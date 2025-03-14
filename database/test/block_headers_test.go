package test

import (
	"context"
	"github.com/google/uuid"
	"github.com/kitshy/treasure-backend/config"
	"github.com/kitshy/treasure-backend/database"
	"testing"
)

func setUp() *database.DB {
	df := config.DBConfig{
		Host: "",
		Port: 5432,
		User: "root",
		Name: "treasure",
		Pass: "",
	}
	db, err := database.NewDB(context.Background(), df)
	if err != nil {
		panic(err)
	}
	return db
}

func TestBlockHeadersDB_QueryBlockHeadersByID(t *testing.T) {
	db := setUp()
	rsp, err := db.BlockHeaders.QueryBlockHeadersByID(uuid.UUID{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)
}

func TestBlockHeadersDB_LatestBlockHeader(t *testing.T) {
	db := setUp()
	rsp, err := db.BlockHeaders.LatestBlockHeader()
	if err != nil {
		t.Error(err)
	}
	t.Log(rsp)
}
