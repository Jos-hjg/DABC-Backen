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

// DABCMetaData contains all meta data concerning the DABC contract.
var DABCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initialAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inviter\",\"type\":\"address\"}],\"name\":\"GetAchievement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"achievement\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lv\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"GetDABC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inviter\",\"type\":\"address\"}],\"name\":\"GetJC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"jc\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"yj\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lv\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetJCBalance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetPledge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inviter\",\"type\":\"address\"}],\"name\":\"GetZT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ena\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"disa\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetZTBalance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"InvalidTimesLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Multiple\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inviter\",\"type\":\"address\"}],\"name\":\"Pledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currtime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SpanMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SpanMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isexist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ZT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"fromwho\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"burning\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ztBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZTRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxAmount\",\"type\":\"uint256\"}],\"name\":\"changeMaxPledge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eachMinedMaxCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eachMinedMinCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emptyPool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inviter\",\"type\":\"address\"}],\"name\":\"getInvitees\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getInvitor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPriceTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"get_invalidtimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"get_valid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"invitees\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxPledgeCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPledgeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"times\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invalidTimes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tblength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ztlength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRevenue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"PDRevenue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ZTRevenue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"JCRevenue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pullSome\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetInvalidTimes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetTimes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"reward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_per\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"total\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DABCABI is the input ABI used to generate the binding from.
// Deprecated: Use DABCMetaData.ABI instead.
var DABCABI = DABCMetaData.ABI

// DABC is an auto generated Go binding around an Ethereum contract.
type DABC struct {
	DABCCaller     // Read-only binding to the contract
	DABCTransactor // Write-only binding to the contract
	DABCFilterer   // Log filterer for contract events
}

// DABCCaller is an auto generated read-only Go binding around an Ethereum contract.
type DABCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DABCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DABCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DABCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DABCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DABCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DABCSession struct {
	Contract     *DABC             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DABCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DABCCallerSession struct {
	Contract *DABCCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DABCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DABCTransactorSession struct {
	Contract     *DABCTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DABCRaw is an auto generated low-level Go binding around an Ethereum contract.
type DABCRaw struct {
	Contract *DABC // Generic contract binding to access the raw methods on
}

// DABCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DABCCallerRaw struct {
	Contract *DABCCaller // Generic read-only contract binding to access the raw methods on
}

// DABCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DABCTransactorRaw struct {
	Contract *DABCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDABC creates a new instance of DABC, bound to a specific deployed contract.
func NewDABC(address common.Address, backend bind.ContractBackend) (*DABC, error) {
	contract, err := bindDABC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DABC{DABCCaller: DABCCaller{contract: contract}, DABCTransactor: DABCTransactor{contract: contract}, DABCFilterer: DABCFilterer{contract: contract}}, nil
}

// NewDABCCaller creates a new read-only instance of DABC, bound to a specific deployed contract.
func NewDABCCaller(address common.Address, caller bind.ContractCaller) (*DABCCaller, error) {
	contract, err := bindDABC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DABCCaller{contract: contract}, nil
}

// NewDABCTransactor creates a new write-only instance of DABC, bound to a specific deployed contract.
func NewDABCTransactor(address common.Address, transactor bind.ContractTransactor) (*DABCTransactor, error) {
	contract, err := bindDABC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DABCTransactor{contract: contract}, nil
}

// NewDABCFilterer creates a new log filterer instance of DABC, bound to a specific deployed contract.
func NewDABCFilterer(address common.Address, filterer bind.ContractFilterer) (*DABCFilterer, error) {
	contract, err := bindDABC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DABCFilterer{contract: contract}, nil
}

// bindDABC binds a generic wrapper to an already deployed contract.
func bindDABC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DABCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DABC *DABCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DABC.Contract.DABCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DABC *DABCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DABC.Contract.DABCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DABC *DABCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DABC.Contract.DABCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DABC *DABCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DABC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DABC *DABCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DABC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DABC *DABCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DABC.Contract.contract.Transact(opts, method, params...)
}

// GetZT is a free data retrieval call binding the contract method 0x2d62a167.
//
// Solidity: function GetZT(address _inviter) view returns(uint256 ena, uint256 disa)
func (_DABC *DABCCaller) GetZT(opts *bind.CallOpts, _inviter common.Address) (struct {
	Ena  *big.Int
	Disa *big.Int
}, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "GetZT", _inviter)

	outstruct := new(struct {
		Ena  *big.Int
		Disa *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ena = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Disa = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetZT is a free data retrieval call binding the contract method 0x2d62a167.
//
// Solidity: function GetZT(address _inviter) view returns(uint256 ena, uint256 disa)
func (_DABC *DABCSession) GetZT(_inviter common.Address) (struct {
	Ena  *big.Int
	Disa *big.Int
}, error) {
	return _DABC.Contract.GetZT(&_DABC.CallOpts, _inviter)
}

// GetZT is a free data retrieval call binding the contract method 0x2d62a167.
//
// Solidity: function GetZT(address _inviter) view returns(uint256 ena, uint256 disa)
func (_DABC *DABCCallerSession) GetZT(_inviter common.Address) (struct {
	Ena  *big.Int
	Disa *big.Int
}, error) {
	return _DABC.Contract.GetZT(&_DABC.CallOpts, _inviter)
}

// InvalidTimesLimit is a free data retrieval call binding the contract method 0xf1714305.
//
// Solidity: function InvalidTimesLimit() view returns(uint256)
func (_DABC *DABCCaller) InvalidTimesLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "InvalidTimesLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvalidTimesLimit is a free data retrieval call binding the contract method 0xf1714305.
//
// Solidity: function InvalidTimesLimit() view returns(uint256)
func (_DABC *DABCSession) InvalidTimesLimit() (*big.Int, error) {
	return _DABC.Contract.InvalidTimesLimit(&_DABC.CallOpts)
}

// InvalidTimesLimit is a free data retrieval call binding the contract method 0xf1714305.
//
// Solidity: function InvalidTimesLimit() view returns(uint256)
func (_DABC *DABCCallerSession) InvalidTimesLimit() (*big.Int, error) {
	return _DABC.Contract.InvalidTimesLimit(&_DABC.CallOpts)
}

// Multiple is a free data retrieval call binding the contract method 0x8e80678a.
//
// Solidity: function Multiple() view returns(uint256)
func (_DABC *DABCCaller) Multiple(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "Multiple")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Multiple is a free data retrieval call binding the contract method 0x8e80678a.
//
// Solidity: function Multiple() view returns(uint256)
func (_DABC *DABCSession) Multiple() (*big.Int, error) {
	return _DABC.Contract.Multiple(&_DABC.CallOpts)
}

// Multiple is a free data retrieval call binding the contract method 0x8e80678a.
//
// Solidity: function Multiple() view returns(uint256)
func (_DABC *DABCCallerSession) Multiple() (*big.Int, error) {
	return _DABC.Contract.Multiple(&_DABC.CallOpts)
}

// OOD is a free data retrieval call binding the contract method 0x368bc0a3.
//
// Solidity: function OOD() view returns(uint256)
func (_DABC *DABCCaller) OOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "OOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OOD is a free data retrieval call binding the contract method 0x368bc0a3.
//
// Solidity: function OOD() view returns(uint256)
func (_DABC *DABCSession) OOD() (*big.Int, error) {
	return _DABC.Contract.OOD(&_DABC.CallOpts)
}

// OOD is a free data retrieval call binding the contract method 0x368bc0a3.
//
// Solidity: function OOD() view returns(uint256)
func (_DABC *DABCCallerSession) OOD() (*big.Int, error) {
	return _DABC.Contract.OOD(&_DABC.CallOpts)
}

// SpanMax is a free data retrieval call binding the contract method 0x3d014014.
//
// Solidity: function SpanMax() view returns(uint256)
func (_DABC *DABCCaller) SpanMax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "SpanMax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpanMax is a free data retrieval call binding the contract method 0x3d014014.
//
// Solidity: function SpanMax() view returns(uint256)
func (_DABC *DABCSession) SpanMax() (*big.Int, error) {
	return _DABC.Contract.SpanMax(&_DABC.CallOpts)
}

// SpanMax is a free data retrieval call binding the contract method 0x3d014014.
//
// Solidity: function SpanMax() view returns(uint256)
func (_DABC *DABCCallerSession) SpanMax() (*big.Int, error) {
	return _DABC.Contract.SpanMax(&_DABC.CallOpts)
}

// SpanMin is a free data retrieval call binding the contract method 0x88eaa4da.
//
// Solidity: function SpanMin() view returns(uint256)
func (_DABC *DABCCaller) SpanMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "SpanMin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpanMin is a free data retrieval call binding the contract method 0x88eaa4da.
//
// Solidity: function SpanMin() view returns(uint256)
func (_DABC *DABCSession) SpanMin() (*big.Int, error) {
	return _DABC.Contract.SpanMin(&_DABC.CallOpts)
}

// SpanMin is a free data retrieval call binding the contract method 0x88eaa4da.
//
// Solidity: function SpanMin() view returns(uint256)
func (_DABC *DABCCallerSession) SpanMin() (*big.Int, error) {
	return _DABC.Contract.SpanMin(&_DABC.CallOpts)
}

// TB is a free data retrieval call binding the contract method 0x6dea05b6.
//
// Solidity: function TB(address , uint256 ) view returns(uint256 time, bool valid, uint256 cost, uint256 balance, bool isexist, uint256 reward)
func (_DABC *DABCCaller) TB(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Time    *big.Int
	Valid   bool
	Cost    *big.Int
	Balance *big.Int
	Isexist bool
	Reward  *big.Int
}, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "TB", arg0, arg1)

	outstruct := new(struct {
		Time    *big.Int
		Valid   bool
		Cost    *big.Int
		Balance *big.Int
		Isexist bool
		Reward  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Time = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Valid = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Cost = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Balance = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Isexist = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Reward = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TB is a free data retrieval call binding the contract method 0x6dea05b6.
//
// Solidity: function TB(address , uint256 ) view returns(uint256 time, bool valid, uint256 cost, uint256 balance, bool isexist, uint256 reward)
func (_DABC *DABCSession) TB(arg0 common.Address, arg1 *big.Int) (struct {
	Time    *big.Int
	Valid   bool
	Cost    *big.Int
	Balance *big.Int
	Isexist bool
	Reward  *big.Int
}, error) {
	return _DABC.Contract.TB(&_DABC.CallOpts, arg0, arg1)
}

// TB is a free data retrieval call binding the contract method 0x6dea05b6.
//
// Solidity: function TB(address , uint256 ) view returns(uint256 time, bool valid, uint256 cost, uint256 balance, bool isexist, uint256 reward)
func (_DABC *DABCCallerSession) TB(arg0 common.Address, arg1 *big.Int) (struct {
	Time    *big.Int
	Valid   bool
	Cost    *big.Int
	Balance *big.Int
	Isexist bool
	Reward  *big.Int
}, error) {
	return _DABC.Contract.TB(&_DABC.CallOpts, arg0, arg1)
}

// ZT is a free data retrieval call binding the contract method 0xeafc623f.
//
// Solidity: function ZT(address , uint256 ) view returns(address fromwho, bool enable, bool burning, uint256 time, uint256 ztBalance)
func (_DABC *DABCCaller) ZT(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Fromwho   common.Address
	Enable    bool
	Burning   bool
	Time      *big.Int
	ZtBalance *big.Int
}, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "ZT", arg0, arg1)

	outstruct := new(struct {
		Fromwho   common.Address
		Enable    bool
		Burning   bool
		Time      *big.Int
		ZtBalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fromwho = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Enable = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Burning = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Time = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ZtBalance = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ZT is a free data retrieval call binding the contract method 0xeafc623f.
//
// Solidity: function ZT(address , uint256 ) view returns(address fromwho, bool enable, bool burning, uint256 time, uint256 ztBalance)
func (_DABC *DABCSession) ZT(arg0 common.Address, arg1 *big.Int) (struct {
	Fromwho   common.Address
	Enable    bool
	Burning   bool
	Time      *big.Int
	ZtBalance *big.Int
}, error) {
	return _DABC.Contract.ZT(&_DABC.CallOpts, arg0, arg1)
}

// ZT is a free data retrieval call binding the contract method 0xeafc623f.
//
// Solidity: function ZT(address , uint256 ) view returns(address fromwho, bool enable, bool burning, uint256 time, uint256 ztBalance)
func (_DABC *DABCCallerSession) ZT(arg0 common.Address, arg1 *big.Int) (struct {
	Fromwho   common.Address
	Enable    bool
	Burning   bool
	Time      *big.Int
	ZtBalance *big.Int
}, error) {
	return _DABC.Contract.ZT(&_DABC.CallOpts, arg0, arg1)
}

// ZTRate is a free data retrieval call binding the contract method 0x4c3c632b.
//
// Solidity: function ZTRate() view returns(uint256)
func (_DABC *DABCCaller) ZTRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "ZTRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZTRate is a free data retrieval call binding the contract method 0x4c3c632b.
//
// Solidity: function ZTRate() view returns(uint256)
func (_DABC *DABCSession) ZTRate() (*big.Int, error) {
	return _DABC.Contract.ZTRate(&_DABC.CallOpts)
}

// ZTRate is a free data retrieval call binding the contract method 0x4c3c632b.
//
// Solidity: function ZTRate() view returns(uint256)
func (_DABC *DABCCallerSession) ZTRate() (*big.Int, error) {
	return _DABC.Contract.ZTRate(&_DABC.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_DABC *DABCCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_DABC *DABCSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _DABC.Contract.Allowance(&_DABC.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_DABC *DABCCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _DABC.Contract.Allowance(&_DABC.CallOpts, _owner, _spender)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_DABC *DABCCaller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "allowed", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_DABC *DABCSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DABC.Contract.Allowed(&_DABC.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_DABC *DABCCallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DABC.Contract.Allowed(&_DABC.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_DABC *DABCCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_DABC *DABCSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DABC.Contract.BalanceOf(&_DABC.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_DABC *DABCCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DABC.Contract.BalanceOf(&_DABC.CallOpts, _owner)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_DABC *DABCCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_DABC *DABCSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _DABC.Contract.Balances(&_DABC.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_DABC *DABCCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _DABC.Contract.Balances(&_DABC.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DABC *DABCCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DABC *DABCSession) Decimals() (uint8, error) {
	return _DABC.Contract.Decimals(&_DABC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DABC *DABCCallerSession) Decimals() (uint8, error) {
	return _DABC.Contract.Decimals(&_DABC.CallOpts)
}

// EachMinedMaxCount is a free data retrieval call binding the contract method 0x7af8c81f.
//
// Solidity: function eachMinedMaxCount() view returns(uint256)
func (_DABC *DABCCaller) EachMinedMaxCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "eachMinedMaxCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EachMinedMaxCount is a free data retrieval call binding the contract method 0x7af8c81f.
//
// Solidity: function eachMinedMaxCount() view returns(uint256)
func (_DABC *DABCSession) EachMinedMaxCount() (*big.Int, error) {
	return _DABC.Contract.EachMinedMaxCount(&_DABC.CallOpts)
}

// EachMinedMaxCount is a free data retrieval call binding the contract method 0x7af8c81f.
//
// Solidity: function eachMinedMaxCount() view returns(uint256)
func (_DABC *DABCCallerSession) EachMinedMaxCount() (*big.Int, error) {
	return _DABC.Contract.EachMinedMaxCount(&_DABC.CallOpts)
}

// EachMinedMinCount is a free data retrieval call binding the contract method 0x2f26a1a5.
//
// Solidity: function eachMinedMinCount() view returns(uint256)
func (_DABC *DABCCaller) EachMinedMinCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "eachMinedMinCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EachMinedMinCount is a free data retrieval call binding the contract method 0x2f26a1a5.
//
// Solidity: function eachMinedMinCount() view returns(uint256)
func (_DABC *DABCSession) EachMinedMinCount() (*big.Int, error) {
	return _DABC.Contract.EachMinedMinCount(&_DABC.CallOpts)
}

// EachMinedMinCount is a free data retrieval call binding the contract method 0x2f26a1a5.
//
// Solidity: function eachMinedMinCount() view returns(uint256)
func (_DABC *DABCCallerSession) EachMinedMinCount() (*big.Int, error) {
	return _DABC.Contract.EachMinedMinCount(&_DABC.CallOpts)
}

// GetInvitees is a free data retrieval call binding the contract method 0xe9881a5e.
//
// Solidity: function getInvitees(address _inviter) view returns(address[])
func (_DABC *DABCCaller) GetInvitees(opts *bind.CallOpts, _inviter common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "getInvitees", _inviter)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetInvitees is a free data retrieval call binding the contract method 0xe9881a5e.
//
// Solidity: function getInvitees(address _inviter) view returns(address[])
func (_DABC *DABCSession) GetInvitees(_inviter common.Address) ([]common.Address, error) {
	return _DABC.Contract.GetInvitees(&_DABC.CallOpts, _inviter)
}

// GetInvitees is a free data retrieval call binding the contract method 0xe9881a5e.
//
// Solidity: function getInvitees(address _inviter) view returns(address[])
func (_DABC *DABCCallerSession) GetInvitees(_inviter common.Address) ([]common.Address, error) {
	return _DABC.Contract.GetInvitees(&_DABC.CallOpts, _inviter)
}

// GetInvitor is a free data retrieval call binding the contract method 0xb0c73bc6.
//
// Solidity: function getInvitor(address ) view returns(address)
func (_DABC *DABCCaller) GetInvitor(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "getInvitor", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInvitor is a free data retrieval call binding the contract method 0xb0c73bc6.
//
// Solidity: function getInvitor(address ) view returns(address)
func (_DABC *DABCSession) GetInvitor(arg0 common.Address) (common.Address, error) {
	return _DABC.Contract.GetInvitor(&_DABC.CallOpts, arg0)
}

// GetInvitor is a free data retrieval call binding the contract method 0xb0c73bc6.
//
// Solidity: function getInvitor(address ) view returns(address)
func (_DABC *DABCCallerSession) GetInvitor(arg0 common.Address) (common.Address, error) {
	return _DABC.Contract.GetInvitor(&_DABC.CallOpts, arg0)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() pure returns(int256)
func (_DABC *DABCCaller) GetLatestPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "getLatestPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() pure returns(int256)
func (_DABC *DABCSession) GetLatestPrice() (*big.Int, error) {
	return _DABC.Contract.GetLatestPrice(&_DABC.CallOpts)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() pure returns(int256)
func (_DABC *DABCCallerSession) GetLatestPrice() (*big.Int, error) {
	return _DABC.Contract.GetLatestPrice(&_DABC.CallOpts)
}

// GetLatestPriceTimestamp is a free data retrieval call binding the contract method 0x0d0d2b01.
//
// Solidity: function getLatestPriceTimestamp() view returns(uint256)
func (_DABC *DABCCaller) GetLatestPriceTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "getLatestPriceTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPriceTimestamp is a free data retrieval call binding the contract method 0x0d0d2b01.
//
// Solidity: function getLatestPriceTimestamp() view returns(uint256)
func (_DABC *DABCSession) GetLatestPriceTimestamp() (*big.Int, error) {
	return _DABC.Contract.GetLatestPriceTimestamp(&_DABC.CallOpts)
}

// GetLatestPriceTimestamp is a free data retrieval call binding the contract method 0x0d0d2b01.
//
// Solidity: function getLatestPriceTimestamp() view returns(uint256)
func (_DABC *DABCCallerSession) GetLatestPriceTimestamp() (*big.Int, error) {
	return _DABC.Contract.GetLatestPriceTimestamp(&_DABC.CallOpts)
}

// GetInvalidtimes is a free data retrieval call binding the contract method 0x24183327.
//
// Solidity: function get_invalidtimes(address _sender) view returns(uint256)
func (_DABC *DABCCaller) GetInvalidtimes(opts *bind.CallOpts, _sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "get_invalidtimes", _sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInvalidtimes is a free data retrieval call binding the contract method 0x24183327.
//
// Solidity: function get_invalidtimes(address _sender) view returns(uint256)
func (_DABC *DABCSession) GetInvalidtimes(_sender common.Address) (*big.Int, error) {
	return _DABC.Contract.GetInvalidtimes(&_DABC.CallOpts, _sender)
}

// GetInvalidtimes is a free data retrieval call binding the contract method 0x24183327.
//
// Solidity: function get_invalidtimes(address _sender) view returns(uint256)
func (_DABC *DABCCallerSession) GetInvalidtimes(_sender common.Address) (*big.Int, error) {
	return _DABC.Contract.GetInvalidtimes(&_DABC.CallOpts, _sender)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x2212dbc3.
//
// Solidity: function get_timestamp() view returns(uint256)
func (_DABC *DABCCaller) GetTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "get_timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0x2212dbc3.
//
// Solidity: function get_timestamp() view returns(uint256)
func (_DABC *DABCSession) GetTimestamp() (*big.Int, error) {
	return _DABC.Contract.GetTimestamp(&_DABC.CallOpts)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x2212dbc3.
//
// Solidity: function get_timestamp() view returns(uint256)
func (_DABC *DABCCallerSession) GetTimestamp() (*big.Int, error) {
	return _DABC.Contract.GetTimestamp(&_DABC.CallOpts)
}

// GetValid is a free data retrieval call binding the contract method 0x53bb6329.
//
// Solidity: function get_valid(address owner, uint256 current) view returns(uint256)
func (_DABC *DABCCaller) GetValid(opts *bind.CallOpts, owner common.Address, current *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "get_valid", owner, current)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValid is a free data retrieval call binding the contract method 0x53bb6329.
//
// Solidity: function get_valid(address owner, uint256 current) view returns(uint256)
func (_DABC *DABCSession) GetValid(owner common.Address, current *big.Int) (*big.Int, error) {
	return _DABC.Contract.GetValid(&_DABC.CallOpts, owner, current)
}

// GetValid is a free data retrieval call binding the contract method 0x53bb6329.
//
// Solidity: function get_valid(address owner, uint256 current) view returns(uint256)
func (_DABC *DABCCallerSession) GetValid(owner common.Address, current *big.Int) (*big.Int, error) {
	return _DABC.Contract.GetValid(&_DABC.CallOpts, owner, current)
}

// Invitees is a free data retrieval call binding the contract method 0x5f2f8797.
//
// Solidity: function invitees(address , uint256 ) view returns(address)
func (_DABC *DABCCaller) Invitees(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "invitees", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Invitees is a free data retrieval call binding the contract method 0x5f2f8797.
//
// Solidity: function invitees(address , uint256 ) view returns(address)
func (_DABC *DABCSession) Invitees(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _DABC.Contract.Invitees(&_DABC.CallOpts, arg0, arg1)
}

// Invitees is a free data retrieval call binding the contract method 0x5f2f8797.
//
// Solidity: function invitees(address , uint256 ) view returns(address)
func (_DABC *DABCCallerSession) Invitees(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _DABC.Contract.Invitees(&_DABC.CallOpts, arg0, arg1)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(uint256 maxPledgeCount, uint256 lastPledgeTime, uint256 times, uint256 invalidTimes, uint256 tblength, uint256 ztlength, uint256 totalBalance, uint256 totalRevenue, uint256 PDRevenue, uint256 ZTRevenue, uint256 JCRevenue)
func (_DABC *DABCCaller) Minters(opts *bind.CallOpts, arg0 common.Address) (struct {
	MaxPledgeCount *big.Int
	LastPledgeTime *big.Int
	Times          *big.Int
	InvalidTimes   *big.Int
	Tblength       *big.Int
	Ztlength       *big.Int
	TotalBalance   *big.Int
	TotalRevenue   *big.Int
	PDRevenue      *big.Int
	ZTRevenue      *big.Int
	JCRevenue      *big.Int
}, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "minters", arg0)

	outstruct := new(struct {
		MaxPledgeCount *big.Int
		LastPledgeTime *big.Int
		Times          *big.Int
		InvalidTimes   *big.Int
		Tblength       *big.Int
		Ztlength       *big.Int
		TotalBalance   *big.Int
		TotalRevenue   *big.Int
		PDRevenue      *big.Int
		ZTRevenue      *big.Int
		JCRevenue      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxPledgeCount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastPledgeTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Times = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.InvalidTimes = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Tblength = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Ztlength = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TotalBalance = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.TotalRevenue = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.PDRevenue = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.ZTRevenue = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.JCRevenue = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(uint256 maxPledgeCount, uint256 lastPledgeTime, uint256 times, uint256 invalidTimes, uint256 tblength, uint256 ztlength, uint256 totalBalance, uint256 totalRevenue, uint256 PDRevenue, uint256 ZTRevenue, uint256 JCRevenue)
func (_DABC *DABCSession) Minters(arg0 common.Address) (struct {
	MaxPledgeCount *big.Int
	LastPledgeTime *big.Int
	Times          *big.Int
	InvalidTimes   *big.Int
	Tblength       *big.Int
	Ztlength       *big.Int
	TotalBalance   *big.Int
	TotalRevenue   *big.Int
	PDRevenue      *big.Int
	ZTRevenue      *big.Int
	JCRevenue      *big.Int
}, error) {
	return _DABC.Contract.Minters(&_DABC.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(uint256 maxPledgeCount, uint256 lastPledgeTime, uint256 times, uint256 invalidTimes, uint256 tblength, uint256 ztlength, uint256 totalBalance, uint256 totalRevenue, uint256 PDRevenue, uint256 ZTRevenue, uint256 JCRevenue)
func (_DABC *DABCCallerSession) Minters(arg0 common.Address) (struct {
	MaxPledgeCount *big.Int
	LastPledgeTime *big.Int
	Times          *big.Int
	InvalidTimes   *big.Int
	Tblength       *big.Int
	Ztlength       *big.Int
	TotalBalance   *big.Int
	TotalRevenue   *big.Int
	PDRevenue      *big.Int
	ZTRevenue      *big.Int
	JCRevenue      *big.Int
}, error) {
	return _DABC.Contract.Minters(&_DABC.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DABC *DABCCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DABC *DABCSession) Name() (string, error) {
	return _DABC.Contract.Name(&_DABC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DABC *DABCCallerSession) Name() (string, error) {
	return _DABC.Contract.Name(&_DABC.CallOpts)
}

// PoolBalance is a free data retrieval call binding the contract method 0x96365d44.
//
// Solidity: function poolBalance() view returns(uint256)
func (_DABC *DABCCaller) PoolBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "poolBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolBalance is a free data retrieval call binding the contract method 0x96365d44.
//
// Solidity: function poolBalance() view returns(uint256)
func (_DABC *DABCSession) PoolBalance() (*big.Int, error) {
	return _DABC.Contract.PoolBalance(&_DABC.CallOpts)
}

// PoolBalance is a free data retrieval call binding the contract method 0x96365d44.
//
// Solidity: function poolBalance() view returns(uint256)
func (_DABC *DABCCallerSession) PoolBalance() (*big.Int, error) {
	return _DABC.Contract.PoolBalance(&_DABC.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_DABC *DABCCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_DABC *DABCSession) Rate() (*big.Int, error) {
	return _DABC.Contract.Rate(&_DABC.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_DABC *DABCCallerSession) Rate() (*big.Int, error) {
	return _DABC.Contract.Rate(&_DABC.CallOpts)
}

// Reward is a free data retrieval call binding the contract method 0x6353586b.
//
// Solidity: function reward(address ) view returns(uint256)
func (_DABC *DABCCaller) Reward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "reward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reward is a free data retrieval call binding the contract method 0x6353586b.
//
// Solidity: function reward(address ) view returns(uint256)
func (_DABC *DABCSession) Reward(arg0 common.Address) (*big.Int, error) {
	return _DABC.Contract.Reward(&_DABC.CallOpts, arg0)
}

// Reward is a free data retrieval call binding the contract method 0x6353586b.
//
// Solidity: function reward(address ) view returns(uint256)
func (_DABC *DABCCallerSession) Reward(arg0 common.Address) (*big.Int, error) {
	return _DABC.Contract.Reward(&_DABC.CallOpts, arg0)
}

// StatePer is a free data retrieval call binding the contract method 0x6ef37748.
//
// Solidity: function state_per() view returns(uint256)
func (_DABC *DABCCaller) StatePer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "state_per")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StatePer is a free data retrieval call binding the contract method 0x6ef37748.
//
// Solidity: function state_per() view returns(uint256)
func (_DABC *DABCSession) StatePer() (*big.Int, error) {
	return _DABC.Contract.StatePer(&_DABC.CallOpts)
}

// StatePer is a free data retrieval call binding the contract method 0x6ef37748.
//
// Solidity: function state_per() view returns(uint256)
func (_DABC *DABCCallerSession) StatePer() (*big.Int, error) {
	return _DABC.Contract.StatePer(&_DABC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DABC *DABCCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DABC *DABCSession) Symbol() (string, error) {
	return _DABC.Contract.Symbol(&_DABC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DABC *DABCCallerSession) Symbol() (string, error) {
	return _DABC.Contract.Symbol(&_DABC.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256)
func (_DABC *DABCCaller) Total(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "total")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256)
func (_DABC *DABCSession) Total() (*big.Int, error) {
	return _DABC.Contract.Total(&_DABC.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256)
func (_DABC *DABCCallerSession) Total() (*big.Int, error) {
	return _DABC.Contract.Total(&_DABC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DABC *DABCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DABC *DABCSession) TotalSupply() (*big.Int, error) {
	return _DABC.Contract.TotalSupply(&_DABC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DABC *DABCCallerSession) TotalSupply() (*big.Int, error) {
	return _DABC.Contract.TotalSupply(&_DABC.CallOpts)
}

// WinTime is a free data retrieval call binding the contract method 0x2498764a.
//
// Solidity: function winTime() view returns(uint256)
func (_DABC *DABCCaller) WinTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DABC.contract.Call(opts, &out, "winTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WinTime is a free data retrieval call binding the contract method 0x2498764a.
//
// Solidity: function winTime() view returns(uint256)
func (_DABC *DABCSession) WinTime() (*big.Int, error) {
	return _DABC.Contract.WinTime(&_DABC.CallOpts)
}

// WinTime is a free data retrieval call binding the contract method 0x2498764a.
//
// Solidity: function winTime() view returns(uint256)
func (_DABC *DABCCallerSession) WinTime() (*big.Int, error) {
	return _DABC.Contract.WinTime(&_DABC.CallOpts)
}

// GetAchievement is a paid mutator transaction binding the contract method 0x00b93891.
//
// Solidity: function GetAchievement(address _inviter) returns(uint256 achievement, uint256 lv)
func (_DABC *DABCTransactor) GetAchievement(opts *bind.TransactOpts, _inviter common.Address) (*types.Transaction, error) {
	return _DABC.contract.Transact(opts, "GetAchievement", _inviter)
}

// GetAchievement is a paid mutator transaction binding the contract method 0x00b93891.
//
// Solidity: function GetAchievement(address _inviter) returns(uint256 achievement, uint256 lv)
func (_DABC *DABCSession) GetAchievement(_inviter common.Address) (*types.Transaction, error) {
	return _DABC.Contract.GetAchievement(&_DABC.TransactOpts, _inviter)
}

// GetAchievement is a paid mutator transaction binding the contract method 0x00b93891.
//
// Solidity: function GetAchievement(address _inviter) returns(uint256 achievement, uint256 lv)
func (_DABC *DABCTransactorSession) GetAchievement(_inviter common.Address) (*types.Transaction, error) {
	return _DABC.Contract.GetAchievement(&_DABC.TransactOpts, _inviter)
}

// GetDABC is a paid mutator transaction binding the contract method 0x112f1780.
//
// Solidity: function GetDABC(uint256 _value) returns()
func (_DABC *DABCTransactor) GetDABC(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _DABC.contract.Transact(opts, "GetDABC", _value)
}

// GetDABC is a paid mutator transaction binding the contract method 0x112f1780.
//
// Solidity: function GetDABC(uint256 _value) returns()
func (_DABC *DABCSession) GetDABC(_value *big.Int) (*types.Transaction, error) {
	return _DABC.Contract.GetDABC(&_DABC.TransactOpts, _value)
}

// GetDABC is a paid mutator transaction binding the contract method 0x112f1780.
//
// Solidity: function GetDABC(uint256 _value) returns()
func (_DABC *DABCTransactorSession) GetDABC(_value *big.Int) (*types.Transaction, error) {
	return _DABC.Contract.GetDABC(&_DABC.TransactOpts, _value)
}

// GetJC is a paid mutator transaction binding the contract method 0x28080e43.
//
// Solidity: function GetJC(address _inviter) returns(uint256 jc, uint256 yj, uint256 lv)
func (_DABC *DABCTransactor) GetJC(opts *bind.TransactOpts, _inviter common.Address) (*types.Transaction, error) {
	return _DABC.contract.Transact(opts, "GetJC", _inviter)
}

// GetJC is a paid mutator transaction binding the contract method 0x28080e43.
//
// Solidity: function GetJC(address _inviter) returns(uint256 jc, uint256 yj, uint256 lv)
func (_DABC *DABCSession) GetJC(_inviter common.Address) (*types.Transaction, error) {
	return _DABC.Contract.GetJC(&_DABC.TransactOpts, _inviter)
}

// GetJC is a paid mutator transaction binding the contract method 0x28080e43.
//
// Solidity: function GetJC(address _inviter) returns(uint256 jc, uint256 yj, uint256 lv)
func (_DABC *DABCTransactorSession) GetJC(_inviter common.Address) (*types.Transaction, error) {
	return _DABC.Contract.GetJC(&_DABC.TransactOpts, _inviter)
}

// GetJCBalance is a paid mutator transaction binding the contract method 0x06d76990.
//
// Solidity: function GetJCBalance() payable returns()
func (_DABC *DABCTransactor) GetJCBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DABC.contract.Transact(opts, "GetJCBalance")
}

// GetJCBalance is a paid mutator transaction binding the contract method 0x06d76990.
//
// Solidity: function GetJCBalance() payable returns()
func (_DABC *DABCSession) GetJCBalance() (*types.Transaction, error) {
	return _DABC.Contract.GetJCBalance(&_DABC.TransactOpts)
}

// GetJCBalance is a paid mutator transaction binding the contract method 0x06d76990.
//
// Solidity: function GetJCBalance() payable returns()
func (_DABC *DABCTransactorSession) GetJCBalance() (*types.Transaction, error) {
	return _DABC.Contract.GetJCBalance(&_DABC.TransactOpts)
}

// GetPledge is a paid mutator transaction binding the contract method 0x4a67810e.
//
// Solidity: function GetPledge() payable returns()
func (_DABC *DABCTransactor) GetPledge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DABC.contract.Transact(opts, "GetPledge")
}

// GetPledge is a paid mutator transaction binding the contract method 0x4a67810e.
//
// Solidity: function GetPledge() payable returns()
func (_DABC *DABCSession) GetPledge() (*types.Transaction, error) {
	return _DABC.Contract.GetPledge(&_DABC.TransactOpts)
}

// GetPledge is a paid mutator transaction binding the contract method 0x4a67810e.
//
// Solidity: function GetPledge() payable returns()
func (_DABC *DABCTransactorSession) GetPledge() (*types.Transaction, error) {
	return _DABC.Contract.GetPledge(&_DABC.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x30f3363c.
//
// Solidity: function GetReward() payable returns()
func (_DABC *DABCTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DABC.contract.Transact(opts, "GetReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x30f3363c.
//
// Solidity: function GetReward() payable returns()
func (_DABC *DABCSession) GetReward() (*types.Transaction, error) {
	return _DABC.Contract.GetReward(&_DABC.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x30f3363c.
//
// Solidity: function GetReward() payable returns()
func (_DABC *DABCTransactorSession) GetReward() (*types.Transaction, error) {
	return _DABC.Contract.GetReward(&_DABC.TransactOpts)
}

// GetZTBalance is a paid mutator transaction binding the contract method 0xbbe35232.
//
// Solidity: function GetZTBalance() payable returns()
func (_DABC *DABCTransactor) GetZTBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DABC.contract.Transact(opts, "GetZTBalance")
}

// GetZTBalance is a paid mutator transaction binding the contract method 0xbbe35232.
//
// Solidity: function GetZTBalance() payable returns()
func (_DABC *DABCSession) GetZTBalance() (*types.Transaction, error) {
	return _DABC.Contract.GetZTBalance(&_DABC.TransactOpts)
}

// GetZTBalance is a paid mutator transaction binding the contract method 0xbbe35232.
//
// Solidity: function GetZTBalance() payable returns()
func (_DABC *DABCTransactorSession) GetZTBalance() (*types.Transaction, error) {
	return _DABC.Contract.GetZTBalance(&_DABC.TransactOpts)
}

// Pledge is a paid mutator transaction binding the contract method 0x1c596371.
//
// Solidity: function Pledge(address _inviter) payable returns(uint256 currtime, uint256 amount)
func (_DABC *DABCTransactor) Pledge(opts *bind.TransactOpts, _inviter common.Address) (*types.Transaction, error) {
	return _DABC.contract.Transact(opts, "Pledge", _inviter)
}

// Pledge is a paid mutator transaction binding the contract method 0x1c596371.
//
// Solidity: function Pledge(address _inviter) payable returns(uint256 currtime, uint256 amount)
func (_DABC *DABCSession) Pledge(_inviter common.Address) (*types.Transaction, error) {
	return _DABC.Contract.Pledge(&_DABC.TransactOpts, _inviter)
}

// Pledge is a paid mutator transaction binding the contract method 0x1c596371.
//
// Solidity: function Pledge(address _inviter) payable returns(uint256 currtime, uint256 amount)
func (_DABC *DABCTransactorSession) Pledge(_inviter common.Address) (*types.Transaction, error) {
	return _DABC.Contract.Pledge(&_DABC.TransactOpts, _inviter)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_DABC *DABCTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DABC.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_DABC *DABCSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DABC.Contract.Approve(&_DABC.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_DABC *DABCTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DABC.Contract.Approve(&_DABC.TransactOpts, _spender, _value)
}

// ChangeMaxPledge is a paid mutator transaction binding the contract method 0x799dba39.
//
// Solidity: function changeMaxPledge(address _target, uint256 _maxAmount) payable returns()
func (_DABC *DABCTransactor) ChangeMaxPledge(opts *bind.TransactOpts, _target common.Address, _maxAmount *big.Int) (*types.Transaction, error) {
	return _DABC.contract.Transact(opts, "changeMaxPledge", _target, _maxAmount)
}

// ChangeMaxPledge is a paid mutator transaction binding the contract method 0x799dba39.
//
// Solidity: function changeMaxPledge(address _target, uint256 _maxAmount) payable returns()
func (_DABC *DABCSession) ChangeMaxPledge(_target common.Address, _maxAmount *big.Int) (*types.Transaction, error) {
	return _DABC.Contract.ChangeMaxPledge(&_DABC.TransactOpts, _target, _maxAmount)
}

// ChangeMaxPledge is a paid mutator transaction binding the contract method 0x799dba39.
//
// Solidity: function changeMaxPledge(address _target, uint256 _maxAmount) payable returns()
func (_DABC *DABCTransactorSession) ChangeMaxPledge(_target common.Address, _maxAmount *big.Int) (*types.Transaction, error) {
	return _DABC.Contract.ChangeMaxPledge(&_DABC.TransactOpts, _target, _maxAmount)
}

// EmptyPool is a paid mutator transaction binding the contract method 0xeecc84e4.
//
// Solidity: function emptyPool() payable returns()
func (_DABC *DABCTransactor) EmptyPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DABC.contract.Transact(opts, "emptyPool")
}

// EmptyPool is a paid mutator transaction binding the contract method 0xeecc84e4.
//
// Solidity: function emptyPool() payable returns()
func (_DABC *DABCSession) EmptyPool() (*types.Transaction, error) {
	return _DABC.Contract.EmptyPool(&_DABC.TransactOpts)
}

// EmptyPool is a paid mutator transaction binding the contract method 0xeecc84e4.
//
// Solidity: function emptyPool() payable returns()
func (_DABC *DABCTransactorSession) EmptyPool() (*types.Transaction, error) {
	return _DABC.Contract.EmptyPool(&_DABC.TransactOpts)
}

// PullSome is a paid mutator transaction binding the contract method 0x75d1b1a4.
//
// Solidity: function pullSome() payable returns()
func (_DABC *DABCTransactor) PullSome(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DABC.contract.Transact(opts, "pullSome")
}

// PullSome is a paid mutator transaction binding the contract method 0x75d1b1a4.
//
// Solidity: function pullSome() payable returns()
func (_DABC *DABCSession) PullSome() (*types.Transaction, error) {
	return _DABC.Contract.PullSome(&_DABC.TransactOpts)
}

// PullSome is a paid mutator transaction binding the contract method 0x75d1b1a4.
//
// Solidity: function pullSome() payable returns()
func (_DABC *DABCTransactorSession) PullSome() (*types.Transaction, error) {
	return _DABC.Contract.PullSome(&_DABC.TransactOpts)
}

// ResetInvalidTimes is a paid mutator transaction binding the contract method 0x7a36d796.
//
// Solidity: function resetInvalidTimes() returns(bool)
func (_DABC *DABCTransactor) ResetInvalidTimes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DABC.contract.Transact(opts, "resetInvalidTimes")
}

// ResetInvalidTimes is a paid mutator transaction binding the contract method 0x7a36d796.
//
// Solidity: function resetInvalidTimes() returns(bool)
func (_DABC *DABCSession) ResetInvalidTimes() (*types.Transaction, error) {
	return _DABC.Contract.ResetInvalidTimes(&_DABC.TransactOpts)
}

// ResetInvalidTimes is a paid mutator transaction binding the contract method 0x7a36d796.
//
// Solidity: function resetInvalidTimes() returns(bool)
func (_DABC *DABCTransactorSession) ResetInvalidTimes() (*types.Transaction, error) {
	return _DABC.Contract.ResetInvalidTimes(&_DABC.TransactOpts)
}

// ResetTimes is a paid mutator transaction binding the contract method 0x74361375.
//
// Solidity: function resetTimes() returns(bool)
func (_DABC *DABCTransactor) ResetTimes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DABC.contract.Transact(opts, "resetTimes")
}

// ResetTimes is a paid mutator transaction binding the contract method 0x74361375.
//
// Solidity: function resetTimes() returns(bool)
func (_DABC *DABCSession) ResetTimes() (*types.Transaction, error) {
	return _DABC.Contract.ResetTimes(&_DABC.TransactOpts)
}

// ResetTimes is a paid mutator transaction binding the contract method 0x74361375.
//
// Solidity: function resetTimes() returns(bool)
func (_DABC *DABCTransactorSession) ResetTimes() (*types.Transaction, error) {
	return _DABC.Contract.ResetTimes(&_DABC.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_DABC *DABCTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DABC.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_DABC *DABCSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DABC.Contract.Transfer(&_DABC.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_DABC *DABCTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DABC.Contract.Transfer(&_DABC.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_DABC *DABCTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DABC.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_DABC *DABCSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DABC.Contract.TransferFrom(&_DABC.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_DABC *DABCTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _DABC.Contract.TransferFrom(&_DABC.TransactOpts, _from, _to, _value)
}

// DABCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DABC contract.
type DABCApprovalIterator struct {
	Event *DABCApproval // Event containing the contract specifics and raw log

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
func (it *DABCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DABCApproval)
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
		it.Event = new(DABCApproval)
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
func (it *DABCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DABCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DABCApproval represents a Approval event raised by the DABC contract.
type DABCApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_DABC *DABCFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*DABCApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _DABC.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &DABCApprovalIterator{contract: _DABC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_DABC *DABCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DABCApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _DABC.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DABCApproval)
				if err := _DABC.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_DABC *DABCFilterer) ParseApproval(log types.Log) (*DABCApproval, error) {
	event := new(DABCApproval)
	if err := _DABC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DABCBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the DABC contract.
type DABCBurnIterator struct {
	Event *DABCBurn // Event containing the contract specifics and raw log

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
func (it *DABCBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DABCBurn)
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
		it.Event = new(DABCBurn)
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
func (it *DABCBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DABCBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DABCBurn represents a Burn event raised by the DABC contract.
type DABCBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_DABC *DABCFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*DABCBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _DABC.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &DABCBurnIterator{contract: _DABC.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_DABC *DABCFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *DABCBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _DABC.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DABCBurn)
				if err := _DABC.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_DABC *DABCFilterer) ParseBurn(log types.Log) (*DABCBurn, error) {
	event := new(DABCBurn)
	if err := _DABC.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DABCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DABC contract.
type DABCTransferIterator struct {
	Event *DABCTransfer // Event containing the contract specifics and raw log

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
func (it *DABCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DABCTransfer)
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
		it.Event = new(DABCTransfer)
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
func (it *DABCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DABCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DABCTransfer represents a Transfer event raised by the DABC contract.
type DABCTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_DABC *DABCFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*DABCTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _DABC.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &DABCTransferIterator{contract: _DABC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_DABC *DABCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DABCTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _DABC.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DABCTransfer)
				if err := _DABC.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_DABC *DABCFilterer) ParseTransfer(log types.Log) (*DABCTransfer, error) {
	event := new(DABCTransfer)
	if err := _DABC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
