// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MiniStore

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

// MiniStoreABI is the input ABI used to generate the binding from.
const MiniStoreABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"addValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deleteLastArrayValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getArrayDataLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getArrayLatestValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getArrayValueAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setArrayValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setNumberValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MiniStoreBin is the compiled bytecode used for deploying new contracts.
var MiniStoreBin = "0x608060405234801561001057600080fd5b5061049f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806375a01c0f1161005b57806375a01c0f1461014d57806399ff5f481461017b578063a366474f14610199578063fc7ef48c146101e557610088565b806314d0c25a1461008d57806356298b51146100ab5780635b9af12b146100ed57806369a127991461012f575b600080fd5b610095610203565b6040518082815260200191505060405180910390f35b6100d7600480360360208110156100c157600080fd5b8101908080359060200190929190505050610241565b6040518082815260200191505060405180910390f35b6101196004803603602081101561010357600080fd5b81019080803590602001909291905050506102db565b6040518082815260200191505060405180910390f35b610137610315565b6040518082815260200191505060405180910390f35b6101796004803603602081101561016357600080fd5b8101908080359060200190929190505050610321565b005b61018361032b565b6040518082815260200191505060405180910390f35b6101cf600480360360408110156101af57600080fd5b810190808035906020019092919080359060200190929190505050610335565b6040518082815260200191505060405180910390f35b6101ed6103a2565b6040518082815260200191505060405180910390f35b60008060008054905011156102395760006001600080549050038154811061022757fe5b9060005260206000200154905061023e565b600090505b90565b6000808054905082106102bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f496e646578206f7574206f662072616e6765000000000000000000000000000081525060200191505060405180910390fd5b600082815481106102c957fe5b90600052602060002001549050919050565b6000808290806001815401808255809150506001900390600052602060002001600090919091909150556001600080549050039050919050565b60008080549050905090565b8060018190555050565b6000600154905090565b6000808054905083101561036757816000848154811061035157fe5b906000526020600020018190555082905061039c565b600082908060018154018082558091505060019003906000526020600020016000909190919091505560016000805490500390505b92915050565b6000806000805490501161041e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f617272617920697320656d70747900000000000000000000000000000000000081525060200191505060405180910390fd5b6000806001600080549050038154811061043457fe5b90600052602060002001549050600080548061044c57fe5b60019003818190600052602060002001600090559055809150509056fea2646970667358221220237bb1feadcde0eea9f11e7c49133914623588ef27644038499423ab948b230e64736f6c63430006000033"

// DeployMiniStore deploys a new Ethereum contract, binding an instance of MiniStore to it.
func DeployMiniStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MiniStore, error) {
	parsed, err := abi.JSON(strings.NewReader(MiniStoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MiniStoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MiniStore{MiniStoreCaller: MiniStoreCaller{contract: contract}, MiniStoreTransactor: MiniStoreTransactor{contract: contract}, MiniStoreFilterer: MiniStoreFilterer{contract: contract}}, nil
}

// MiniStore is an auto generated Go binding around an Ethereum contract.
type MiniStore struct {
	MiniStoreCaller     // Read-only binding to the contract
	MiniStoreTransactor // Write-only binding to the contract
	MiniStoreFilterer   // Log filterer for contract events
}

// MiniStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type MiniStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiniStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MiniStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiniStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MiniStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiniStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MiniStoreSession struct {
	Contract     *MiniStore        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MiniStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MiniStoreCallerSession struct {
	Contract *MiniStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MiniStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MiniStoreTransactorSession struct {
	Contract     *MiniStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MiniStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type MiniStoreRaw struct {
	Contract *MiniStore // Generic contract binding to access the raw methods on
}

// MiniStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MiniStoreCallerRaw struct {
	Contract *MiniStoreCaller // Generic read-only contract binding to access the raw methods on
}

// MiniStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MiniStoreTransactorRaw struct {
	Contract *MiniStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMiniStore creates a new instance of MiniStore, bound to a specific deployed contract.
func NewMiniStore(address common.Address, backend bind.ContractBackend) (*MiniStore, error) {
	contract, err := bindMiniStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MiniStore{MiniStoreCaller: MiniStoreCaller{contract: contract}, MiniStoreTransactor: MiniStoreTransactor{contract: contract}, MiniStoreFilterer: MiniStoreFilterer{contract: contract}}, nil
}

// NewMiniStoreCaller creates a new read-only instance of MiniStore, bound to a specific deployed contract.
func NewMiniStoreCaller(address common.Address, caller bind.ContractCaller) (*MiniStoreCaller, error) {
	contract, err := bindMiniStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MiniStoreCaller{contract: contract}, nil
}

// NewMiniStoreTransactor creates a new write-only instance of MiniStore, bound to a specific deployed contract.
func NewMiniStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*MiniStoreTransactor, error) {
	contract, err := bindMiniStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MiniStoreTransactor{contract: contract}, nil
}

// NewMiniStoreFilterer creates a new log filterer instance of MiniStore, bound to a specific deployed contract.
func NewMiniStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*MiniStoreFilterer, error) {
	contract, err := bindMiniStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MiniStoreFilterer{contract: contract}, nil
}

// bindMiniStore binds a generic wrapper to an already deployed contract.
func bindMiniStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MiniStoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiniStore *MiniStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MiniStore.Contract.MiniStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiniStore *MiniStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiniStore.Contract.MiniStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiniStore *MiniStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiniStore.Contract.MiniStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiniStore *MiniStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MiniStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiniStore *MiniStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiniStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiniStore *MiniStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiniStore.Contract.contract.Transact(opts, method, params...)
}

// GetArrayDataLength is a free data retrieval call binding the contract method 0x69a12799.
//
// Solidity: function getArrayDataLength() view returns(uint256)
func (_MiniStore *MiniStoreCaller) GetArrayDataLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MiniStore.contract.Call(opts, &out, "getArrayDataLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetArrayDataLength is a free data retrieval call binding the contract method 0x69a12799.
//
// Solidity: function getArrayDataLength() view returns(uint256)
func (_MiniStore *MiniStoreSession) GetArrayDataLength() (*big.Int, error) {
	return _MiniStore.Contract.GetArrayDataLength(&_MiniStore.CallOpts)
}

// GetArrayDataLength is a free data retrieval call binding the contract method 0x69a12799.
//
// Solidity: function getArrayDataLength() view returns(uint256)
func (_MiniStore *MiniStoreCallerSession) GetArrayDataLength() (*big.Int, error) {
	return _MiniStore.Contract.GetArrayDataLength(&_MiniStore.CallOpts)
}

// GetArrayLatestValue is a free data retrieval call binding the contract method 0x14d0c25a.
//
// Solidity: function getArrayLatestValue() view returns(uint256)
func (_MiniStore *MiniStoreCaller) GetArrayLatestValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MiniStore.contract.Call(opts, &out, "getArrayLatestValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetArrayLatestValue is a free data retrieval call binding the contract method 0x14d0c25a.
//
// Solidity: function getArrayLatestValue() view returns(uint256)
func (_MiniStore *MiniStoreSession) GetArrayLatestValue() (*big.Int, error) {
	return _MiniStore.Contract.GetArrayLatestValue(&_MiniStore.CallOpts)
}

// GetArrayLatestValue is a free data retrieval call binding the contract method 0x14d0c25a.
//
// Solidity: function getArrayLatestValue() view returns(uint256)
func (_MiniStore *MiniStoreCallerSession) GetArrayLatestValue() (*big.Int, error) {
	return _MiniStore.Contract.GetArrayLatestValue(&_MiniStore.CallOpts)
}

// GetArrayValueAt is a free data retrieval call binding the contract method 0x56298b51.
//
// Solidity: function getArrayValueAt(uint256 i) view returns(uint256)
func (_MiniStore *MiniStoreCaller) GetArrayValueAt(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MiniStore.contract.Call(opts, &out, "getArrayValueAt", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetArrayValueAt is a free data retrieval call binding the contract method 0x56298b51.
//
// Solidity: function getArrayValueAt(uint256 i) view returns(uint256)
func (_MiniStore *MiniStoreSession) GetArrayValueAt(i *big.Int) (*big.Int, error) {
	return _MiniStore.Contract.GetArrayValueAt(&_MiniStore.CallOpts, i)
}

// GetArrayValueAt is a free data retrieval call binding the contract method 0x56298b51.
//
// Solidity: function getArrayValueAt(uint256 i) view returns(uint256)
func (_MiniStore *MiniStoreCallerSession) GetArrayValueAt(i *big.Int) (*big.Int, error) {
	return _MiniStore.Contract.GetArrayValueAt(&_MiniStore.CallOpts, i)
}

// GetNumberValue is a free data retrieval call binding the contract method 0x99ff5f48.
//
// Solidity: function getNumberValue() view returns(uint256)
func (_MiniStore *MiniStoreCaller) GetNumberValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MiniStore.contract.Call(opts, &out, "getNumberValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberValue is a free data retrieval call binding the contract method 0x99ff5f48.
//
// Solidity: function getNumberValue() view returns(uint256)
func (_MiniStore *MiniStoreSession) GetNumberValue() (*big.Int, error) {
	return _MiniStore.Contract.GetNumberValue(&_MiniStore.CallOpts)
}

// GetNumberValue is a free data retrieval call binding the contract method 0x99ff5f48.
//
// Solidity: function getNumberValue() view returns(uint256)
func (_MiniStore *MiniStoreCallerSession) GetNumberValue() (*big.Int, error) {
	return _MiniStore.Contract.GetNumberValue(&_MiniStore.CallOpts)
}

// AddValue is a paid mutator transaction binding the contract method 0x5b9af12b.
//
// Solidity: function addValue(uint256 _value) returns(uint256)
func (_MiniStore *MiniStoreTransactor) AddValue(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _MiniStore.contract.Transact(opts, "addValue", _value)
}

// AddValue is a paid mutator transaction binding the contract method 0x5b9af12b.
//
// Solidity: function addValue(uint256 _value) returns(uint256)
func (_MiniStore *MiniStoreSession) AddValue(_value *big.Int) (*types.Transaction, error) {
	return _MiniStore.Contract.AddValue(&_MiniStore.TransactOpts, _value)
}

// AddValue is a paid mutator transaction binding the contract method 0x5b9af12b.
//
// Solidity: function addValue(uint256 _value) returns(uint256)
func (_MiniStore *MiniStoreTransactorSession) AddValue(_value *big.Int) (*types.Transaction, error) {
	return _MiniStore.Contract.AddValue(&_MiniStore.TransactOpts, _value)
}

// DeleteLastArrayValue is a paid mutator transaction binding the contract method 0xfc7ef48c.
//
// Solidity: function deleteLastArrayValue() returns(uint256)
func (_MiniStore *MiniStoreTransactor) DeleteLastArrayValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiniStore.contract.Transact(opts, "deleteLastArrayValue")
}

// DeleteLastArrayValue is a paid mutator transaction binding the contract method 0xfc7ef48c.
//
// Solidity: function deleteLastArrayValue() returns(uint256)
func (_MiniStore *MiniStoreSession) DeleteLastArrayValue() (*types.Transaction, error) {
	return _MiniStore.Contract.DeleteLastArrayValue(&_MiniStore.TransactOpts)
}

// DeleteLastArrayValue is a paid mutator transaction binding the contract method 0xfc7ef48c.
//
// Solidity: function deleteLastArrayValue() returns(uint256)
func (_MiniStore *MiniStoreTransactorSession) DeleteLastArrayValue() (*types.Transaction, error) {
	return _MiniStore.Contract.DeleteLastArrayValue(&_MiniStore.TransactOpts)
}

// SetArrayValue is a paid mutator transaction binding the contract method 0xa366474f.
//
// Solidity: function setArrayValue(uint256 _i, uint256 _value) returns(uint256)
func (_MiniStore *MiniStoreTransactor) SetArrayValue(opts *bind.TransactOpts, _i *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _MiniStore.contract.Transact(opts, "setArrayValue", _i, _value)
}

// SetArrayValue is a paid mutator transaction binding the contract method 0xa366474f.
//
// Solidity: function setArrayValue(uint256 _i, uint256 _value) returns(uint256)
func (_MiniStore *MiniStoreSession) SetArrayValue(_i *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _MiniStore.Contract.SetArrayValue(&_MiniStore.TransactOpts, _i, _value)
}

// SetArrayValue is a paid mutator transaction binding the contract method 0xa366474f.
//
// Solidity: function setArrayValue(uint256 _i, uint256 _value) returns(uint256)
func (_MiniStore *MiniStoreTransactorSession) SetArrayValue(_i *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _MiniStore.Contract.SetArrayValue(&_MiniStore.TransactOpts, _i, _value)
}

// SetNumberValue is a paid mutator transaction binding the contract method 0x75a01c0f.
//
// Solidity: function setNumberValue(uint256 _value) returns()
func (_MiniStore *MiniStoreTransactor) SetNumberValue(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _MiniStore.contract.Transact(opts, "setNumberValue", _value)
}

// SetNumberValue is a paid mutator transaction binding the contract method 0x75a01c0f.
//
// Solidity: function setNumberValue(uint256 _value) returns()
func (_MiniStore *MiniStoreSession) SetNumberValue(_value *big.Int) (*types.Transaction, error) {
	return _MiniStore.Contract.SetNumberValue(&_MiniStore.TransactOpts, _value)
}

// SetNumberValue is a paid mutator transaction binding the contract method 0x75a01c0f.
//
// Solidity: function setNumberValue(uint256 _value) returns()
func (_MiniStore *MiniStoreTransactorSession) SetNumberValue(_value *big.Int) (*types.Transaction, error) {
	return _MiniStore.Contract.SetNumberValue(&_MiniStore.TransactOpts, _value)
}
