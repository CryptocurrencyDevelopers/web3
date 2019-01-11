package main

import (
	"context"
	"log"
	"math/big"

	"github.com/gochain-io/gochain/consensus/clique"
	"github.com/gochain-io/gochain/common"
	"github.com/gochain-io/gochain/core/types"
	"github.com/gochain-io/gochain/goclient"
)

type RPCClient struct {
	url    string
	client *goclient.Client
}

func GetClient(rpcURL string) *RPCClient {
	client, err := goclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Cannot connect to the network %q: %v", rpcURL, err)
	}
	rpc := &RPCClient{
		url:    rpcURL,
		client: client,
	}
	return rpc
}

func (rpc *RPCClient) GetBalance(address string, blockNumber *big.Int) (*big.Int, error) {
	return rpc.client.BalanceAt(context.Background(), common.HexToAddress(address), blockNumber)
}

func (rpc *RPCClient) GetCode(address string, blockNumber *big.Int) ([]byte, error) {
	return rpc.client.CodeAt(context.Background(), common.HexToAddress(address), blockNumber)
}

func (rpc *RPCClient) GetBlockByNumber(number *big.Int) (*types.Block, error) {
	return rpc.client.BlockByNumber(context.Background(), number)
}

func (rpc *RPCClient) GetTransactionByHash(hash string) (*types.Transaction, bool, error) {
	return rpc.client.TransactionByHash(context.Background(), common.HexToHash(hash))
}

func (rpc *RPCClient) GetSnapshot() (*clique.Snapshot, error) {
	return rpc.client.SnapshotAt(context.Background(), nil)
}
