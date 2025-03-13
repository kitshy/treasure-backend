package node

import (
	"context"
	"github.com/ethereum/go-ethereum/rpc"
)

type RPC interface {
	Close()
	CallContext(ctx context.Context, result interface{}, method string, args ...any) error
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
}

type rpcClient struct {
	rpc *rpc.Client
}

func NewRPC(client *rpc.Client) RPC {
	return &rpcClient{
		rpc: client,
	}
}

func (c *rpcClient) Close() {
	c.rpc.Close()
}

func (c *rpcClient) CallContext(ctx context.Context, result interface{}, method string, args ...any) error {
	return c.rpc.CallContext(ctx, result, method, args...)
}

func (c *rpcClient) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	return c.rpc.BatchCallContext(ctx, b)
}
