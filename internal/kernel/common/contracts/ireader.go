package contracts

import (
	"github.com/ybbus/jsonrpc"
)

type Block struct {
	Number           uint64
	Hash             string
	Timestamp        uint64
	Transactions     []*Transaction
	TransactionCount int
	Original         interface{}
}

type Transaction struct {
	Hash        string
	BlockHash   string
	BlockHumber uint64
	From        string
	To          string
	Value       float64
	Original    interface{}
}

type IReader interface {
	GetBalance(id string) float64
	GetBlock(id interface{}, withTransactions bool) *Block
	GetTransaction(id string) *Transaction
	//GetTransactionsForAddress(address string) []Transaction
	//GetTransactionsForBlock(block string) []Transaction
}

type Reader struct {
	client jsonrpc.RPCClient
}
