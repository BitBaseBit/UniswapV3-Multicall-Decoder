package tradeExecution

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

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
