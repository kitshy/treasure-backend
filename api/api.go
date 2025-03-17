package api

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kitshy/treasure-backend/api/common/httputil"
	"github.com/kitshy/treasure-backend/api/routes"
	"github.com/kitshy/treasure-backend/api/service"
	"github.com/kitshy/treasure-backend/api/validator"
	"github.com/kitshy/treasure-backend/config"
	"github.com/kitshy/treasure-backend/database"
)

const (
	HealthPath = "/healthz"

	GetDepositTokensListV1Path = "/api/v1/deposit/GetDepositTokensList"
	GetGrantRewardTokensV1Path = "/api/v1/grant/GetGrantRewardTokens"
)

type APIConfig struct {
	HTTPServer    config.ServerConfig
	MetricsServer config.ServerConfig
}

type Api struct {
	router    *chi.Mux
	apiServer *httputil.HTTPServer
	db        *database.DB

	context context.Context
	stopped atomic.Bool
}

func NewApi(conf *config.Config, context context.Context) (*Api, error) {
	api := Api{}
	err := api.initDB(context, conf)
	if err != nil {
		return nil, errors.Join(err, api.Stop(context))
	}

	if api.initRoute(context, conf); err != nil {
		return nil, errors.Join(err, api.Stop(context))
	}

	if api.startServer(conf.HttpService); err != nil {
		return nil, errors.Join(err, api.Stop(context))
	}

	return &api, nil
}

func (api *Api) Start(ctx context.Context) error {
	return nil
}

func (api *Api) Stop(ctx context.Context) error {

	return nil
}

func (api *Api) Stopped() bool {
	return api.stopped.Load()
}

func (api *Api) initDB(ctx context.Context, cfg *config.Config) error {

	var initDb *database.DB
	var err error
	if !cfg.SlaveDbEnable {
		initDb, err = database.NewDB(ctx, cfg.MasterDB)
		if err != nil {
			log.Error("failed to connect to master database", "err", err)
			return err
		}
	} else {
		initDb, err = database.NewDB(ctx, cfg.SlaveDB)
		if err != nil {
			log.Error("failed to connect to slave database", "err", err)
			return err
		}
	}
	api.db = initDb

	return nil
}

func (api *Api) initRoute(ctx context.Context, cfg *config.Config) error {

	v := new(validator.Validator)

	db := service.NewService(v, api.db.DepositTokens)
	apiRouter := chi.NewRouter()
	service := routes.NewRoutes(apiRouter, db)
	apiRouter.Use(middleware.Timeout(time.Second * 12))
	apiRouter.Use(middleware.Recoverer)

	apiRouter.Use(middleware.Heartbeat(HealthPath))

	// deposit token
	apiRouter.Get(fmt.Sprintf(GetDepositTokensListV1Path), service.GetDepositTokensListHandler)

	// grant reward token
	apiRouter.Get(fmt.Sprintf(GetGrantRewardTokensV1Path), service.GetGrantRewardTokensListHandler)

	api.router = apiRouter
	return nil
}

func (api *Api) startServer(serverConfig config.ServerConfig) error {
	log.Debug("API server listening...", "port", serverConfig.Port)
	addr := net.JoinHostPort(serverConfig.Host, strconv.Itoa(serverConfig.Port))
	srv, err := httputil.StartHTTPServer(addr, api.router)
	if err != nil {
		return fmt.Errorf("failed to start API server: %w", err)
	}
	log.Info("API server started", "addr", srv.Addr().String())
	api.apiServer = srv
	return nil
}
