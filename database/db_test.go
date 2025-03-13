package database

import (
	"context"
	"github.com/kitshy/treasure-backend/config"
	"testing"
)

func setUp() *DB {
	conf := config.DBConfig{
		Host: "localhost",
		Port: 5432,
		User: "postgres",
		Pass: "postgres",
		Name: "postgres",
	}
	db, err := NewDB(context.Background(), conf)
	if err != nil {
		panic(err)
	}
	return db
}

func TestDB_ExecuteSQLMigration(t *testing.T) {
	db := setUp()
	db.ExecuteSQLMigration("")
}
