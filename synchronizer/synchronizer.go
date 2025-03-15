package synchronizer

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/kitshy/treasure-backend/common/retry"
	"github.com/kitshy/treasure-backend/database/chain"
	"github.com/kitshy/treasure-backend/database/event"
	"github.com/kitshy/treasure-backend/database/utils"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/kitshy/treasure-backend/common/tasks"
	"github.com/kitshy/treasure-backend/config"
	"github.com/kitshy/treasure-backend/database"
	"github.com/kitshy/treasure-backend/synchronizer/node"
)

type Synchronizer struct {
	evmClient node.EvmClient
	conf      *config.Config
	db        *database.DB

	startHeight    *big.Int
	headers        []types.Header
	latestHeader   *types.Header
	blockTraversal *node.BlockTraversal

	resourceCtx    context.Context
	resourceCancel context.CancelFunc
	tasks          tasks.Group
}

func NewSynchronizer(conf *config.Config, db *database.DB, evmClient node.EvmClient, shutdown context.CancelCauseFunc) (*Synchronizer, error) {

	latestBlock, err := db.BlockHeaders.LatestBlockHeader()
	if err != nil {
		return nil, err
	}
	var fromBlock *types.Header
	if latestBlock != nil {
		fromBlock = latestBlock.RlpHeader.Header()
	} else {

		header, err := evmClient.BlockHeaderByNumber(big.NewInt(int64(conf.Chain.StartHeight)))
		if err != nil {
			return nil, err
		}
		log.Info("evmClient.BlockHeaderByNumber : ", header.Number, conf.Chain.StartHeight)
		fromBlock = header
	}
	log.Info("synchronizer initialized", "fromBlock", fromBlock.Number)

	blockTraversal := node.NewBlockTraversal(evmClient, fromBlock, big.NewInt(0), conf.Chain.ChainId)

	ctx, cancel := context.WithCancel(context.Background())

	return &Synchronizer{
		evmClient:      evmClient,
		conf:           conf,
		db:             db,
		startHeight:    fromBlock.Number,
		blockTraversal: blockTraversal,

		resourceCtx:    ctx,
		resourceCancel: cancel,
		tasks: tasks.Group{
			HandleCrit: func(err error) {
				shutdown(fmt.Errorf("error handling critical transaction: %v", err))
			},
		},
	}, nil
}

func (synchronizer *Synchronizer) Start() error {

	tickerSyncer := time.NewTicker(time.Second * 3)
	synchronizer.tasks.Go(func() error {

		for range tickerSyncer.C {
			log.Info("synchronizer starting ..... ")

			if len(synchronizer.headers) > 0 {
				log.Info("execute block step headers")
			} else {

				newHeaders, err := synchronizer.blockTraversal.ExecuteBlockTraversal(synchronizer.conf.Chain.BlockStep)
				if err != nil {
					log.Error("error executing block traversal", "err", err)
					continue
				} else if len(newHeaders) == 0 {
					log.Warn("not new block ... ")
				} else {
					synchronizer.headers = newHeaders
				}
				latestBlock := synchronizer.blockTraversal.LatestHeader()
				if latestBlock != nil {
					log.Info("latest block number", "number", latestBlock.Number)
				}
			}

			err := synchronizer.processBatch(synchronizer.headers, synchronizer.conf)
			log.Info("executing block traversal", "err", err)
			if err == nil {
				synchronizer.headers = nil
			}
		}

		return nil
	})

	return nil
}

func (synchronizer *Synchronizer) processBatch(headers []types.Header, config *config.Config) error {

	if len(headers) == 0 {
		return nil
	}

	startBlock := headers[0]
	endBlock := headers[len(headers)-1]
	headerMap := make(map[common.Hash]*types.Header, len(headers))
	for i := range headers {
		header := headers[i]
		headerMap[header.Hash()] = &header
	}

	log.Info("processing block traversal", "start : ", startBlock.Number.String(), "  end : ", endBlock.Number.String())

	filter := ethereum.FilterQuery{
		FromBlock: startBlock.Number,
		ToBlock:   endBlock.Number,
		Addresses: []common.Address{config.Chain.Contracts[0]},
	}
	blockLogs, err := synchronizer.evmClient.BatchBlockAndLogs(filter)
	if err != nil {
		log.Error("error processing block traversal BatchBlockAndLogs", "err", err)
		return err
	}

	if blockLogs.ToBlockHeader.Number.Cmp(endBlock.Number) != 0 {
		return fmt.Errorf("block traversal batch out of order")
	} else if blockLogs.ToBlockHeader.Hash() != endBlock.Hash() {
		return fmt.Errorf("err to block hash not equal end block hash")
	}

	blockHeaders := make([]chain.BlockHeaders, len(headers))
	for i, header := range headers {
		blockHeader := chain.BlockHeaders{
			GUID:       uuid.New(),
			Hash:       header.Hash(),
			ParentHash: header.ParentHash,
			Number:     header.Number,
			Timestamp:  header.Time,
			RlpHeader:  (*utils.RLPHeader)(&header),
		}
		blockHeaders[i] = blockHeader
	}

	chainContractEvents := make([]event.ContractEvents, 0, len(blockLogs.Logs))
	for i := range blockLogs.Logs {
		logEvent := blockLogs.Logs[i]

		if _, ok := headerMap[logEvent.BlockHash]; !ok {
			continue
		}

		time := headerMap[logEvent.BlockHash].Time
		chainContractEvent := event.ContractEventsFromLog(&logEvent, time)
		chainContractEvents = append(chainContractEvents, *chainContractEvent)
	}

	log.Info("source block number and log num and copy after block and log num : ", blockLogs.ToBlockHeader.Number, len(blockLogs.Logs), len(blockHeaders), len(chainContractEvents))

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](synchronizer.resourceCtx, 3, retryStrategy, func() (interface{}, error) {

		if err := synchronizer.db.Transaction(func(db *database.DB) error {
			if err := db.BlockHeaders.SaveBlockHeaders(&blockHeaders); err != nil {
				return err
			}
			if len(chainContractEvents) > 0 {
				if err := db.ContractEvents.SaveContractEvents(&chainContractEvents); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			log.Error("batch save block and event log error", "err", err)
			return nil, err
		}
		return nil, nil
	}); err != nil {
		return err
	}

	log.Info("batch save block and logs for Transaction success .... ")
	return nil
}

func (synchronizer *Synchronizer) Close() error {
	return nil
}
