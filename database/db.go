package database

import (
	"context"
	"fmt"
	"github.com/kitshy/treasure-backend/database/eventlog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/log"

	"github.com/kitshy/treasure-backend/database/chain"
	"github.com/kitshy/treasure-backend/database/event"

	"github.com/kitshy/treasure-backend/common/retry"
	"github.com/kitshy/treasure-backend/config"
	_ "github.com/kitshy/treasure-backend/database/utils/serializers"
)

type DB struct {
	gorm              *gorm.DB
	BlockHeaders      chain.BlockHeadersDB
	ContractEvents    event.ContractEventsDB
	DepositTokens     eventlog.DepositTokensDB
	GrantRewardTokens eventlog.GrantRewardTokensDB
}

func NewDB(ctx context.Context, dbConfig config.DBConfig) (*DB, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Name)
	if dbConfig.Port != 0 {
		dsn += fmt.Sprintf(" port=%d", dbConfig.Port)
	}
	if dbConfig.User != "" {
		dsn += fmt.Sprintf(" user=%s", dbConfig.User)
	}
	if dbConfig.Pass != "" {
		dsn += fmt.Sprintf(" password=%s", dbConfig.Pass)
	}

	gromConfig := gorm.Config{
		SkipDefaultTransaction: true,
		CreateBatchSize:        3_000,
	}

	log.Info("db config .. ", dsn)
	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	gorm, err := retry.Do[*gorm.DB](context.Background(), 3, retryStrategy, func() (*gorm.DB, error) {
		gorm, err := gorm.Open(postgres.Open(dsn), &gromConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		return gorm, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return &DB{
		gorm:              gorm,
		ContractEvents:    event.NewContractEventsDB(gorm),
		BlockHeaders:      chain.NewBlockHeaders(gorm),
		DepositTokens:     eventlog.NewDepositTokensDB(gorm),
		GrantRewardTokens: eventlog.NewGrantRewardTokensDB(gorm),
	}, nil
}

func (db *DB) Transaction(fn func(db *DB) error) error {
	return db.gorm.Transaction(func(tx *gorm.DB) error {
		txDB := &DB{
			gorm:              tx,
			BlockHeaders:      chain.NewBlockHeaders(tx),
			ContractEvents:    event.NewContractEventsDB(tx),
			DepositTokens:     eventlog.NewDepositTokensDB(tx),
			GrantRewardTokens: eventlog.NewGrantRewardTokensDB(tx),
		}
		return fn(txDB)
	})
}

func (db *DB) Close() error {
	sql, err := db.gorm.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}

func (db *DB) ExecuteSQLMigration(migratorDir string) error {

	err := filepath.Walk(migratorDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to process migration file: %s", path))
		}
		if info.IsDir() {
			return nil
		}
		fileContent, readErr := os.ReadFile(path)
		if readErr != nil {
			return errors.Wrap(readErr, fmt.Sprintf("Error reading SQL file: %s", path))
		}

		execErr := db.gorm.Exec(string(fileContent)).Error
		if execErr != nil {
			return errors.Wrap(execErr, fmt.Sprintf("Error executing SQL script: %s", path))
		}
		return nil
	})
	return err
}
