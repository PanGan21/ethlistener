package ethlistener

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type Contract struct {
	Name    string
	ABI     abi.ABI
	Address common.Address
	events  map[common.Hash]string
}

func NewContract(name string, abiJSON string, address common.Address) (*Contract, error) {
	contract := &Contract{
		Name:    name,
		Address: address,
		events:  make(map[common.Hash]string),
	}

	abi, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, err
	}
	contract.ABI = abi
	for _, event := range abi.Events {
		contract.events[event.ID] = event.Name
	}
	return contract, nil
}
