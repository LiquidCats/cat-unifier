package eth

import (
	"cat-unifier/internal/kernel/common/contracts"
	"cat-unifier/pkg/reader/eth/common"
	"github.com/ybbus/jsonrpc"
	"log"
)

type ethReader struct {
	client jsonrpc.RPCClient
}

func mapTransactionsFromResponse(response map[string]interface{}) ([]*common.Transaction, []*contracts.Transaction) {
	responseTransaction := response["rawTransactions"].([]map[string]interface{})
	rawTransactions := make([]*common.Transaction, len(responseTransaction))
	transactions := make([]*contracts.Transaction, len(responseTransaction))

	for k, v := range responseTransaction {
		rawTransactions[k] = common.CreateRawTransaction(v)
	}

	for k, v := range rawTransactions {
		transactions[k] = common.CreateTransaction(v)
	}

	return rawTransactions, transactions
}

func (eth *ethReader) GetBalance(id string) float64 {
	var resp string

	err := eth.client.CallFor(&resp, "eth_getBalance", id)
	if nil != err {
		log.Fatal("ETH node is not reachable ERROR:", err)
	}

	return common.WeiToEth(common.DecodeHex(resp))
}

func (eth *ethReader) GetBlock(id interface{}, loadTransactions bool) *contracts.Block {
	var response map[string]interface{}
	var command string

	if _, ok := id.(int); ok {
		command = "eth_getBlockByNumber"
	} else {
		command = "eth_getBlockByHash"
	}

	err := eth.client.CallFor(&response, command, id, loadTransactions)

	if nil != err {
		log.Fatal("ETH node is not reachable ERROR:", err)
	}

	rawTransactions, transactions := mapTransactionsFromResponse(response)

	return common.CreateBlock(common.CreateRawBlock(response, rawTransactions), transactions)
}

func (eth *ethReader) GetTransaction(id string) *contracts.Transaction {
	var response map[string]interface{}

	err := eth.client.CallFor(&response, "eth_getTransaction", id)
	if nil != err {
		log.Fatal("ETH node is not reachable ERROR:", err)
	}

	return common.CreateTransaction(common.CreateRawTransaction(response))
}

func NewReader(nodeUrl string) contracts.IReader {
	client := jsonrpc.NewClient(nodeUrl)

	return &ethReader{
		client,
	}
}
