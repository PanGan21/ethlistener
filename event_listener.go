package ethlistener

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	bufferedSize = 1000
)

type EventListener struct {
	client                    ethclient.Client
	logChannel                chan types.Log
	contractAddressToContract map[common.Address]*Contract
}

func NewEventListener(client *ethclient.Client, contracts []*Contract) *EventListener {
	listener := &EventListener{
		client:                    *client,
		contractAddressToContract: make(map[common.Address]*Contract),
		logChannel:                make(chan types.Log, bufferedSize),
	}

	for _, contract := range contracts {
		listener.contractAddressToContract[contract.Address] = contract
	}

	return listener
}

func (el *EventListener) Listen(fromBlock *big.Int, eventChannel chan<- *ContractEvent, stopChannel <-chan struct{}) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	contracts := make([]common.Address, 0)
	for addr := range el.contractAddressToContract {
		contracts = append(contracts, addr)
	}
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		Addresses: contracts,
	}
	subscription, err := el.client.SubscribeFilterLogs(ctx, query, el.logChannel)
	if err != nil {
		return err
	}
	defer subscription.Unsubscribe()

	pastLogs, err := el.client.FilterLogs(context.Background(), query)
	if err != nil {
		return err
	}

	go func() {
		for _, pastLog := range pastLogs {
			el.logChannel <- pastLog
		}
	}()

	for {
		select {
		case err := <-subscription.Err():
			return err
		case log := <-el.logChannel:
			if receivedEvent := el.Parse(log); receivedEvent != nil {
				eventChannel <- receivedEvent
			}
		case <-stopChannel:
			return nil
		}
	}
}

func (el *EventListener) Parse(log types.Log) *ContractEvent {
	contract, ok := el.contractAddressToContract[log.Address]
	if !ok {
		return nil
	}
	if len(log.Topics) == 0 {
		return nil
	}
	name, ok := contract.events[log.Topics[0]]
	if !ok {
		return nil
	}
	return &ContractEvent{
		BlockNumber: log.BlockNumber,
		BlockHash:   log.BlockHash,
		TxHash:      log.TxHash,
		Contract:    contract,
		Name:        name,
		Data:        log.Data,
		Removed:     log.Removed,
	}
}
