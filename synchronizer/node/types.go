package node

import "github.com/ethereum/go-ethereum/core/types"

type BlockAndLogs struct {
	Logs          []types.Log
	ToBlockHeader *types.Header
}
