package node

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/url"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/kitshy/treasure-backend/common/global_const"
	"github.com/kitshy/treasure-backend/common/retry"
)

const (
	defaultDialTimeout    = 5 * time.Second
	defaultDialAttempts   = 3
	defaultRequestTimeout = 100 * time.Second
)

type Client struct {
	rpc RPC
}

type EvmClient interface {
	BlockHeaderByNumber(*big.Int) (*types.Header, error)
	LatestSafeBlockHeader() (*types.Header, error)
	LatestFinalizedBlockHeader() (*types.Header, error)
	BlockHeaderByHash(common.Hash) (*types.Header, error)
	BlockHeadersByRange(*big.Int, *big.Int, uint) ([]types.Header, error)
	TxByHash(common.Hash) (*types.Transaction, error)
	StorageHash(common.Address, *big.Int) (common.Hash, error)

	BlockReceipts(common.Hash) ([]types.Receipt, error)
	TransactionReceipts(common.Hash) (types.Receipt, error)

	NewFilter(*big.Int, *big.Int, common.Address) (string, error)
	UninstallFilter(string) error
	FilterChanges(string) ([]types.Log, error)
	FilterLogs(string) ([]types.Log, error)

	GetLogs(*big.Int, *big.Int, common.Address) ([]types.Log, error)

	BatchBlockAndLogs(ethereum.FilterQuery) (BlockAndLogs, error)

	Close()
}

func NewEvmClient(ctx context.Context, rpcUrl string) (EvmClient, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultDialTimeout)
	defer cancel()
	off := retry.Exponential()
	rpcClient, err := retry.Do(ctx, defaultDialAttempts, off, func() (*rpc.Client, error) {
		if !IsURLAvailable(rpcUrl) {
			return nil, fmt.Errorf("rpc url %s is not available", rpcUrl)
		}
		client, err := rpc.DialContext(ctx, rpcUrl)
		if err != nil {
			return nil, err
		}
		return client, nil
	})
	if err != nil {
		return nil, err
	}
	return Client{rpc: NewRPC(rpcClient)}, nil
}

func (c Client) BlockHeaderByNumber(number *big.Int) (*types.Header, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var header *types.Header
	err := c.rpc.CallContext(ctxwt, &header, "eth_getBlockByNumber", toBlockNumArg(number), false)
	if err != nil {
		log.Fatalln("Call eth_getBlockByNumber method fail", "err", err)
		return nil, err
	} else if header == nil {
		log.Println("header not found")
		return nil, ethereum.NotFound
	}

	return header, nil
}

func (c Client) LatestSafeBlockHeader() (*types.Header, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var header *types.Header
	err := c.rpc.CallContext(ctxwt, &header, "eth_getBlockByNumber", "safe", false)
	if err != nil {
		return nil, err
	} else if header == nil {
		return nil, ethereum.NotFound
	}

	return header, nil
}

func (c Client) LatestFinalizedBlockHeader() (*types.Header, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var header *types.Header
	err := c.rpc.CallContext(ctxwt, &header, "eth_getBlockByNumber", "finalized", false)
	if err != nil {
		return nil, err
	} else if header == nil {
		return nil, ethereum.NotFound
	}

	return header, nil
}

func (c Client) BlockHeaderByHash(hash common.Hash) (*types.Header, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var header *types.Header
	err := c.rpc.CallContext(ctxwt, &header, "eth_getBlockByHash", hash, false)
	if err != nil {
		return nil, err
	} else if header == nil {
		return nil, ethereum.NotFound
	}

	if header.Hash() != hash {
		return nil, errors.New("header mismatch")
	}

	return header, nil
}

func (c Client) BlockHeadersByRange(startHeight *big.Int, endHeight *big.Int, chainId uint) ([]types.Header, error) {
	if startHeight.Cmp(endHeight) == 0 {
		header, err := c.BlockHeaderByNumber(startHeight)
		if err != nil {
			return nil, err
		}
		return []types.Header{*header}, nil
	}

	count := new(big.Int).Sub(endHeight, startHeight).Uint64() + 1
	headers := make([]types.Header, count)
	batchElems := make([]rpc.BatchElem, count)

	if chainId != uint(global_const.PolygonChainId) {
		for i := uint64(0); i < count; i++ {
			height := new(big.Int).Add(startHeight, new(big.Int).SetUint64(i))
			batchElems[i] = rpc.BatchElem{Method: "eth_getBlockByNumber", Args: []interface{}{toBlockNumArg(height), false}, Result: &headers[i]}
		}

		ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
		defer cancel()
		err := c.rpc.BatchCallContext(ctxwt, batchElems)
		if err != nil {
			return nil, err
		}
	} else {
		groupSize := 100
		var wg sync.WaitGroup
		numGroups := (int(count)-1)/groupSize + 1
		wg.Add(numGroups)

		for i := 0; i < int(count); i += groupSize {
			start := i
			end := i + groupSize - 1
			if end > int(count) {
				end = int(count) - 1
			}
			go func(start, end int) {
				defer wg.Done()
				for j := start; j <= end; j++ {
					ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
					defer cancel()
					height := new(big.Int).Add(startHeight, new(big.Int).SetUint64(uint64(j)))
					batchElems[j] = rpc.BatchElem{
						Method: "eth_getBlockByNumber",
						Result: new(types.Header),
						Error:  nil,
					}
					header := new(types.Header)
					batchElems[j].Error = c.rpc.CallContext(ctxwt, header, batchElems[j].Method, toBlockNumArg(height), false)
					batchElems[j].Result = header
				}
			}(start, end)
		}

		wg.Wait()
	}
	size := 0
	for i, batchElem := range batchElems {
		header, ok := batchElem.Result.(*types.Header)
		if !ok {
			return nil, fmt.Errorf("unable to transform rpc response %v into utils.Header", batchElem.Result)
		}

		headers[i] = *header

		size = size + 1
	}
	headers = headers[:size]

	return headers, nil
}

func (c Client) TxByHash(hash common.Hash) (*types.Transaction, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var tx *types.Transaction
	err := c.rpc.CallContext(ctxwt, &tx, "eth_getTransactionByHash", hash)
	if err != nil {
		return nil, err
	} else if tx == nil {
		return nil, ethereum.NotFound
	}

	return tx, nil
}

func (c Client) StorageHash(address common.Address, blockNumber *big.Int) (common.Hash, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	proof := struct{ StorageHash common.Hash }{}
	err := c.rpc.CallContext(ctxwt, &proof, "eth_getProof", address, nil, toBlockNumArg(blockNumber))
	if err != nil {
		return common.Hash{}, err
	}

	return proof.StorageHash, nil
}

func (c Client) BlockReceipts(blockHash common.Hash) ([]types.Receipt, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var receipt []types.Receipt
	err := c.rpc.CallContext(ctxwt, &receipt, "eth_getBlockReceipts", blockHash)
	if err != nil {
		return []types.Receipt{}, err
	}
	return receipt, nil
}

func (c Client) TransactionReceipts(txHash common.Hash) (types.Receipt, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var receipt types.Receipt
	err := c.rpc.CallContext(ctxwt, &receipt, "eth_getTransactionReceipt", txHash)
	if err != nil {
		return types.Receipt{}, err
	}
	return receipt, nil
}

func (c Client) NewFilter(fromBlock *big.Int, toBlock *big.Int, contractAddress common.Address) (string, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var filterId string
	var args map[string]interface{}
	from := toBlockNumArg(fromBlock)
	if toBlock != nil {
		to := toBlockNumArg(toBlock)
		args = map[string]interface{}{
			"fromBlock": from,
			"toBlock":   to,
			"address":   []common.Address{contractAddress},
		}
	} else {
		args = map[string]interface{}{
			"fromBlock": from,
			"toBlock":   "safe",
			"address":   []common.Address{contractAddress},
		}
	}

	err := c.rpc.CallContext(ctxwt, &filterId, "eth_newFilter", args)
	if err != nil {
		return "", err
	}
	return filterId, nil
}

func (c Client) UninstallFilter(filterId string) error {

	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	err := c.rpc.CallContext(ctxwt, nil, "eth_uninstallFilter", filterId)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) FilterChanges(filterId string) ([]types.Log, error) {

	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var logs []types.Log
	err := c.rpc.CallContext(ctxwt, &logs, "eth_getFilterChanges", filterId)
	if err != nil {
		return []types.Log{}, err
	}
	return logs, nil

}

func (c Client) FilterLogs(filterId string) ([]types.Log, error) {

	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var logs []types.Log
	err := c.rpc.CallContext(ctxwt, &logs, "eth_getFilterLogs", filterId)
	if err != nil {
		return []types.Log{}, err
	}
	return logs, nil

}

func (c Client) GetLogs(fromBlock *big.Int, toBlock *big.Int, contractAddress common.Address) ([]types.Log, error) {

	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var logs []types.Log

	var args = map[string]interface{}{
		"fromBlock": toBlockNumArg(fromBlock),
		"toBlock":   toBlockNumArg(toBlock),
		"address":   []common.Address{contractAddress},
	}

	err := c.rpc.CallContext(ctxwt, &logs, "eth_getLogs", args)
	if err != nil {
		return []types.Log{}, err
	}
	return logs, nil
}

func (c Client) BatchBlockAndLogs(query ethereum.FilterQuery) (BlockAndLogs, error) {

	params, err := toFilterArg(query)
	if err != nil {
		return BlockAndLogs{}, err
	}

	var logResult []types.Log
	var blockResult types.Header

	batchElems := make([]rpc.BatchElem, 2)
	batchElems[0] = rpc.BatchElem{
		Method: "eth_getBlockByNumber",
		Args:   []interface{}{toBlockNumArg(query.ToBlock), false},
		Result: &blockResult,
	}
	batchElems[1] = rpc.BatchElem{
		Method: "eth_getLogs",
		Args:   []interface{}{params},
		Result: &logResult,
	}

	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	err = c.rpc.BatchCallContext(ctxwt, batchElems)
	if err != nil {
		return BlockAndLogs{}, err
	}
	return BlockAndLogs{
		Logs:          logResult,
		ToBlockHeader: &blockResult,
	}, nil
}

func (c Client) Close() {
	c.rpc.Close()
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	if number.Sign() >= 0 {
		return hexutil.EncodeBig(number)
	}
	return rpc.BlockNumber(number.Int64()).String()
}

func toFilterArg(q ethereum.FilterQuery) (interface{}, error) {
	arg := map[string]interface{}{"address": q.Addresses, "topics": q.Topics}
	if q.BlockHash != nil {
		arg["blockHash"] = *q.BlockHash
		if q.FromBlock != nil || q.ToBlock != nil {
			return nil, errors.New("cannot specify both BlockHash and FromBlock/ToBlock")
		}
	} else {
		if q.FromBlock == nil {
			arg["fromBlock"] = "0x0"
		} else {
			arg["fromBlock"] = toBlockNumArg(q.FromBlock)
		}
		arg["toBlock"] = toBlockNumArg(q.ToBlock)
	}
	return arg, nil
}

func IsURLAvailable(address string) bool {
	u, err := url.Parse(address)
	if err != nil {
		return false
	}
	addr := u.Host
	if u.Port() == "" {
		switch u.Scheme {
		case "http", "ws":
			addr += ":80"
		case "https", "wss":
			addr += ":443"
		default:
			return true
		}
	}
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		return false
	}
	err = conn.Close()
	if err != nil {
		return false
	}
	return true
}
