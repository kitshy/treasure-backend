package main

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	treasure_backend "github.com/kitshy/treasure-backend"
	"github.com/kitshy/treasure-backend/api"
	"github.com/kitshy/treasure-backend/common/cliapp"
	"github.com/urfave/cli/v2"

	"github.com/kitshy/treasure-backend/common/opio"
	"github.com/kitshy/treasure-backend/config"
	"github.com/kitshy/treasure-backend/database"
)

func runApi(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Info("run treasure-backend api ")
	config, err := config.NewConfig(ctx)
	if err != nil {
		return nil, err
	}
	return api.NewApi(config, ctx.Context)
}

func runEventSyn(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Info("run treasure-backend event syn")
	config, err := config.NewConfig(ctx)
	if err != nil {
		return nil, err
	}
	return treasure_backend.NewEventSynchronizer(ctx.Context, config, shutdown)
}

func runMigrations(ctx *cli.Context) error {

	conf, err := config.NewConfig(ctx)
	if err != nil {
		log.Error("Failed to load config", "err", err)
		return err
	}

	ctx.Context = opio.CancelOnInterrupt(ctx.Context)
	db, err := database.NewDB(ctx.Context, conf.MasterDB)
	if err != nil {
		log.Error("Failed to init database", "err", err)
		return err
	}

	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			log.Error("Failed to close database", "err", err)
		}
	}(db)
	return db.ExecuteSQLMigration(conf.Migrations)
}

func NewCli(GitCommit string, GitDate string) *cli.App {

	flags := config.Flags

	return &cli.App{
		Version:              "params.VersionWithCommit(GitCommit, GitDate)",
		Description:          "An indexer of all optimism events with a serving api layer",
		EnableBashCompletion: true,

		Commands: []*cli.Command{
			{
				Name:        "api",
				Flags:       flags,
				Description: "run api server",
				Action:      cliapp.LifecycleCmd(runApi),
			},
			{
				Name:        "eventSyn",
				Flags:       flags,
				Description: "run Event Syn",
				Action:      cliapp.LifecycleCmd(runEventSyn),
			},
			{
				Name:        "migration",
				Flags:       flags,
				Description: "Run migrations",
				Action:      runMigrations,
			},
		},
	}

}
