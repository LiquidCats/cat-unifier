package common

import (
	"cat-unifier/internal/kernel/common/contracts"
	"fmt"
)

func CreateRawBlock(data map[string]interface{}, rawTransactions []*Transaction) *Block {
	return &Block{
		Difficulty:       uint64(DecodeHex(fmt.Sprintf("%v", data["difficulty"]))),
		GasLimit:         uint64(DecodeHex(fmt.Sprintf("%v", data["gasLimit"]))),
		GasUsed:          uint64(DecodeHex(fmt.Sprintf("%v", data["gasUsed"]))),
		Hash:             fmt.Sprintf("%v", data["hash"]),
		LogsBloom:        fmt.Sprintf("%v", data["logsBloom"]),
		Miner:            fmt.Sprintf("%v", data["miner"]),
		MixHash:          fmt.Sprintf("%v", data["mixHash"]),
		Nonce:            fmt.Sprintf("%v", data["nonce"]),
		Number:           uint64(DecodeHex(fmt.Sprintf("%v", data["number"]))),
		ParentHash:       fmt.Sprintf("%v", data["parentHash"]),
		ReceiptsRoot:     fmt.Sprintf("%v", data["receiptsRoot"]),
		Sha3Uncles:       fmt.Sprintf("%v", data["sha3Uncles"]),
		Size:             uint64(DecodeHex(fmt.Sprintf("%v", data["size"]))),
		StateRoot:        fmt.Sprintf("%v", data["stateRoot"]),
		Timestamp:        uint64(DecodeHex(fmt.Sprintf("%v", data["timestamp"]))),
		TotalDifficulty:  uint64(DecodeHex(fmt.Sprintf("%v", data["totalDifficulty"]))),
		TransactionsRoot: fmt.Sprintf("%v", data["transactionsRoot"]),
		Transactions:     rawTransactions,
		Uncles:           data["uncles"],
	}
}

func CreateBlock(rawBlock *Block, transactions []*contracts.Transaction) *contracts.Block {
	return &contracts.Block{
		Hash:             rawBlock.Hash,
		Number:           rawBlock.Number,
		Timestamp:        rawBlock.Timestamp,
		Transactions:     transactions,
		TransactionCount: len(transactions),
		Original:         rawBlock,
	}
}

func CreateRawTransaction(data map[string]interface{}) *Transaction {
	return &Transaction{
		Hash:             fmt.Sprintf("%v", data["hash"]),
		Nonce:            uint64(DecodeHex(fmt.Sprintf("%v", data["nonce"]))),
		BlockHash:        fmt.Sprintf("%v", data["blockHash"]),
		BlockNumber:      uint64(DecodeHex(fmt.Sprintf("%v", data["blockNumber"]))),
		TransactionIndex: uint64(DecodeHex(fmt.Sprintf("%v", data["transactionIndex"]))),
		From:             fmt.Sprintf("%v", data["from"]),
		To:               fmt.Sprintf("%v", data["to"]),
		Value:            WeiToEth(DecodeHex(fmt.Sprintf("%v", data["value"]))),
		Gas:              uint64(DecodeHex(fmt.Sprintf("%v", data["gas"]))),
		GasPrice:         uint64(DecodeHex(fmt.Sprintf("%v", data["gasPrice"]))),
		Input:            fmt.Sprintf("%v", data["input"]),
	}
}

func CreateTransaction(rawTransaction *Transaction) *contracts.Transaction {
	return &contracts.Transaction{
		Hash:        rawTransaction.Hash,
		BlockHash:   rawTransaction.BlockHash,
		BlockHumber: rawTransaction.BlockNumber,
		From:        rawTransaction.From,
		To:          rawTransaction.To,
		Value:       rawTransaction.Value,
		Original:    rawTransaction,
	}
}
