package node

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"os"
	"testing"
)

const URL = "https://eth-sepolia.g.alchemy.com/v2/tQ6olBJhhJY0_v04kavsliE758vz_qE_"

func setUp(url string) (EvmClient, error) {
	client, err := NewEvmClient(context.Background(), url)
	if err != nil {
		panic(err)
	}
	return client, nil
}

func TestClient_BlockReceipts(t *testing.T) {

	client, err := setUp(URL)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.BlockReceipts(common.HexToHash("0xeeda5deb531db32abc96967b3dbad0434cc9134acc1e0136ba56074e37eb6fc2"))

	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)

}
func TestClient_TransactionReceipts(t *testing.T) {

	client, err := setUp(URL)
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := client.TransactionReceipts(common.HexToHash("0x246cc16a96df3592762cf2bbb3a5c7bd085f249b7d64c4ee3343fc95228a0539"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)

}

func TestClient_NewFilter(t *testing.T) {
	client, err := setUp(URL)
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := client.NewFilter(big.NewInt(0), nil, common.HexToAddress("0xDFC97b057eF039772F1bD7e8acf18949B660Cff1"))

	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)

}

func TestClient_UninstallFilter(t *testing.T) {
	client, err := setUp(URL)
	if err != nil {
		t.Fatal(err)
	}
	err = client.UninstallFilter("0x17159b13d8defba5e9258e51bbf0779b")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("uninstall filter success")
	}
}

func TestClient_FilterChanges(t *testing.T) {
	client, err := setUp(URL)
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := client.FilterChanges("0x40cb989c51cb10348074557d052aa756")
	if err != nil {

	}
	t.Log(rsp)

}

func TestClient_FilterLogs(t *testing.T) {
	client, err := setUp(URL)
	if err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := client.FilterLogs("0x17159b13d8defba5e9258e51bbf0779b")
	if err != nil {

	}
	t.Log(rsp)

}

func TestClient_GetLogs(t *testing.T) {

	client, err := setUp(URL)
	if err != nil {
		t.Fatal(err)
	}
	h := log.NewTerminalHandler(os.Stdout, true)
	log.SetDefault(log.NewLogger(h))

	rsp, err := client.GetLogs(big.NewInt(7872300), big.NewInt(7872350), common.HexToAddress("0xDFC97b057eF039772F1bD7e8acf18949B660Cff1"))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rsp)

}
