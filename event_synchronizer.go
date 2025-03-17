package treasure_backend

import (
	"context"
	"github.com/kitshy/treasure-backend/event"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/log"
	"github.com/kitshy/treasure-backend/config"
	"github.com/kitshy/treasure-backend/database"
	"github.com/kitshy/treasure-backend/synchronizer"
	"github.com/kitshy/treasure-backend/synchronizer/node"
)

type EventSynchronizer struct {
	synchronizer *synchronizer.Synchronizer
	event        *event.Event

	shutdown context.CancelCauseFunc
	stopped  atomic.Bool
}

func NewEventSynchronizer(ctx context.Context, conf *config.Config, shutdown context.CancelCauseFunc) (*EventSynchronizer, error) {

	db, err := database.NewDB(ctx, conf.MasterDB)
	if err != nil {
		log.Error("init db error")
		return nil, err
	}

	log.Info("init event synchronizer ====== ", "chain_rpc_url", conf.Chain.ChainRpcUrl)
	evmClient, err := node.NewEvmClient(ctx, conf.Chain.ChainRpcUrl)
	if err != nil {
		log.Error("init evm client error")
		return nil, err
	}

	synchronizer, err := synchronizer.NewSynchronizer(conf, db, evmClient, shutdown)
	if err != nil {
		log.Error("init synchronizer error")
		return nil, err
	}

	event, err := event.NewEvent(conf, db, shutdown)
	if err != nil {
		log.Error("init event error")
		return nil, err
	}

	return &EventSynchronizer{
		synchronizer: synchronizer,
		shutdown:     shutdown,
		event:        event,
	}, nil

}

func (s *EventSynchronizer) Start(ctx context.Context) error {
	err := s.synchronizer.Start()
	if err != nil {
		log.Error("start treasure synchronizer error")
		return err
	}
	if err := s.event.Start(); err != nil {
		log.Error("start treasure synchronizer error")
		return err
	}
	return nil
}

func (s *EventSynchronizer) Stop(ctx context.Context) error {
	err := s.synchronizer.Close()
	if err != nil {
		log.Error("close treasure synchronizer error")
		return err
	}
	if err := s.event.Close(); err != nil {
		log.Error("close event error")
		return err
	}
	return nil
}

func (s *EventSynchronizer) Stopped() bool {
	return s.stopped.Load()
}
