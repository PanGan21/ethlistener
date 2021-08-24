// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EventExampleABI is the input ABI used to generate the binding from.
const EventExampleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"data2\",\"type\":\"uint256\"}],\"name\":\"DataStored\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_data1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_data2\",\"type\":\"uint256\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EventExampleFuncSigs maps the 4-byte function signature to its string representation.
var EventExampleFuncSigs = map[string]string{
	"711c1c4a": "storeData(uint256,uint256)",
}

// EventExampleBin is the compiled bytecode used for deploying new contracts.
var EventExampleBin = "0x608060405234801561001057600080fd5b5060d88061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063711c1c4a14602d575b600080fd5b603c60383660046081565b603e565b005b6000829055600181905560405182815281907fd3b6c594451bea3204716c9e34653b521ace3900157ba74eabc1390c20bb8c989060200160405180910390a25050565b60008060408385031215609357600080fd5b5050803592602090910135915056fea26469706673582212208102091857fb3fc18f5cfb52ce67fa4bbf7889fda25350c58e06f61f011f547064736f6c63430008060033"

// DeployEventExample deploys a new Ethereum contract, binding an instance of EventExample to it.
func DeployEventExample(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EventExample, error) {
	parsed, err := abi.JSON(strings.NewReader(EventExampleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EventExampleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EventExample{EventExampleCaller: EventExampleCaller{contract: contract}, EventExampleTransactor: EventExampleTransactor{contract: contract}, EventExampleFilterer: EventExampleFilterer{contract: contract}}, nil
}

// EventExample is an auto generated Go binding around an Ethereum contract.
type EventExample struct {
	EventExampleCaller     // Read-only binding to the contract
	EventExampleTransactor // Write-only binding to the contract
	EventExampleFilterer   // Log filterer for contract events
}

// EventExampleCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventExampleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventExampleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventExampleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventExampleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventExampleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventExampleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventExampleSession struct {
	Contract     *EventExample     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventExampleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventExampleCallerSession struct {
	Contract *EventExampleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EventExampleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventExampleTransactorSession struct {
	Contract     *EventExampleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EventExampleRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventExampleRaw struct {
	Contract *EventExample // Generic contract binding to access the raw methods on
}

// EventExampleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventExampleCallerRaw struct {
	Contract *EventExampleCaller // Generic read-only contract binding to access the raw methods on
}

// EventExampleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventExampleTransactorRaw struct {
	Contract *EventExampleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventExample creates a new instance of EventExample, bound to a specific deployed contract.
func NewEventExample(address common.Address, backend bind.ContractBackend) (*EventExample, error) {
	contract, err := bindEventExample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EventExample{EventExampleCaller: EventExampleCaller{contract: contract}, EventExampleTransactor: EventExampleTransactor{contract: contract}, EventExampleFilterer: EventExampleFilterer{contract: contract}}, nil
}

// NewEventExampleCaller creates a new read-only instance of EventExample, bound to a specific deployed contract.
func NewEventExampleCaller(address common.Address, caller bind.ContractCaller) (*EventExampleCaller, error) {
	contract, err := bindEventExample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventExampleCaller{contract: contract}, nil
}

// NewEventExampleTransactor creates a new write-only instance of EventExample, bound to a specific deployed contract.
func NewEventExampleTransactor(address common.Address, transactor bind.ContractTransactor) (*EventExampleTransactor, error) {
	contract, err := bindEventExample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventExampleTransactor{contract: contract}, nil
}

// NewEventExampleFilterer creates a new log filterer instance of EventExample, bound to a specific deployed contract.
func NewEventExampleFilterer(address common.Address, filterer bind.ContractFilterer) (*EventExampleFilterer, error) {
	contract, err := bindEventExample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventExampleFilterer{contract: contract}, nil
}

// bindEventExample binds a generic wrapper to an already deployed contract.
func bindEventExample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EventExampleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventExample *EventExampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventExample.Contract.EventExampleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventExample *EventExampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventExample.Contract.EventExampleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventExample *EventExampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventExample.Contract.EventExampleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventExample *EventExampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventExample.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventExample *EventExampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventExample.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventExample *EventExampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventExample.Contract.contract.Transact(opts, method, params...)
}

// StoreData is a paid mutator transaction binding the contract method 0x711c1c4a.
//
// Solidity: function storeData(uint256 _data1, uint256 _data2) returns()
func (_EventExample *EventExampleTransactor) StoreData(opts *bind.TransactOpts, _data1 *big.Int, _data2 *big.Int) (*types.Transaction, error) {
	return _EventExample.contract.Transact(opts, "storeData", _data1, _data2)
}

// StoreData is a paid mutator transaction binding the contract method 0x711c1c4a.
//
// Solidity: function storeData(uint256 _data1, uint256 _data2) returns()
func (_EventExample *EventExampleSession) StoreData(_data1 *big.Int, _data2 *big.Int) (*types.Transaction, error) {
	return _EventExample.Contract.StoreData(&_EventExample.TransactOpts, _data1, _data2)
}

// StoreData is a paid mutator transaction binding the contract method 0x711c1c4a.
//
// Solidity: function storeData(uint256 _data1, uint256 _data2) returns()
func (_EventExample *EventExampleTransactorSession) StoreData(_data1 *big.Int, _data2 *big.Int) (*types.Transaction, error) {
	return _EventExample.Contract.StoreData(&_EventExample.TransactOpts, _data1, _data2)
}

// EventExampleDataStoredIterator is returned from FilterDataStored and is used to iterate over the raw logs and unpacked data for DataStored events raised by the EventExample contract.
type EventExampleDataStoredIterator struct {
	Event *EventExampleDataStored // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventExampleDataStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventExampleDataStored)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventExampleDataStored)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventExampleDataStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventExampleDataStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventExampleDataStored represents a DataStored event raised by the EventExample contract.
type EventExampleDataStored struct {
	Data1 *big.Int
	Data2 *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDataStored is a free log retrieval operation binding the contract event 0xd3b6c594451bea3204716c9e34653b521ace3900157ba74eabc1390c20bb8c98.
//
// Solidity: event DataStored(uint256 data1, uint256 indexed data2)
func (_EventExample *EventExampleFilterer) FilterDataStored(opts *bind.FilterOpts, data2 []*big.Int) (*EventExampleDataStoredIterator, error) {

	var data2Rule []interface{}
	for _, data2Item := range data2 {
		data2Rule = append(data2Rule, data2Item)
	}

	logs, sub, err := _EventExample.contract.FilterLogs(opts, "DataStored", data2Rule)
	if err != nil {
		return nil, err
	}
	return &EventExampleDataStoredIterator{contract: _EventExample.contract, event: "DataStored", logs: logs, sub: sub}, nil
}

// WatchDataStored is a free log subscription operation binding the contract event 0xd3b6c594451bea3204716c9e34653b521ace3900157ba74eabc1390c20bb8c98.
//
// Solidity: event DataStored(uint256 data1, uint256 indexed data2)
func (_EventExample *EventExampleFilterer) WatchDataStored(opts *bind.WatchOpts, sink chan<- *EventExampleDataStored, data2 []*big.Int) (event.Subscription, error) {

	var data2Rule []interface{}
	for _, data2Item := range data2 {
		data2Rule = append(data2Rule, data2Item)
	}

	logs, sub, err := _EventExample.contract.WatchLogs(opts, "DataStored", data2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventExampleDataStored)
				if err := _EventExample.contract.UnpackLog(event, "DataStored", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDataStored is a log parse operation binding the contract event 0xd3b6c594451bea3204716c9e34653b521ace3900157ba74eabc1390c20bb8c98.
//
// Solidity: event DataStored(uint256 data1, uint256 indexed data2)
func (_EventExample *EventExampleFilterer) ParseDataStored(log types.Log) (*EventExampleDataStored, error) {
	event := new(EventExampleDataStored)
	if err := _EventExample.contract.UnpackLog(event, "DataStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
