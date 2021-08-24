package ethlistener

import (
	"github.com/ethereum/go-ethereum/common"
)

type ContractEvent struct {
	BlockNumber uint64
	BlockHash   common.Hash
	TxHash      common.Hash
	Contract    *Contract
	Name        string
	Data        []byte
	Removed     bool
}
