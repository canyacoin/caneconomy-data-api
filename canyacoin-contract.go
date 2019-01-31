// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CanYaCoinABI is the input ABI used to generate the binding from.
const CanYaCoinABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// CanYaCoinBin is the compiled bytecode used for deploying new contracts.
const CanYaCoinBin = `0x608060405234801561001057600080fd5b50336000908152602081905260409020655af3107a400090556105be806100386000396000f3006080604052600436106100b95763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100be578063095ea7b31461014857806318160ddd1461018057806323b872dd146101a757806327e235e3146101d1578063313ce567146101f25780635c6581651461020757806370a082311461022e5780637e1c0c091461024f57806395d89b4114610264578063a9059cbb14610279578063dd62ed3e1461029d575b600080fd5b3480156100ca57600080fd5b506100d36102c4565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561010d5781810151838201526020016100f5565b50505050905090810190601f16801561013a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015457600080fd5b5061016c600160a060020a03600435166024356102fb565b604080519115158252519081900360200190f35b34801561018c57600080fd5b50610195610362565b60408051918252519081900360200190f35b3480156101b357600080fd5b5061016c600160a060020a036004358116906024351660443561036c565b3480156101dd57600080fd5b50610195600160a060020a036004351661044d565b3480156101fe57600080fd5b5061019561045f565b34801561021357600080fd5b50610195600160a060020a0360043581169060243516610464565b34801561023a57600080fd5b50610195600160a060020a0360043516610481565b34801561025b57600080fd5b5061019561049c565b34801561027057600080fd5b506100d36104a6565b34801561028557600080fd5b5061016c600160a060020a03600435166024356104dd565b3480156102a957600080fd5b50610195600160a060020a0360043581169060243516610567565b60408051808201909152600981527f43616e5961436f696e0000000000000000000000000000000000000000000000602082015281565b336000818152600160209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35060015b92915050565b655af3107a400090565b600160a060020a03831660009081526020819052604081205482118015906103b75750600160a060020a03841660009081526001602090815260408083203384529091529020548211155b1561044257600160a060020a038085166000818152602081815260408083208054889003905560018252808320338452825280832080548890039055938716808352828252918490208054870190558351868152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a3506001610446565b5060005b9392505050565b60006020819052908152604090205481565b600681565b600160209081526000928352604080842090915290825290205481565b600160a060020a031660009081526020819052604090205490565b655af3107a400081565b60408051808201909152600381527f43414e0000000000000000000000000000000000000000000000000000000000602082015281565b33600090815260208190526040812054821161055e573360008181526020818152604080832080548790039055600160a060020a03871680845292819020805487019055805186815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a350600161035c565b50600092915050565b600160a060020a039182166000908152600160209081526040808320939094168252919091522054905600a165627a7a723058203f265399683b1a9f1f77440bf597ca68a949ad597a9ee0288e2c28b18e0e39400029`

// DeployCanYaCoin deploys a new Ethereum contract, binding an instance of CanYaCoin to it.
func DeployCanYaCoin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CanYaCoin, error) {
	parsed, err := abi.JSON(strings.NewReader(CanYaCoinABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CanYaCoinBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CanYaCoin{CanYaCoinCaller: CanYaCoinCaller{contract: contract}, CanYaCoinTransactor: CanYaCoinTransactor{contract: contract}, CanYaCoinFilterer: CanYaCoinFilterer{contract: contract}}, nil
}

// CanYaCoin is an auto generated Go binding around an Ethereum contract.
type CanYaCoin struct {
	CanYaCoinCaller     // Read-only binding to the contract
	CanYaCoinTransactor // Write-only binding to the contract
	CanYaCoinFilterer   // Log filterer for contract events
}

// CanYaCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type CanYaCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanYaCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CanYaCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanYaCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CanYaCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanYaCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CanYaCoinSession struct {
	Contract     *CanYaCoin        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CanYaCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CanYaCoinCallerSession struct {
	Contract *CanYaCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CanYaCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CanYaCoinTransactorSession struct {
	Contract     *CanYaCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CanYaCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type CanYaCoinRaw struct {
	Contract *CanYaCoin // Generic contract binding to access the raw methods on
}

// CanYaCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CanYaCoinCallerRaw struct {
	Contract *CanYaCoinCaller // Generic read-only contract binding to access the raw methods on
}

// CanYaCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CanYaCoinTransactorRaw struct {
	Contract *CanYaCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCanYaCoin creates a new instance of CanYaCoin, bound to a specific deployed contract.
func NewCanYaCoin(address common.Address, backend bind.ContractBackend) (*CanYaCoin, error) {
	contract, err := bindCanYaCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CanYaCoin{CanYaCoinCaller: CanYaCoinCaller{contract: contract}, CanYaCoinTransactor: CanYaCoinTransactor{contract: contract}, CanYaCoinFilterer: CanYaCoinFilterer{contract: contract}}, nil
}

// NewCanYaCoinCaller creates a new read-only instance of CanYaCoin, bound to a specific deployed contract.
func NewCanYaCoinCaller(address common.Address, caller bind.ContractCaller) (*CanYaCoinCaller, error) {
	contract, err := bindCanYaCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CanYaCoinCaller{contract: contract}, nil
}

// NewCanYaCoinTransactor creates a new write-only instance of CanYaCoin, bound to a specific deployed contract.
func NewCanYaCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*CanYaCoinTransactor, error) {
	contract, err := bindCanYaCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CanYaCoinTransactor{contract: contract}, nil
}

// NewCanYaCoinFilterer creates a new log filterer instance of CanYaCoin, bound to a specific deployed contract.
func NewCanYaCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*CanYaCoinFilterer, error) {
	contract, err := bindCanYaCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CanYaCoinFilterer{contract: contract}, nil
}

// bindCanYaCoin binds a generic wrapper to an already deployed contract.
func bindCanYaCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CanYaCoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanYaCoin *CanYaCoinRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CanYaCoin.Contract.CanYaCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanYaCoin *CanYaCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanYaCoin.Contract.CanYaCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanYaCoin *CanYaCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanYaCoin.Contract.CanYaCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanYaCoin *CanYaCoinCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CanYaCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanYaCoin *CanYaCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanYaCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanYaCoin *CanYaCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanYaCoin.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_CanYaCoin *CanYaCoinCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CanYaCoin.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_CanYaCoin *CanYaCoinSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _CanYaCoin.Contract.Allowance(&_CanYaCoin.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_CanYaCoin *CanYaCoinCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _CanYaCoin.Contract.Allowance(&_CanYaCoin.CallOpts, _owner, _spender)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_CanYaCoin *CanYaCoinCaller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CanYaCoin.contract.Call(opts, out, "allowed", arg0, arg1)
	return *ret0, err
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_CanYaCoin *CanYaCoinSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CanYaCoin.Contract.Allowed(&_CanYaCoin.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_CanYaCoin *CanYaCoinCallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CanYaCoin.Contract.Allowed(&_CanYaCoin.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_CanYaCoin *CanYaCoinCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CanYaCoin.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_CanYaCoin *CanYaCoinSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _CanYaCoin.Contract.BalanceOf(&_CanYaCoin.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_CanYaCoin *CanYaCoinCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _CanYaCoin.Contract.BalanceOf(&_CanYaCoin.CallOpts, _owner)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_CanYaCoin *CanYaCoinCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CanYaCoin.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_CanYaCoin *CanYaCoinSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _CanYaCoin.Contract.Balances(&_CanYaCoin.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_CanYaCoin *CanYaCoinCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _CanYaCoin.Contract.Balances(&_CanYaCoin.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_CanYaCoin *CanYaCoinCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CanYaCoin.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_CanYaCoin *CanYaCoinSession) Decimals() (*big.Int, error) {
	return _CanYaCoin.Contract.Decimals(&_CanYaCoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_CanYaCoin *CanYaCoinCallerSession) Decimals() (*big.Int, error) {
	return _CanYaCoin.Contract.Decimals(&_CanYaCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CanYaCoin *CanYaCoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CanYaCoin.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CanYaCoin *CanYaCoinSession) Name() (string, error) {
	return _CanYaCoin.Contract.Name(&_CanYaCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CanYaCoin *CanYaCoinCallerSession) Name() (string, error) {
	return _CanYaCoin.Contract.Name(&_CanYaCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CanYaCoin *CanYaCoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CanYaCoin.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CanYaCoin *CanYaCoinSession) Symbol() (string, error) {
	return _CanYaCoin.Contract.Symbol(&_CanYaCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CanYaCoin *CanYaCoinCallerSession) Symbol() (string, error) {
	return _CanYaCoin.Contract.Symbol(&_CanYaCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CanYaCoin *CanYaCoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CanYaCoin.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CanYaCoin *CanYaCoinSession) TotalSupply() (*big.Int, error) {
	return _CanYaCoin.Contract.TotalSupply(&_CanYaCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CanYaCoin *CanYaCoinCallerSession) TotalSupply() (*big.Int, error) {
	return _CanYaCoin.Contract.TotalSupply(&_CanYaCoin.CallOpts)
}

// TotalTokens is a free data retrieval call binding the contract method 0x7e1c0c09.
//
// Solidity: function totalTokens() constant returns(uint256)
func (_CanYaCoin *CanYaCoinCaller) TotalTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CanYaCoin.contract.Call(opts, out, "totalTokens")
	return *ret0, err
}

// TotalTokens is a free data retrieval call binding the contract method 0x7e1c0c09.
//
// Solidity: function totalTokens() constant returns(uint256)
func (_CanYaCoin *CanYaCoinSession) TotalTokens() (*big.Int, error) {
	return _CanYaCoin.Contract.TotalTokens(&_CanYaCoin.CallOpts)
}

// TotalTokens is a free data retrieval call binding the contract method 0x7e1c0c09.
//
// Solidity: function totalTokens() constant returns(uint256)
func (_CanYaCoin *CanYaCoinCallerSession) TotalTokens() (*big.Int, error) {
	return _CanYaCoin.Contract.TotalTokens(&_CanYaCoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_CanYaCoin *CanYaCoinTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CanYaCoin.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_CanYaCoin *CanYaCoinSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CanYaCoin.Contract.Approve(&_CanYaCoin.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_CanYaCoin *CanYaCoinTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CanYaCoin.Contract.Approve(&_CanYaCoin.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_CanYaCoin *CanYaCoinTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CanYaCoin.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_CanYaCoin *CanYaCoinSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CanYaCoin.Contract.Transfer(&_CanYaCoin.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_CanYaCoin *CanYaCoinTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CanYaCoin.Contract.Transfer(&_CanYaCoin.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_CanYaCoin *CanYaCoinTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CanYaCoin.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_CanYaCoin *CanYaCoinSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CanYaCoin.Contract.TransferFrom(&_CanYaCoin.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_CanYaCoin *CanYaCoinTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CanYaCoin.Contract.TransferFrom(&_CanYaCoin.TransactOpts, _from, _to, _value)
}

// CanYaCoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CanYaCoin contract.
type CanYaCoinApprovalIterator struct {
	Event *CanYaCoinApproval // Event containing the contract specifics and raw log

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
func (it *CanYaCoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanYaCoinApproval)
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
		it.Event = new(CanYaCoinApproval)
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
func (it *CanYaCoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanYaCoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanYaCoinApproval represents a Approval event raised by the CanYaCoin contract.
type CanYaCoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_CanYaCoin *CanYaCoinFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CanYaCoinApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CanYaCoin.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CanYaCoinApprovalIterator{contract: _CanYaCoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_CanYaCoin *CanYaCoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CanYaCoinApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CanYaCoin.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanYaCoinApproval)
				if err := _CanYaCoin.contract.UnpackLog(event, "Approval", log); err != nil {
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

// CanYaCoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CanYaCoin contract.
type CanYaCoinTransferIterator struct {
	Event *CanYaCoinTransfer // Event containing the contract specifics and raw log

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
func (it *CanYaCoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanYaCoinTransfer)
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
		it.Event = new(CanYaCoinTransfer)
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
func (it *CanYaCoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanYaCoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanYaCoinTransfer represents a Transfer event raised by the CanYaCoin contract.
type CanYaCoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_CanYaCoin *CanYaCoinFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CanYaCoinTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CanYaCoin.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CanYaCoinTransferIterator{contract: _CanYaCoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_CanYaCoin *CanYaCoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CanYaCoinTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CanYaCoin.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanYaCoinTransfer)
				if err := _CanYaCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC20TokenInterfaceABI is the input ABI used to generate the binding from.
const ERC20TokenInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"supply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20TokenInterfaceBin is the compiled bytecode used for deploying new contracts.
const ERC20TokenInterfaceBin = `0x`

// DeployERC20TokenInterface deploys a new Ethereum contract, binding an instance of ERC20TokenInterface to it.
func DeployERC20TokenInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20TokenInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20TokenInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20TokenInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20TokenInterface{ERC20TokenInterfaceCaller: ERC20TokenInterfaceCaller{contract: contract}, ERC20TokenInterfaceTransactor: ERC20TokenInterfaceTransactor{contract: contract}, ERC20TokenInterfaceFilterer: ERC20TokenInterfaceFilterer{contract: contract}}, nil
}

// ERC20TokenInterface is an auto generated Go binding around an Ethereum contract.
type ERC20TokenInterface struct {
	ERC20TokenInterfaceCaller     // Read-only binding to the contract
	ERC20TokenInterfaceTransactor // Write-only binding to the contract
	ERC20TokenInterfaceFilterer   // Log filterer for contract events
}

// ERC20TokenInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TokenInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TokenInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TokenInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TokenInterfaceSession struct {
	Contract     *ERC20TokenInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ERC20TokenInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TokenInterfaceCallerSession struct {
	Contract *ERC20TokenInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ERC20TokenInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TokenInterfaceTransactorSession struct {
	Contract     *ERC20TokenInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ERC20TokenInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TokenInterfaceRaw struct {
	Contract *ERC20TokenInterface // Generic contract binding to access the raw methods on
}

// ERC20TokenInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TokenInterfaceCallerRaw struct {
	Contract *ERC20TokenInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TokenInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TokenInterfaceTransactorRaw struct {
	Contract *ERC20TokenInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20TokenInterface creates a new instance of ERC20TokenInterface, bound to a specific deployed contract.
func NewERC20TokenInterface(address common.Address, backend bind.ContractBackend) (*ERC20TokenInterface, error) {
	contract, err := bindERC20TokenInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenInterface{ERC20TokenInterfaceCaller: ERC20TokenInterfaceCaller{contract: contract}, ERC20TokenInterfaceTransactor: ERC20TokenInterfaceTransactor{contract: contract}, ERC20TokenInterfaceFilterer: ERC20TokenInterfaceFilterer{contract: contract}}, nil
}

// NewERC20TokenInterfaceCaller creates a new read-only instance of ERC20TokenInterface, bound to a specific deployed contract.
func NewERC20TokenInterfaceCaller(address common.Address, caller bind.ContractCaller) (*ERC20TokenInterfaceCaller, error) {
	contract, err := bindERC20TokenInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenInterfaceCaller{contract: contract}, nil
}

// NewERC20TokenInterfaceTransactor creates a new write-only instance of ERC20TokenInterface, bound to a specific deployed contract.
func NewERC20TokenInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TokenInterfaceTransactor, error) {
	contract, err := bindERC20TokenInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenInterfaceTransactor{contract: contract}, nil
}

// NewERC20TokenInterfaceFilterer creates a new log filterer instance of ERC20TokenInterface, bound to a specific deployed contract.
func NewERC20TokenInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TokenInterfaceFilterer, error) {
	contract, err := bindERC20TokenInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenInterfaceFilterer{contract: contract}, nil
}

// bindERC20TokenInterface binds a generic wrapper to an already deployed contract.
func bindERC20TokenInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20TokenInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenInterface *ERC20TokenInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20TokenInterface.Contract.ERC20TokenInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenInterface *ERC20TokenInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenInterface.Contract.ERC20TokenInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenInterface *ERC20TokenInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenInterface.Contract.ERC20TokenInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenInterface *ERC20TokenInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20TokenInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenInterface *ERC20TokenInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenInterface *ERC20TokenInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenInterface.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ERC20TokenInterface *ERC20TokenInterfaceCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20TokenInterface.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ERC20TokenInterface *ERC20TokenInterfaceSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20TokenInterface.Contract.Allowance(&_ERC20TokenInterface.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ERC20TokenInterface *ERC20TokenInterfaceCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20TokenInterface.Contract.Allowance(&_ERC20TokenInterface.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC20TokenInterface *ERC20TokenInterfaceCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20TokenInterface.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC20TokenInterface *ERC20TokenInterfaceSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC20TokenInterface.Contract.BalanceOf(&_ERC20TokenInterface.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC20TokenInterface *ERC20TokenInterfaceCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC20TokenInterface.Contract.BalanceOf(&_ERC20TokenInterface.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_ERC20TokenInterface *ERC20TokenInterfaceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20TokenInterface.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_ERC20TokenInterface *ERC20TokenInterfaceSession) TotalSupply() (*big.Int, error) {
	return _ERC20TokenInterface.Contract.TotalSupply(&_ERC20TokenInterface.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_ERC20TokenInterface *ERC20TokenInterfaceCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20TokenInterface.Contract.TotalSupply(&_ERC20TokenInterface.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ERC20TokenInterface *ERC20TokenInterfaceTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenInterface.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ERC20TokenInterface *ERC20TokenInterfaceSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenInterface.Contract.Approve(&_ERC20TokenInterface.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ERC20TokenInterface *ERC20TokenInterfaceTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenInterface.Contract.Approve(&_ERC20TokenInterface.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ERC20TokenInterface *ERC20TokenInterfaceTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenInterface.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ERC20TokenInterface *ERC20TokenInterfaceSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenInterface.Contract.Transfer(&_ERC20TokenInterface.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ERC20TokenInterface *ERC20TokenInterfaceTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenInterface.Contract.Transfer(&_ERC20TokenInterface.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ERC20TokenInterface *ERC20TokenInterfaceTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenInterface.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ERC20TokenInterface *ERC20TokenInterfaceSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenInterface.Contract.TransferFrom(&_ERC20TokenInterface.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ERC20TokenInterface *ERC20TokenInterfaceTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenInterface.Contract.TransferFrom(&_ERC20TokenInterface.TransactOpts, _from, _to, _value)
}

// ERC20TokenInterfaceApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20TokenInterface contract.
type ERC20TokenInterfaceApprovalIterator struct {
	Event *ERC20TokenInterfaceApproval // Event containing the contract specifics and raw log

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
func (it *ERC20TokenInterfaceApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenInterfaceApproval)
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
		it.Event = new(ERC20TokenInterfaceApproval)
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
func (it *ERC20TokenInterfaceApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenInterfaceApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenInterfaceApproval represents a Approval event raised by the ERC20TokenInterface contract.
type ERC20TokenInterfaceApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20TokenInterface *ERC20TokenInterfaceFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20TokenInterfaceApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20TokenInterface.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenInterfaceApprovalIterator{contract: _ERC20TokenInterface.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20TokenInterface *ERC20TokenInterfaceFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20TokenInterfaceApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20TokenInterface.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenInterfaceApproval)
				if err := _ERC20TokenInterface.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20TokenInterfaceTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20TokenInterface contract.
type ERC20TokenInterfaceTransferIterator struct {
	Event *ERC20TokenInterfaceTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20TokenInterfaceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenInterfaceTransfer)
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
		it.Event = new(ERC20TokenInterfaceTransfer)
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
func (it *ERC20TokenInterfaceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenInterfaceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenInterfaceTransfer represents a Transfer event raised by the ERC20TokenInterface contract.
type ERC20TokenInterfaceTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20TokenInterface *ERC20TokenInterfaceFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TokenInterfaceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20TokenInterface.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenInterfaceTransferIterator{contract: _ERC20TokenInterface.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20TokenInterface *ERC20TokenInterfaceFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20TokenInterfaceTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20TokenInterface.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenInterfaceTransfer)
				if err := _ERC20TokenInterface.contract.UnpackLog(event, "Transfer", log); err != nil {
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

