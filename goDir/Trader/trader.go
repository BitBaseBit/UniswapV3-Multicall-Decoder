// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trader

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

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Session) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20CallerSession) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Session) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20CallerSession) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Session) Symbol() (string, error) {
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20CallerSession) Symbol() (string, error) {
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV2PairABI is the input ABI used to generate the binding from.
const IUniswapV2PairABI = "[{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IUniswapV2PairFuncSigs maps the 4-byte function signature to its string representation.
var IUniswapV2PairFuncSigs = map[string]string{
	"0902f1ac": "getReserves()",
	"022c0d9f": "swap(uint256,uint256,address,bytes)",
}

// IUniswapV2Pair is an auto generated Go binding around an Ethereum contract.
type IUniswapV2Pair struct {
	IUniswapV2PairCaller     // Read-only binding to the contract
	IUniswapV2PairTransactor // Write-only binding to the contract
	IUniswapV2PairFilterer   // Log filterer for contract events
}

// IUniswapV2PairCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV2PairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2PairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV2PairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2PairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV2PairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2PairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV2PairSession struct {
	Contract     *IUniswapV2Pair   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUniswapV2PairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV2PairCallerSession struct {
	Contract *IUniswapV2PairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IUniswapV2PairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV2PairTransactorSession struct {
	Contract     *IUniswapV2PairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IUniswapV2PairRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV2PairRaw struct {
	Contract *IUniswapV2Pair // Generic contract binding to access the raw methods on
}

// IUniswapV2PairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV2PairCallerRaw struct {
	Contract *IUniswapV2PairCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV2PairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV2PairTransactorRaw struct {
	Contract *IUniswapV2PairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV2Pair creates a new instance of IUniswapV2Pair, bound to a specific deployed contract.
func NewIUniswapV2Pair(address common.Address, backend bind.ContractBackend) (*IUniswapV2Pair, error) {
	contract, err := bindIUniswapV2Pair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2Pair{IUniswapV2PairCaller: IUniswapV2PairCaller{contract: contract}, IUniswapV2PairTransactor: IUniswapV2PairTransactor{contract: contract}, IUniswapV2PairFilterer: IUniswapV2PairFilterer{contract: contract}}, nil
}

// NewIUniswapV2PairCaller creates a new read-only instance of IUniswapV2Pair, bound to a specific deployed contract.
func NewIUniswapV2PairCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV2PairCaller, error) {
	contract, err := bindIUniswapV2Pair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2PairCaller{contract: contract}, nil
}

// NewIUniswapV2PairTransactor creates a new write-only instance of IUniswapV2Pair, bound to a specific deployed contract.
func NewIUniswapV2PairTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV2PairTransactor, error) {
	contract, err := bindIUniswapV2Pair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2PairTransactor{contract: contract}, nil
}

// NewIUniswapV2PairFilterer creates a new log filterer instance of IUniswapV2Pair, bound to a specific deployed contract.
func NewIUniswapV2PairFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV2PairFilterer, error) {
	contract, err := bindIUniswapV2Pair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2PairFilterer{contract: contract}, nil
}

// bindIUniswapV2Pair binds a generic wrapper to an already deployed contract.
func bindIUniswapV2Pair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapV2PairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV2Pair *IUniswapV2PairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV2Pair.Contract.IUniswapV2PairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV2Pair *IUniswapV2PairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV2Pair.Contract.IUniswapV2PairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV2Pair *IUniswapV2PairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV2Pair.Contract.IUniswapV2PairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV2Pair *IUniswapV2PairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV2Pair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV2Pair *IUniswapV2PairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV2Pair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV2Pair *IUniswapV2PairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV2Pair.Contract.contract.Transact(opts, method, params...)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_IUniswapV2Pair *IUniswapV2PairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _IUniswapV2Pair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_IUniswapV2Pair *IUniswapV2PairSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _IUniswapV2Pair.Contract.GetReserves(&_IUniswapV2Pair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_IUniswapV2Pair *IUniswapV2PairCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _IUniswapV2Pair.Contract.GetReserves(&_IUniswapV2Pair.CallOpts)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_IUniswapV2Pair *IUniswapV2PairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IUniswapV2Pair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_IUniswapV2Pair *IUniswapV2PairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IUniswapV2Pair.Contract.Swap(&_IUniswapV2Pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_IUniswapV2Pair *IUniswapV2PairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IUniswapV2Pair.Contract.Swap(&_IUniswapV2Pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// IWETHABI is the input ABI used to generate the binding from.
const IWETHABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IWETHFuncSigs maps the 4-byte function signature to its string representation.
var IWETHFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"d0e30db0": "deposit()",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"2e1a7d4d": "withdraw(uint256)",
}

// IWETH is an auto generated Go binding around an Ethereum contract.
type IWETH struct {
	IWETHCaller     // Read-only binding to the contract
	IWETHTransactor // Write-only binding to the contract
	IWETHFilterer   // Log filterer for contract events
}

// IWETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWETHSession struct {
	Contract     *IWETH            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWETHCallerSession struct {
	Contract *IWETHCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IWETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWETHTransactorSession struct {
	Contract     *IWETHTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWETHRaw struct {
	Contract *IWETH // Generic contract binding to access the raw methods on
}

// IWETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWETHCallerRaw struct {
	Contract *IWETHCaller // Generic read-only contract binding to access the raw methods on
}

// IWETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWETHTransactorRaw struct {
	Contract *IWETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWETH creates a new instance of IWETH, bound to a specific deployed contract.
func NewIWETH(address common.Address, backend bind.ContractBackend) (*IWETH, error) {
	contract, err := bindIWETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWETH{IWETHCaller: IWETHCaller{contract: contract}, IWETHTransactor: IWETHTransactor{contract: contract}, IWETHFilterer: IWETHFilterer{contract: contract}}, nil
}

// NewIWETHCaller creates a new read-only instance of IWETH, bound to a specific deployed contract.
func NewIWETHCaller(address common.Address, caller bind.ContractCaller) (*IWETHCaller, error) {
	contract, err := bindIWETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWETHCaller{contract: contract}, nil
}

// NewIWETHTransactor creates a new write-only instance of IWETH, bound to a specific deployed contract.
func NewIWETHTransactor(address common.Address, transactor bind.ContractTransactor) (*IWETHTransactor, error) {
	contract, err := bindIWETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWETHTransactor{contract: contract}, nil
}

// NewIWETHFilterer creates a new log filterer instance of IWETH, bound to a specific deployed contract.
func NewIWETHFilterer(address common.Address, filterer bind.ContractFilterer) (*IWETHFilterer, error) {
	contract, err := bindIWETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWETHFilterer{contract: contract}, nil
}

// bindIWETH binds a generic wrapper to an already deployed contract.
func bindIWETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH *IWETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH.Contract.IWETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH *IWETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH.Contract.IWETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH *IWETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH.Contract.IWETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH *IWETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH *IWETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH *IWETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH *IWETHCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH *IWETHSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IWETH.Contract.Allowance(&_IWETH.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH *IWETHCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IWETH.Contract.Allowance(&_IWETH.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IWETH *IWETHCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IWETH *IWETHSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IWETH.Contract.BalanceOf(&_IWETH.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IWETH *IWETHCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IWETH.Contract.BalanceOf(&_IWETH.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IWETH *IWETHCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IWETH.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IWETH *IWETHSession) Decimals() (uint8, error) {
	return _IWETH.Contract.Decimals(&_IWETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IWETH *IWETHCallerSession) Decimals() (uint8, error) {
	return _IWETH.Contract.Decimals(&_IWETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IWETH *IWETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IWETH.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IWETH *IWETHSession) Name() (string, error) {
	return _IWETH.Contract.Name(&_IWETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IWETH *IWETHCallerSession) Name() (string, error) {
	return _IWETH.Contract.Name(&_IWETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IWETH *IWETHCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IWETH.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IWETH *IWETHSession) Symbol() (string, error) {
	return _IWETH.Contract.Symbol(&_IWETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IWETH *IWETHCallerSession) Symbol() (string, error) {
	return _IWETH.Contract.Symbol(&_IWETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH *IWETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWETH.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH *IWETHSession) TotalSupply() (*big.Int, error) {
	return _IWETH.Contract.TotalSupply(&_IWETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH *IWETHCallerSession) TotalSupply() (*big.Int, error) {
	return _IWETH.Contract.TotalSupply(&_IWETH.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IWETH *IWETHTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IWETH *IWETHSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Approve(&_IWETH.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IWETH *IWETHTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Approve(&_IWETH.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH *IWETHTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH *IWETHSession) Deposit() (*types.Transaction, error) {
	return _IWETH.Contract.Deposit(&_IWETH.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH *IWETHTransactorSession) Deposit() (*types.Transaction, error) {
	return _IWETH.Contract.Deposit(&_IWETH.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWETH *IWETHTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWETH *IWETHSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Transfer(&_IWETH.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWETH *IWETHTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Transfer(&_IWETH.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IWETH *IWETHTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IWETH *IWETHSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.TransferFrom(&_IWETH.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IWETH *IWETHTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.TransferFrom(&_IWETH.TransactOpts, from, to, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH *IWETHTransactor) Withdraw(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "withdraw", arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH *IWETHSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Withdraw(&_IWETH.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH *IWETHTransactorSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Withdraw(&_IWETH.TransactOpts, arg0)
}

// IWETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IWETH contract.
type IWETHApprovalIterator struct {
	Event *IWETHApproval // Event containing the contract specifics and raw log

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
func (it *IWETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWETHApproval)
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
		it.Event = new(IWETHApproval)
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
func (it *IWETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWETHApproval represents a Approval event raised by the IWETH contract.
type IWETHApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH *IWETHFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IWETHApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IWETH.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IWETHApprovalIterator{contract: _IWETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH *IWETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IWETHApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IWETH.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWETHApproval)
				if err := _IWETH.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH *IWETHFilterer) ParseApproval(log types.Log) (*IWETHApproval, error) {
	event := new(IWETHApproval)
	if err := _IWETH.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IWETH contract.
type IWETHTransferIterator struct {
	Event *IWETHTransfer // Event containing the contract specifics and raw log

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
func (it *IWETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWETHTransfer)
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
		it.Event = new(IWETHTransfer)
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
func (it *IWETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWETHTransfer represents a Transfer event raised by the IWETH contract.
type IWETHTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH *IWETHFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IWETHTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IWETH.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IWETHTransferIterator{contract: _IWETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH *IWETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IWETHTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IWETH.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWETHTransfer)
				if err := _IWETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH *IWETHFilterer) ParseTransfer(log types.Log) (*IWETHTransfer, error) {
	event := new(IWETHTransfer)
	if err := _IWETH.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniswapV2TraderABI is the input ABI used to generate the binding from.
const UniswapV2TraderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETHAddress\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"fromIsToken0\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pairAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"WETHAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveInMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"Buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"fromIsToken0\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pairAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reserveInMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"Sell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"convertETHtoWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"convertWETHtoETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// UniswapV2TraderFuncSigs maps the 4-byte function signature to its string representation.
var UniswapV2TraderFuncSigs = map[string]string{
	"02cd36ad": "Buy(bool,address,address,uint256,uint256,uint256)",
	"6849233c": "Sell(bool,address,address,uint256,uint256)",
	"2de9800a": "convertETHtoWETH(uint256)",
	"a2206833": "convertWETHtoETH(uint256)",
	"41c0e1b5": "kill()",
	"f14210a6": "withdrawETH(uint256)",
}

// UniswapV2TraderBin is the compiled bytecode used for deploying new contracts.
var UniswapV2TraderBin = "0x60c060405260405161138e38038061138e8339810160408190526100229161003f565b6001600160601b0319606091821b1660805233901b60a05261006d565b600060208284031215610050578081fd5b81516001600160a01b0381168114610066578182fd5b9392505050565b60805160601c60a05160601c6112a26100ec60003960008181610177015281816105ab015281816106bc015281816107fa0152818161082f01528181610be801528181610cad01528181610e1501528181610ec20152610f3401526000818161063b0152818161070e01528181610c3b0152610d0801526112a26000f3fe6080604052600436106100595760003560e01c806302cd36ad146100b75780632de9800a146100d757806341c0e1b5146100f75780636849233c1461010c578063a22068331461012c578063f14210a61461014c57610060565b3661006057005b34801561006c57600080fd5b5036156100b55760405162461bcd60e51b815260206004820152601260248201527111905313109050d2d7d5149251d1d154915160721b60448201526064015b60405180910390fd5b005b3480156100c357600080fd5b506100b56100d2366004611019565b61016c565b3480156100e357600080fd5b506100b56100f23660046110c5565b6105a0565b34801561010357600080fd5b506100b56106b1565b34801561011857600080fd5b506100b5610127366004610fc5565b610824565b34801561013857600080fd5b506100b56101473660046110c5565b610bdd565b34801561015857600080fd5b506100b56101673660046110c5565b610ca2565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101b45760405162461bcd60e51b81526004016100ac90611148565b80428110156101ef5760405162461bcd60e51b81526020600482015260076024820152661156141254915160ca1b60448201526064016100ac565b6000859050600080826001600160a01b0316630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b15801561023057600080fd5b505afa158015610244573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102689190611077565b506001600160701b03918216935016905060006102876103e5896111f4565b90508a1561041357868310156102af5760405162461bcd60e51b81526004016100ac90611111565b604080516001600160a01b038b8116602483015260448083018c905283518084039091018152606490920183526020820180516001600160e01b031663a9059cbb60e01b1790529151918c169161030691906110f5565b6000604051808303816000865af19150503d8060008114610343576040519150601f19603f3d011682016040523d82523d6000602084013e610348565b606091505b5050506000816103e88561035c91906111f4565b61036691906111bc565b61037084846111f4565b61037a91906111d4565b90506001600160a01b03851663022c0d9f60008330825b6040519080825280601f01601f1916602001820160405280156103bb576020820181803683370190505b506040518563ffffffff1660e01b81526004016103db949392919061116c565b600060405180830381600087803b1580156103f557600080fd5b505af1158015610409573d6000803e3d6000fd5b5050505050610593565b868210156104335760405162461bcd60e51b81526004016100ac90611111565b604080516001600160a01b038b8116602483015260448083018c905283518084039091018152606490920183526020820180516001600160e01b031663a9059cbb60e01b1790529151918c169161048a91906110f5565b6000604051808303816000865af19150503d80600081146104c7576040519150601f19603f3d011682016040523d82523d6000602084013e6104cc565b606091505b5050506000816103e8846104e091906111f4565b6104ea91906111bc565b6104f485846111f4565b6104fe91906111d4565b90506001600160a01b03851663022c0d9f82600030815b6040519080825280601f01601f19166020018201604052801561053f576020820181803683370190505b506040518563ffffffff1660e01b815260040161055f949392919061116c565b600060405180830381600087803b15801561057957600080fd5b505af115801561058d573d6000803e3d6000fd5b50505050505b5050505050505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105e85760405162461bcd60e51b81526004016100ac90611148565b80478111156106395760405162461bcd60e51b815260206004820152601860248201527f494e53554646494349454e545f4554485f42414c414e4345000000000000000060448201526064016100ac565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db0836040518263ffffffff1660e01b81526004016000604051808303818588803b15801561069457600080fd5b505af11580156106a8573d6000803e3d6000fd5b50505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146106f95760405162461bcd60e51b81526004016100ac90611148565b6040516370a0823160e01b81523060048201527f0000000000000000000000000000000000000000000000000000000000000000906000906001600160a01b038316906370a082319060240160206040518083038186803b15801561075d57600080fd5b505afa158015610771573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079591906110dd565b905080156107f857604051632e1a7d4d60e01b8152600481018290526001600160a01b03831690632e1a7d4d90602401600060405180830381600087803b1580156107df57600080fd5b505af11580156107f3573d6000803e3d6000fd5b505050505b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316ff5b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461086c5760405162461bcd60e51b81526004016100ac90611148565b80428110156108a75760405162461bcd60e51b81526020600482015260076024820152661156141254915160ca1b60448201526064016100ac565b6000849050600080826001600160a01b0316630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b1580156108e857600080fd5b505afa1580156108fc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109209190611077565b506040516370a0823160e01b81523060048201526001600160701b039283169450911691506000906001906001600160a01b038b16906370a082319060240160206040518083038186803b15801561097757600080fd5b505afa15801561098b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109af91906110dd565b6109b99190611213565b905060006109c96103e5836111f4565b90508a15610ad757878410156109f15760405162461bcd60e51b81526004016100ac90611111565b604080516001600160a01b038b81166024830152604480830186905283518084039091018152606490920183526020820180516001600160e01b031663a9059cbb60e01b1790529151918c1691610a4891906110f5565b6000604051808303816000865af19150503d8060008114610a85576040519150601f19603f3d011682016040523d82523d6000602084013e610a8a565b606091505b5050506000816103e886610a9e91906111f4565b610aa891906111bc565b610ab285846111f4565b610abc91906111d4565b90506001600160a01b03861663022c0d9f6000833082610391565b87831015610af75760405162461bcd60e51b81526004016100ac90611111565b604080516001600160a01b038b81166024830152604480830186905283518084039091018152606490920183526020820180516001600160e01b031663a9059cbb60e01b1790529151918c1691610b4e91906110f5565b6000604051808303816000865af19150503d8060008114610b8b576040519150601f19603f3d011682016040523d82523d6000602084013e610b90565b606091505b5050506000816103e885610ba491906111f4565b610bae91906111bc565b610bb886846111f4565b610bc291906111d4565b90506001600160a01b03861663022c0d9f8260003081610515565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610c255760405162461bcd60e51b81526004016100ac90611148565b604051632e1a7d4d60e01b8152600481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632e1a7d4d90602401600060405180830381600087803b158015610c8757600080fd5b505af1158015610c9b573d6000803e3d6000fd5b5050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610cea5760405162461bcd60e51b81526004016100ac90611148565b4780821115610f27576040516370a0823160e01b81523060048201527f0000000000000000000000000000000000000000000000000000000000000000906000906001600160a01b038316906370a082319060240160206040518083038186803b158015610d5757600080fd5b505afa158015610d6b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d8f91906110dd565b905083610d9c84836111bc565b10610e66576001600160a01b038216632e1a7d4d610dba8587611213565b6040518263ffffffff1660e01b8152600401610dd891815260200190565b600060405180830381600087803b158015610df257600080fd5b505af1158015610e06573d6000803e3d6000fd5b50506040516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925086156108fc02915086906000818181858888f19350505050158015610e60573d6000803e3d6000fd5b50610f21565b604051632e1a7d4d60e01b8152600481018290526001600160a01b03831690632e1a7d4d90602401600060405180830381600087803b158015610ea857600080fd5b505af1158015610ebc573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166108fc8285610ef991906111bc565b6040518115909202916000818181858888f19350505050158015610c9b573d6000803e3d6000fd5b50505050565b6040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169083156108fc029084906000818181858888f19350505050158015610f7d573d6000803e3d6000fd5b505050565b80356001600160a01b0381168114610f9957600080fd5b919050565b80358015158114610f9957600080fd5b80516001600160701b0381168114610f9957600080fd5b600080600080600060a08688031215610fdc578081fd5b610fe586610f9e565b9450610ff360208701610f82565b935061100160408701610f82565b94979396509394606081013594506080013592915050565b60008060008060008060c08789031215611031578081fd5b61103a87610f9e565b955061104860208801610f82565b945061105660408801610f82565b9350606087013592506080870135915060a087013590509295509295509295565b60008060006060848603121561108b578283fd5b61109484610fae565b92506110a260208501610fae565b9150604084015163ffffffff811681146110ba578182fd5b809150509250925092565b6000602082840312156110d6578081fd5b5035919050565b6000602082840312156110ee578081fd5b5051919050565b6000825161110781846020870161122a565b9190910192915050565b60208082526019908201527f494e53554646494349454e545f494e5055545f414d4f554e5400000000000000604082015260600190565b6020808252600a908201526927a7262cafa7aba722a960b11b604082015260600190565b84815283602082015260018060a01b038316604082015260806060820152600082518060808401526111a58160a085016020870161122a565b601f01601f19169190910160a00195945050505050565b600082198211156111cf576111cf611256565b500190565b6000826111ef57634e487b7160e01b81526012600452602481fd5b500490565b600081600019048311821515161561120e5761120e611256565b500290565b60008282101561122557611225611256565b500390565b60005b8381101561124557818101518382015260200161122d565b83811115610f215750506000910152565b634e487b7160e01b600052601160045260246000fdfea2646970667358221220b940937cbd3d47fd0ea83434a86c84f6a8593b45cda23b17de73a284ccc3e50264736f6c63430008040033"

// DeployUniswapV2Trader deploys a new Ethereum contract, binding an instance of UniswapV2Trader to it.
func DeployUniswapV2Trader(auth *bind.TransactOpts, backend bind.ContractBackend, _WETHAddress common.Address) (common.Address, *types.Transaction, *UniswapV2Trader, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapV2TraderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UniswapV2TraderBin), backend, _WETHAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniswapV2Trader{UniswapV2TraderCaller: UniswapV2TraderCaller{contract: contract}, UniswapV2TraderTransactor: UniswapV2TraderTransactor{contract: contract}, UniswapV2TraderFilterer: UniswapV2TraderFilterer{contract: contract}}, nil
}

// UniswapV2Trader is an auto generated Go binding around an Ethereum contract.
type UniswapV2Trader struct {
	UniswapV2TraderCaller     // Read-only binding to the contract
	UniswapV2TraderTransactor // Write-only binding to the contract
	UniswapV2TraderFilterer   // Log filterer for contract events
}

// UniswapV2TraderCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV2TraderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2TraderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV2TraderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2TraderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV2TraderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2TraderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV2TraderSession struct {
	Contract     *UniswapV2Trader  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV2TraderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV2TraderCallerSession struct {
	Contract *UniswapV2TraderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UniswapV2TraderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV2TraderTransactorSession struct {
	Contract     *UniswapV2TraderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UniswapV2TraderRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV2TraderRaw struct {
	Contract *UniswapV2Trader // Generic contract binding to access the raw methods on
}

// UniswapV2TraderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV2TraderCallerRaw struct {
	Contract *UniswapV2TraderCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV2TraderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV2TraderTransactorRaw struct {
	Contract *UniswapV2TraderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV2Trader creates a new instance of UniswapV2Trader, bound to a specific deployed contract.
func NewUniswapV2Trader(address common.Address, backend bind.ContractBackend) (*UniswapV2Trader, error) {
	contract, err := bindUniswapV2Trader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV2Trader{UniswapV2TraderCaller: UniswapV2TraderCaller{contract: contract}, UniswapV2TraderTransactor: UniswapV2TraderTransactor{contract: contract}, UniswapV2TraderFilterer: UniswapV2TraderFilterer{contract: contract}}, nil
}

// NewUniswapV2TraderCaller creates a new read-only instance of UniswapV2Trader, bound to a specific deployed contract.
func NewUniswapV2TraderCaller(address common.Address, caller bind.ContractCaller) (*UniswapV2TraderCaller, error) {
	contract, err := bindUniswapV2Trader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2TraderCaller{contract: contract}, nil
}

// NewUniswapV2TraderTransactor creates a new write-only instance of UniswapV2Trader, bound to a specific deployed contract.
func NewUniswapV2TraderTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV2TraderTransactor, error) {
	contract, err := bindUniswapV2Trader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2TraderTransactor{contract: contract}, nil
}

// NewUniswapV2TraderFilterer creates a new log filterer instance of UniswapV2Trader, bound to a specific deployed contract.
func NewUniswapV2TraderFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV2TraderFilterer, error) {
	contract, err := bindUniswapV2Trader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV2TraderFilterer{contract: contract}, nil
}

// bindUniswapV2Trader binds a generic wrapper to an already deployed contract.
func bindUniswapV2Trader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapV2TraderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2Trader *UniswapV2TraderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2Trader.Contract.UniswapV2TraderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2Trader *UniswapV2TraderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.UniswapV2TraderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2Trader *UniswapV2TraderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.UniswapV2TraderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2Trader *UniswapV2TraderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2Trader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2Trader *UniswapV2TraderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2Trader *UniswapV2TraderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.contract.Transact(opts, method, params...)
}

// Buy is a paid mutator transaction binding the contract method 0x02cd36ad.
//
// Solidity: function Buy(bool fromIsToken0, address from, address pairAddress, uint256 WETHAmount, uint256 reserveInMin, uint256 deadline) returns()
func (_UniswapV2Trader *UniswapV2TraderTransactor) Buy(opts *bind.TransactOpts, fromIsToken0 bool, from common.Address, pairAddress common.Address, WETHAmount *big.Int, reserveInMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trader.contract.Transact(opts, "Buy", fromIsToken0, from, pairAddress, WETHAmount, reserveInMin, deadline)
}

// Buy is a paid mutator transaction binding the contract method 0x02cd36ad.
//
// Solidity: function Buy(bool fromIsToken0, address from, address pairAddress, uint256 WETHAmount, uint256 reserveInMin, uint256 deadline) returns()
func (_UniswapV2Trader *UniswapV2TraderSession) Buy(fromIsToken0 bool, from common.Address, pairAddress common.Address, WETHAmount *big.Int, reserveInMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.Buy(&_UniswapV2Trader.TransactOpts, fromIsToken0, from, pairAddress, WETHAmount, reserveInMin, deadline)
}

// Buy is a paid mutator transaction binding the contract method 0x02cd36ad.
//
// Solidity: function Buy(bool fromIsToken0, address from, address pairAddress, uint256 WETHAmount, uint256 reserveInMin, uint256 deadline) returns()
func (_UniswapV2Trader *UniswapV2TraderTransactorSession) Buy(fromIsToken0 bool, from common.Address, pairAddress common.Address, WETHAmount *big.Int, reserveInMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.Buy(&_UniswapV2Trader.TransactOpts, fromIsToken0, from, pairAddress, WETHAmount, reserveInMin, deadline)
}

// Sell is a paid mutator transaction binding the contract method 0x6849233c.
//
// Solidity: function Sell(bool fromIsToken0, address from, address pairAddress, uint256 reserveInMin, uint256 deadline) returns()
func (_UniswapV2Trader *UniswapV2TraderTransactor) Sell(opts *bind.TransactOpts, fromIsToken0 bool, from common.Address, pairAddress common.Address, reserveInMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trader.contract.Transact(opts, "Sell", fromIsToken0, from, pairAddress, reserveInMin, deadline)
}

// Sell is a paid mutator transaction binding the contract method 0x6849233c.
//
// Solidity: function Sell(bool fromIsToken0, address from, address pairAddress, uint256 reserveInMin, uint256 deadline) returns()
func (_UniswapV2Trader *UniswapV2TraderSession) Sell(fromIsToken0 bool, from common.Address, pairAddress common.Address, reserveInMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.Sell(&_UniswapV2Trader.TransactOpts, fromIsToken0, from, pairAddress, reserveInMin, deadline)
}

// Sell is a paid mutator transaction binding the contract method 0x6849233c.
//
// Solidity: function Sell(bool fromIsToken0, address from, address pairAddress, uint256 reserveInMin, uint256 deadline) returns()
func (_UniswapV2Trader *UniswapV2TraderTransactorSession) Sell(fromIsToken0 bool, from common.Address, pairAddress common.Address, reserveInMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.Sell(&_UniswapV2Trader.TransactOpts, fromIsToken0, from, pairAddress, reserveInMin, deadline)
}

// ConvertETHtoWETH is a paid mutator transaction binding the contract method 0x2de9800a.
//
// Solidity: function convertETHtoWETH(uint256 amount) returns()
func (_UniswapV2Trader *UniswapV2TraderTransactor) ConvertETHtoWETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trader.contract.Transact(opts, "convertETHtoWETH", amount)
}

// ConvertETHtoWETH is a paid mutator transaction binding the contract method 0x2de9800a.
//
// Solidity: function convertETHtoWETH(uint256 amount) returns()
func (_UniswapV2Trader *UniswapV2TraderSession) ConvertETHtoWETH(amount *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.ConvertETHtoWETH(&_UniswapV2Trader.TransactOpts, amount)
}

// ConvertETHtoWETH is a paid mutator transaction binding the contract method 0x2de9800a.
//
// Solidity: function convertETHtoWETH(uint256 amount) returns()
func (_UniswapV2Trader *UniswapV2TraderTransactorSession) ConvertETHtoWETH(amount *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.ConvertETHtoWETH(&_UniswapV2Trader.TransactOpts, amount)
}

// ConvertWETHtoETH is a paid mutator transaction binding the contract method 0xa2206833.
//
// Solidity: function convertWETHtoETH(uint256 amount) returns()
func (_UniswapV2Trader *UniswapV2TraderTransactor) ConvertWETHtoETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trader.contract.Transact(opts, "convertWETHtoETH", amount)
}

// ConvertWETHtoETH is a paid mutator transaction binding the contract method 0xa2206833.
//
// Solidity: function convertWETHtoETH(uint256 amount) returns()
func (_UniswapV2Trader *UniswapV2TraderSession) ConvertWETHtoETH(amount *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.ConvertWETHtoETH(&_UniswapV2Trader.TransactOpts, amount)
}

// ConvertWETHtoETH is a paid mutator transaction binding the contract method 0xa2206833.
//
// Solidity: function convertWETHtoETH(uint256 amount) returns()
func (_UniswapV2Trader *UniswapV2TraderTransactorSession) ConvertWETHtoETH(amount *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.ConvertWETHtoETH(&_UniswapV2Trader.TransactOpts, amount)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_UniswapV2Trader *UniswapV2TraderTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Trader.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_UniswapV2Trader *UniswapV2TraderSession) Kill() (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.Kill(&_UniswapV2Trader.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_UniswapV2Trader *UniswapV2TraderTransactorSession) Kill() (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.Kill(&_UniswapV2Trader.TransactOpts)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_UniswapV2Trader *UniswapV2TraderTransactor) WithdrawETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trader.contract.Transact(opts, "withdrawETH", amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_UniswapV2Trader *UniswapV2TraderSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.WithdrawETH(&_UniswapV2Trader.TransactOpts, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_UniswapV2Trader *UniswapV2TraderTransactorSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.WithdrawETH(&_UniswapV2Trader.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_UniswapV2Trader *UniswapV2TraderTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _UniswapV2Trader.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_UniswapV2Trader *UniswapV2TraderSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.Fallback(&_UniswapV2Trader.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_UniswapV2Trader *UniswapV2TraderTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.Fallback(&_UniswapV2Trader.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV2Trader *UniswapV2TraderTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Trader.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV2Trader *UniswapV2TraderSession) Receive() (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.Receive(&_UniswapV2Trader.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV2Trader *UniswapV2TraderTransactorSession) Receive() (*types.Transaction, error) {
	return _UniswapV2Trader.Contract.Receive(&_UniswapV2Trader.TransactOpts)
}
