package test

import (
	"github.com/google/uuid"
	"log"
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
