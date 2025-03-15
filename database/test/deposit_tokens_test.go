package test

import (
	"log"
	"testing"
)

func TestDepositTokens_QueryDepositTokensByPage(t *testing.T) {
	db := setUp()
	rsp, err := db.DepositTokens.QueryDepositTokensByPage(1, 10, "timestamp desc")
	if err != nil {
		t.Fatal(err)
	}
	log.Println(rsp)
}

func TestDepositTokens_SaveDepositTokens(t *testing.T) {

}
