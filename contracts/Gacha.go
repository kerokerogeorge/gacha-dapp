// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Bought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_total\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"account\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052683635c9adc5dea0000060005534801561001d57600080fd5b506111408061002d6000396000f3fe6080604052600436106100c25760003560e01c806373b9aa911161007f5780639dca362f116100595780639dca362f146102a4578063a6f2ae3a146102cf578063a9059cbb146102d9578063dd62ed3e14610316576100c2565b806373b9aa91146102115780637a27da791461024e57806395d89b4114610279576100c2565b806306fdde03146100c7578063095ea7b3146100f257806318160ddd1461012f57806323b872dd1461015a57806327e235e31461019757806370a08231146101d4575b600080fd5b3480156100d357600080fd5b506100dc610353565b6040516100e99190610d14565b60405180910390f35b3480156100fe57600080fd5b5061011960048036038101906101149190610dcf565b61038c565b6040516101269190610e2a565b60405180910390f35b34801561013b57600080fd5b5061014461047e565b6040516101519190610e54565b60405180910390f35b34801561016657600080fd5b50610181600480360381019061017c9190610e6f565b610487565b60405161018e9190610e2a565b60405180910390f35b3480156101a357600080fd5b506101be60048036038101906101b99190610ec2565b6107f2565b6040516101cb9190610e54565b60405180910390f35b3480156101e057600080fd5b506101fb60048036038101906101f69190610ec2565b61080a565b6040516102089190610e54565b60405180910390f35b34801561021d57600080fd5b5061023860048036038101906102339190610ec2565b610853565b6040516102459190610e2a565b60405180910390f35b34801561025a57600080fd5b50610263610873565b6040516102709190610e54565b60405180910390f35b34801561028557600080fd5b5061028e610879565b60405161029b9190610d14565b60405180910390f35b3480156102b057600080fd5b506102b96108b2565b6040516102c69190610e2a565b60405180910390f35b6102d7610910565b005b3480156102e557600080fd5b5061030060048036038101906102fb9190610dcf565b6109ee565b60405161030d9190610e2a565b60405180910390f35b34801561032257600080fd5b5061033d60048036038101906103389190610eef565b610bfd565b60405161034a9190610e54565b60405180910390f35b6040518060400160405280600a81526020017f474143484120636f696e0000000000000000000000000000000000000000000081525081565b600081600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161046c9190610e54565b60405180910390a36001905092915050565b60008054905090565b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548211156104d557600080fd5b600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482111561055e57600080fd5b81600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546105a99190610f5e565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546106749190610f5e565b600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461073f9190610f92565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516107df9190610e54565b60405180910390a3600190509392505050565b60016020528060005260406000206000915090505481565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60026020528060005260406000206000915054906101000a900460ff1681565b60005481565b6040518060400160405280600381526020017f474143000000000000000000000000000000000000000000000000000000000081525081565b60006001600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055905090565b600034905060006109203061080a565b905060008211610965576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095c90611012565b60405180910390fd5b808211156109a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099f9061107e565b60405180910390fd5b6109b233836109ee565b507f4e08ba899977cf7d4c2964bce71c6b9a7ef76ee5166a4c1249a1e08016e33ef1826040516109e29190610e54565b60405180910390a15050565b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054821115610a72576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a69906110ea565b60405180910390fd5b81600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610abd9190610f5e565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610b4b9190610f92565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610beb9190610e54565b60405180910390a36001905092915050565b6000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610cbe578082015181840152602081019050610ca3565b60008484015250505050565b6000601f19601f8301169050919050565b6000610ce682610c84565b610cf08185610c8f565b9350610d00818560208601610ca0565b610d0981610cca565b840191505092915050565b60006020820190508181036000830152610d2e8184610cdb565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610d6682610d3b565b9050919050565b610d7681610d5b565b8114610d8157600080fd5b50565b600081359050610d9381610d6d565b92915050565b6000819050919050565b610dac81610d99565b8114610db757600080fd5b50565b600081359050610dc981610da3565b92915050565b60008060408385031215610de657610de5610d36565b5b6000610df485828601610d84565b9250506020610e0585828601610dba565b9150509250929050565b60008115159050919050565b610e2481610e0f565b82525050565b6000602082019050610e3f6000830184610e1b565b92915050565b610e4e81610d99565b82525050565b6000602082019050610e696000830184610e45565b92915050565b600080600060608486031215610e8857610e87610d36565b5b6000610e9686828701610d84565b9350506020610ea786828701610d84565b9250506040610eb886828701610dba565b9150509250925092565b600060208284031215610ed857610ed7610d36565b5b6000610ee684828501610d84565b91505092915050565b60008060408385031215610f0657610f05610d36565b5b6000610f1485828601610d84565b9250506020610f2585828601610d84565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610f6982610d99565b9150610f7483610d99565b9250828203905081811115610f8c57610f8b610f2f565b5b92915050565b6000610f9d82610d99565b9150610fa883610d99565b9250828201905080821115610fc057610fbf610f2f565b5b92915050565b7f596f75206e65656420746f2073656e6420736f6d652065746865720000000000600082015250565b6000610ffc601b83610c8f565b915061100782610fc6565b602082019050919050565b6000602082019050818103600083015261102b81610fef565b9050919050565b7f4e6f7420656e6f75676820746f6b656e7320696e207468652072657365727665600082015250565b6000611068602083610c8f565b915061107382611032565b602082019050919050565b600060208201905081810360008301526110978161105b565b9050919050565b7f6e6f7420656e6f75676820746f6b656e7320746f2073656e6400000000000000600082015250565b60006110d4601983610c8f565b91506110df8261109e565b602082019050919050565b60006020820190508181036000830152611103816110c7565b905091905056fea2646970667358221220bf95ddb787356dfa46e39c8c2780a7b57548b7b9075b57f94c4307ef02d7bac464736f6c63430008110033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Total is a free data retrieval call binding the contract method 0x7a27da79.
//
// Solidity: function _total() view returns(uint256)
func (_Contract *ContractCaller) Total(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_total")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Total is a free data retrieval call binding the contract method 0x7a27da79.
//
// Solidity: function _total() view returns(uint256)
func (_Contract *ContractSession) Total() (*big.Int, error) {
	return _Contract.Contract.Total(&_Contract.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x7a27da79.
//
// Solidity: function _total() view returns(uint256)
func (_Contract *ContractCallerSession) Total() (*big.Int, error) {
	return _Contract.Contract.Total(&_Contract.CallOpts)
}

// Account is a free data retrieval call binding the contract method 0x73b9aa91.
//
// Solidity: function account(address ) view returns(bool)
func (_Contract *ContractCaller) Account(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "account", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Account is a free data retrieval call binding the contract method 0x73b9aa91.
//
// Solidity: function account(address ) view returns(bool)
func (_Contract *ContractSession) Account(arg0 common.Address) (bool, error) {
	return _Contract.Contract.Account(&_Contract.CallOpts, arg0)
}

// Account is a free data retrieval call binding the contract method 0x73b9aa91.
//
// Solidity: function account(address ) view returns(bool)
func (_Contract *ContractCallerSession) Account(arg0 common.Address) (bool, error) {
	return _Contract.Contract.Account(&_Contract.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_Contract *ContractCaller) Allowance(opts *bind.CallOpts, owner common.Address, delegate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allowance", owner, delegate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_Contract *ContractSession) Allowance(owner common.Address, delegate common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowance(&_Contract.CallOpts, owner, delegate)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_Contract *ContractCallerSession) Allowance(owner common.Address, delegate common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowance(&_Contract.CallOpts, owner, delegate)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_Contract *ContractCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balanceOf", tokenOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_Contract *ContractSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_Contract *ContractCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, tokenOwner)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Contract *ContractCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Contract *ContractSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Balances(&_Contract.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Contract *ContractCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Balances(&_Contract.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCallerSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCallerSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", delegate, numTokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_Contract *ContractSession) Approve(delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, delegate, numTokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_Contract *ContractTransactorSession) Approve(delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, delegate, numTokens)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_Contract *ContractTransactor) Buy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "buy")
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_Contract *ContractSession) Buy() (*types.Transaction, error) {
	return _Contract.Contract.Buy(&_Contract.TransactOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_Contract *ContractTransactorSession) Buy() (*types.Transaction, error) {
	return _Contract.Contract.Buy(&_Contract.TransactOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bool)
func (_Contract *ContractTransactor) CreateAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createAccount")
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bool)
func (_Contract *ContractSession) CreateAccount() (*types.Transaction, error) {
	return _Contract.Contract.CreateAccount(&_Contract.TransactOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bool)
func (_Contract *ContractTransactorSession) CreateAccount() (*types.Transaction, error) {
	return _Contract.Contract.CreateAccount(&_Contract.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_Contract *ContractTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transfer", receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_Contract *ContractSession) Transfer(receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Transfer(&_Contract.TransactOpts, receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_Contract *ContractTransactorSession) Transfer(receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Transfer(&_Contract.TransactOpts, receiver, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_Contract *ContractTransactor) TransferFrom(opts *bind.TransactOpts, owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom", owner, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_Contract *ContractSession) TransferFrom(owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, owner, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom(owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, owner, buyer, numTokens)
}

// ContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contract contract.
type ContractApprovalIterator struct {
	Event *ContractApproval // Event containing the contract specifics and raw log

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
func (it *ContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApproval)
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
		it.Event = new(ContractApproval)
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
func (it *ContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApproval represents a Approval event raised by the Contract contract.
type ContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contract *ContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalIterator{contract: _Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contract *ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApproval)
				if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Contract *ContractFilterer) ParseApproval(log types.Log) (*ContractApproval, error) {
	event := new(ContractApproval)
	if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBoughtIterator is returned from FilterBought and is used to iterate over the raw logs and unpacked data for Bought events raised by the Contract contract.
type ContractBoughtIterator struct {
	Event *ContractBought // Event containing the contract specifics and raw log

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
func (it *ContractBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBought)
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
		it.Event = new(ContractBought)
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
func (it *ContractBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBought represents a Bought event raised by the Contract contract.
type ContractBought struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBought is a free log retrieval operation binding the contract event 0x4e08ba899977cf7d4c2964bce71c6b9a7ef76ee5166a4c1249a1e08016e33ef1.
//
// Solidity: event Bought(uint256 amount)
func (_Contract *ContractFilterer) FilterBought(opts *bind.FilterOpts) (*ContractBoughtIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Bought")
	if err != nil {
		return nil, err
	}
	return &ContractBoughtIterator{contract: _Contract.contract, event: "Bought", logs: logs, sub: sub}, nil
}

// WatchBought is a free log subscription operation binding the contract event 0x4e08ba899977cf7d4c2964bce71c6b9a7ef76ee5166a4c1249a1e08016e33ef1.
//
// Solidity: event Bought(uint256 amount)
func (_Contract *ContractFilterer) WatchBought(opts *bind.WatchOpts, sink chan<- *ContractBought) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Bought")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBought)
				if err := _Contract.contract.UnpackLog(event, "Bought", log); err != nil {
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

// ParseBought is a log parse operation binding the contract event 0x4e08ba899977cf7d4c2964bce71c6b9a7ef76ee5166a4c1249a1e08016e33ef1.
//
// Solidity: event Bought(uint256 amount)
func (_Contract *ContractFilterer) ParseBought(log types.Log) (*ContractBought, error) {
	event := new(ContractBought)
	if err := _Contract.contract.UnpackLog(event, "Bought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contract contract.
type ContractTransferIterator struct {
	Event *ContractTransfer // Event containing the contract specifics and raw log

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
func (it *ContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransfer)
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
		it.Event = new(ContractTransfer)
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
func (it *ContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransfer represents a Transfer event raised by the Contract contract.
type ContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contract *ContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransferIterator{contract: _Contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contract *ContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransfer)
				if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Contract *ContractFilterer) ParseTransfer(log types.Log) (*ContractTransfer, error) {
	event := new(ContractTransfer)
	if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
