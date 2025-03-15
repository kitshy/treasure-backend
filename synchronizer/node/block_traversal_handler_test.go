package node

import (
	"log"
	"math/big"
	"testing"
)

func TestBlockTraversal_ExecuteBlockTraversal(t *testing.T) {
	client, err := setUp(URL)
	if err != nil {
		t.Fatal(err)
	}
	bt := NewBlockTraversal(client, nil, big.NewInt(0), 11155111)
	rsp, err := bt.ExecuteBlockTraversal(2)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(rsp)
}
