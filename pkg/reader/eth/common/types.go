package common

type Block struct {
	Difficulty       uint64
	GasLimit         uint64
	GasUsed          uint64
	Hash             string
	LogsBloom        string
	Miner            string
	MixHash          string
	Nonce            string
	Number           uint64
	ParentHash       string
	ReceiptsRoot     string
	Sha3Uncles       string
	Size             uint64
	StateRoot        string
	Timestamp        uint64
	TotalDifficulty  uint64
	Transactions     []*Transaction
	TransactionsRoot string
	Uncles           interface{}
}

type Transaction struct {
	Hash             string
	Nonce            uint64
	BlockHash        string
	BlockNumber      uint64
	TransactionIndex uint64
	From             string
	To               string
	Value            float64
	Gas              uint64
	GasPrice         uint64
	Input            string
}
