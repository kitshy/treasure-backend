package test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/kitshy/treasure-backend/database/event"
	"log"
	"math/big"
	"testing"
)

func TestContractEvents_QueryContractEventsByID(t *testing.T) {

	db := setUp()
	id := uuid.New()
	rsp, err := db.ContractEvents.QueryContractEventsByID(id)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(rsp)

}

func TestContractEvents_ContractEventsWithFilter(t *testing.T) {
	db := setUp()
	c := event.ContractEvents{ContractAddress: common.HexToAddress("0xDFC97b057eF039772F1bD7e8acf18949B660Cff1")}
	rsp, err := db.ContractEvents.ContractEventsWithFilter(c, big.NewInt(0), big.NewInt(7872416))
	if err != nil {
		t.Fatal(err)
	}
	log.Println(rsp)
}
