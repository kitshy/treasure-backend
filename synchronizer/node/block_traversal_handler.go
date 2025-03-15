package node

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/kitshy/treasure-backend/common/bigint"
	"math/big"
)

type BlockTraversal struct {
	evmClient EvmClient

	latestHeader        types.Header
	lastTraversedHeader *types.Header

	blockConfirmationDepth *big.Int
	chainId                uint
}

func NewBlockTraversal(evmClient EvmClient, fromHeader *types.Header, confDepth *big.Int, chainId uint) *BlockTraversal {
	return &BlockTraversal{
		evmClient:              evmClient,
		chainId:                chainId,
		lastTraversedHeader:    fromHeader,
		blockConfirmationDepth: confDepth,
	}
}

func (b *BlockTraversal) LatestHeader() *types.Header {
	return &b.latestHeader
}

func (f *BlockTraversal) LastTraversedHeader() *types.Header {
	return f.lastTraversedHeader
}

func (b *BlockTraversal) ExecuteBlockTraversal(step uint64) ([]types.Header, error) {

	latestBlock, err := b.evmClient.LatestFinalizedBlockHeader()
	if err != nil {
		log.Error("Failed to get latest finalized block header", "err", err)
		return nil, fmt.Errorf("failed to fetch latest finalized block header: %v", err)
	} else if latestBlock == nil {
		log.Error("block header not found")
		return nil, fmt.Errorf("latest finalized block header is nil")
	} else {
		b.latestHeader = *latestBlock
	}

	endHeight := new(big.Int).Sub(b.latestHeader.Number, b.blockConfirmationDepth)
	if endHeight.Sign() < 0 {
		log.Info("wait blocking .... ")
		return nil, nil
	}

	if b.lastTraversedHeader != nil {
		cmp := b.lastTraversedHeader.Number.Cmp(endHeight)
		if cmp == 0 {
			return nil, nil
		} else if cmp > 0 {
			errors.New("treasure scan block in front of the RPC node ")
			return nil, fmt.Errorf("treasure scan block in front of the RPC node ")
		}
	}

	startHeight := bigint.Zero
	if b.lastTraversedHeader != nil {
		startHeight = new(big.Int).Add(b.lastTraversedHeader.Number, bigint.One)
	}
	endHeight = bigint.Clamp(startHeight, endHeight, step)
	headers, err := b.evmClient.BlockHeadersByRange(startHeight, endHeight, b.chainId)
	if err != nil {
		return nil, err
	}

	count := len(headers)
	if count == 0 {
		return nil, nil
	} else if b.lastTraversedHeader != nil && b.lastTraversedHeader.Hash() != headers[0].ParentHash {
		return nil, fmt.Errorf("scan node index bug")
	}

	b.lastTraversedHeader = &headers[count-1]

	return headers, nil
}
