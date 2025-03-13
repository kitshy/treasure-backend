package synchronizer

import (
	"context"
	"fmt"
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
	evmClient   node.EvmClient
	conf        *config.Config
	db          *database.DB
	startHeight *big.Int

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
		fromBlock = header
	}
	log.Info("synchronizer initialized", "fromBlock", fromBlock.Number.String())

	ctx, cancel := context.WithCancel(context.Background())

	return &Synchronizer{
		evmClient:      evmClient,
		conf:           conf,
		db:             db,
		startHeight:    fromBlock.Number,
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
			synchronizer.processBatch()
		}

		return nil
	})

	return nil
}

func (synchronizer *Synchronizer) processBatch() error {
	return nil
}

func (synchronizer *Synchronizer) Close() error {
	return nil
}
