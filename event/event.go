package event

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/kitshy/treasure-backend/common/bigint"
	"github.com/kitshy/treasure-backend/common/tasks"
	"github.com/kitshy/treasure-backend/config"
	"github.com/kitshy/treasure-backend/database"
	"github.com/kitshy/treasure-backend/database/chain"
	"github.com/kitshy/treasure-backend/event/contracts"
	"gorm.io/gorm"
	"math/big"
	"time"
)

var blocksLimit = 10_000

type Event struct {
	conf *config.Config
	db   *database.DB

	LatestBlockHeader *chain.BlockHeaders

	resourceCtx    context.Context
	resourceCancel context.CancelFunc
	tasks          tasks.Group
}

func NewEvent(conf *config.Config, db *database.DB, shutdown context.CancelCauseFunc) (*Event, error) {

	startHeight := new(big.Int).SetUint64(conf.Chain.StartHeight)

	LatestBlock, err := db.BlockHeaders.LatestBlockHeader()
	if err != nil {
		return nil, err
	}

	if LatestBlock.Number.Cmp(startHeight) < 0 {
		return nil, fmt.Errorf("conf start height %d is too high", conf.Chain.StartHeight)
	} else if LatestBlock.Number.Cmp(startHeight) > 0 {
		block, err := db.BlockHeaders.QueryBlockHeaders(&chain.BlockHeaders{Number: new(big.Int).Add(startHeight, bigint.One)})
		if err != nil {
			return nil, err
		}
		if len(block) <= 0 {
			return nil, fmt.Errorf("query block config height bug", startHeight)
		}
		LatestBlock = &block[0]
	}

	log.Info("init LatestBlock is : ", LatestBlock.Number)

	ctx, cancel := context.WithCancel(context.Background())
	return &Event{
		conf:              conf,
		db:                db,
		LatestBlockHeader: LatestBlock,
		resourceCtx:       ctx,
		resourceCancel:    cancel,
		tasks: tasks.Group{
			HandleCrit: func(err error) {
				shutdown(fmt.Errorf("error handling critical transaction: %v", err))
			},
		},
	}, nil
}

func (e *Event) Start() error {

	log.Info("starting event ..... ")

	tickerSyncer := time.NewTicker(time.Second * 3)
	e.tasks.Go(func() error {

		for range tickerSyncer.C {
			err := e.processTreasureEvent()
			if err != nil {
				return err
			}
		}
		return nil
	})

	return nil
}

func (e *Event) Close() error {
	e.resourceCancel()
	return e.db.Close()
}

func (e *Event) processTreasureEvent() error {

	lastBlockNumber := big.NewInt(int64(e.conf.Chain.StartHeight))
	if e.LatestBlockHeader != nil {
		lastBlockNumber = e.LatestBlockHeader.Number
	}

	latestHeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Model(chain.BlockHeaders{}).Where("number > ?", lastBlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(blocksLimit)).Select("MAX(number)"))
	}
	if latestHeaderScope == nil {
		return fmt.Errorf("latest block header scope is nil")
	}
	latestHeader, err := e.db.BlockHeaders.BlockHeaderWithScope(latestHeaderScope)
	if err != nil {
		return fmt.Errorf("failed to query for latest unfinalized treasure manager events state: %w", err)
	} else if latestHeader == nil {
		log.Debug("no new state to process event")
		return nil
	}

	fromHeight, toHeight := new(big.Int).Add(lastBlockNumber, bigint.One), latestHeader.Number

	if err := e.db.Transaction(func(tx *database.DB) error {
		log.Info("scanning for treasure manager events", "fromHeight", fromHeight, "toHeight", toHeight)
		return contracts.TreasureProcessor(tx, e.conf.Chain, fromHeight, toHeight)
	}); err != nil {
		return err
	}
	e.LatestBlockHeader = latestHeader
	log.Info("finished processing treasure manager events...end block..", latestHeader.Number)
	return nil
}
