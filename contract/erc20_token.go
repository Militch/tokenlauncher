// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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


// ERC20TokenABI is the input ABI used to generate the binding from.
const ERC20TokenABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_initailSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"superBurnFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20TokenFuncSigs maps the 4-byte function signature to its string representation.
var ERC20TokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"42966c68": "burn(uint256)",
	"79cc6790": "burnFrom(address,uint256)",
	"313ce567": "decimals()",
	"07546172": "minter()",
	"06fdde03": "name()",
	"624c4d72": "superBurnFrom(address,uint256)",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20TokenBin is the compiled bytecode used for deploying new contracts.
var ERC20TokenBin = "0x60806040523480156200001157600080fd5b50604051620010bd380380620010bd83398101604081905262000034916200026a565b83516200004990600090602087019062000111565b5082516200005f90600190602086019062000111565b506002805460ff191660ff8381169190911791829055620000839116600a62000340565b6200008f908362000438565b6003819055600480546001600160a01b03191633178082556001600160a01b0390811660009081526005602052604080822085905592549251929091169290917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91620000ff9190815260200190565b60405180910390a350505050620004c3565b8280546200011f906200045a565b90600052602060002090601f0160209004810192826200014357600085556200018e565b82601f106200015e57805160ff19168380011785556200018e565b828001600101855582156200018e579182015b828111156200018e57825182559160200191906001019062000171565b506200019c929150620001a0565b5090565b5b808211156200019c5760008155600101620001a1565b600082601f830112620001c8578081fd5b81516001600160401b0380821115620001e557620001e5620004ad565b604051601f8301601f19908116603f01168101908282118183101715620002105762000210620004ad565b816040528381526020925086838588010111156200022c578485fd5b8491505b838210156200024f578582018301518183018401529082019062000230565b838211156200026057848385830101525b9695505050505050565b6000806000806080858703121562000280578384fd5b84516001600160401b038082111562000297578586fd5b620002a588838901620001b7565b95506020870151915080821115620002bb578485fd5b50620002ca87828801620001b7565b93505060408501519150606085015160ff81168114620002e8578182fd5b939692955090935050565b80825b600180861162000307575062000337565b8187048211156200031c576200031c62000497565b808616156200032a57918102915b9490941c938002620002f6565b94509492505050565b60006200035460001960ff8516846200035b565b9392505050565b6000826200036c5750600162000354565b816200037b5750600062000354565b81600181146200039457600281146200039f57620003d3565b600191505062000354565b60ff841115620003b357620003b362000497565b6001841b915084821115620003cc57620003cc62000497565b5062000354565b5060208310610133831016604e8410600b84101617156200040b575081810a8381111562000405576200040562000497565b62000354565b6200041a8484846001620002f3565b8086048211156200042f576200042f62000497565b02949350505050565b600081600019048311821515161562000455576200045562000497565b500290565b6002810460018216806200046f57607f821691505b602082108114156200049157634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b610bea80620004d36000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806342966c681161008c57806379cc67901161006657806379cc6790146101d357806395d89b41146101e6578063a9059cbb146101ee578063dd62ed3e14610201576100cf565b806342966c6814610184578063624c4d721461019757806370a08231146101aa576100cf565b806306fdde03146100d457806307546172146100f2578063095ea7b31461011d57806318160ddd1461014057806323b872dd14610152578063313ce56714610165575b600080fd5b6100dc61023a565b6040516100e99190610a25565b60405180910390f35b600454610105906001600160a01b031681565b6040516001600160a01b0390911681526020016100e9565b61013061012b3660046109e4565b6102c8565b60405190151581526020016100e9565b6003545b6040519081526020016100e9565b6101306101603660046109a9565b610334565b6002546101729060ff1681565b60405160ff90911681526020016100e9565b610130610192366004610a0d565b610506565b6101306101a53660046109e4565b6105a6565b6101446101b836600461095d565b6001600160a01b031660009081526005602052604090205490565b6101306101e13660046109e4565b6106aa565b6100dc6107d5565b6101306101fc3660046109e4565b6107e2565b61014461020f366004610977565b6001600160a01b03918216600090815260066020908152604080832093909416825291909152205490565b6000805461024790610b63565b80601f016020809104026020016040519081016040528092919081815260200182805461027390610b63565b80156102c05780601f10610295576101008083540402835291602001916102c0565b820191906000526020600020905b8154815290600101906020018083116102a357829003601f168201915b505050505081565b3360008181526006602090815260408083206001600160a01b038716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103239086815260200190565b60405180910390a350600192915050565b6001600160a01b0383166000908152600560205260408120548211156103755760405162461bcd60e51b815260040161036c90610a78565b60405180910390fd5b6001600160a01b03841660009081526006602090815260408083203384529091529020548211156103b85760405162461bcd60e51b815260040161036c90610aee565b6001600160a01b03831661040e5760405162461bcd60e51b815260206004820181905260248201527f43616e6e6f742073656e6420746f20616c6c207a65726f20616464726573732e604482015260640161036c565b6001600160a01b0384166000908152600560205260409020546104319083610900565b6001600160a01b0380861660009081526005602052604080822093909355908516815220546104609083610923565b6001600160a01b03808516600090815260056020908152604080832094909455918716815260068252828120338252909152205461049e9083610900565b6001600160a01b03858116600081815260066020908152604080832033845282529182902094909455518581529186169290917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35060019392505050565b336000908152600560205260408120548211156105355760405162461bcd60e51b815260040161036c90610aad565b3360009081526005602052604090205461054f9083610900565b33600081815260056020526040908190209290925590517f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7906105959085815260200190565b60405180910390a25060015b919050565b6001600160a01b0382166000908152600560205260408120548211156105de5760405162461bcd60e51b815260040161036c90610a78565b6004546001600160a01b031633146106275760405162461bcd60e51b815260206004820152600c60248201526b155b985d5d1a1bdc9a5e995960a21b604482015260640161036c565b6001600160a01b03831660009081526005602052604090205461064a9083610900565b6001600160a01b038416600081815260056020526040908190209290925590517f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7906106999085815260200190565b60405180910390a250600192915050565b6001600160a01b0382166000908152600560205260408120548211156106e25760405162461bcd60e51b815260040161036c90610a78565b6001600160a01b03831660009081526006602090815260408083203384529091529020548211156107255760405162461bcd60e51b815260040161036c90610aee565b6001600160a01b0383166000908152600560205260409020546107489083610900565b6001600160a01b038416600090815260056020908152604080832093909355600681528282203383529052205461077f9083610900565b6001600160a01b0384166000818152600660209081526040808320338452825291829020939093555184815290917f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df79101610699565b6001805461024790610b63565b60006001600160a01b03831661083a5760405162461bcd60e51b815260206004820181905260248201527f43616e6e6f742073656e6420746f20616c6c207a65726f20616464726573732e604482015260640161036c565b336000908152600560205260409020548211156108695760405162461bcd60e51b815260040161036c90610aad565b336000908152600560205260409020546108839083610900565b33600090815260056020526040808220929092556001600160a01b038516815220546108af9083610923565b6001600160a01b0384166000818152600560205260409081902092909255905133907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906103239086815260200190565b60008282111561090f57600080fd5b600061091b8385610b4c565b949350505050565b6000806109308385610b34565b90508381101561093f57600080fd5b9392505050565b80356001600160a01b03811681146105a157600080fd5b60006020828403121561096e578081fd5b61093f82610946565b60008060408385031215610989578081fd5b61099283610946565b91506109a060208401610946565b90509250929050565b6000806000606084860312156109bd578081fd5b6109c684610946565b92506109d460208501610946565b9150604084013590509250925092565b600080604083850312156109f6578182fd5b6109ff83610946565b946020939093013593505050565b600060208284031215610a1e578081fd5b5035919050565b6000602080835283518082850152825b81811015610a5157858101830151858201604001528201610a35565b81811115610a625783604083870101525b50601f01601f1916929092016040019392505050565b6020808252818101527f66726f6d20646f65736e74206861766520656e6f7567682062616c616e63652e604082015260600190565b60208082526021908201527f6d73672e73656e6465722062616c616e6365206973206e6f7420656e6f7567686040820152601760f91b606082015260800190565b60208082526026908201527f416c6c6f77616e6365206f66206d73672e73656e646572206973206e6f7420656040820152653737bab3b41760d11b606082015260800190565b60008219821115610b4757610b47610b9e565b500190565b600082821015610b5e57610b5e610b9e565b500390565b600281046001821680610b7757607f821691505b60208210811415610b9857634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fdfea26469706673582212203260f93f48cd0917060d56d21bdc4fa81c35fdb555cbbc7ab6930cd29f0af51164736f6c63430008020033"

// DeployERC20Token deploys a new Ethereum contract, binding an instance of ERC20Token to it.
func DeployERC20Token(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _initailSupply *big.Int, _decimals uint8) (common.Address, *types.Transaction, *ERC20Token, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20TokenBin), backend, _name, _symbol, _initailSupply, _decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Token{ERC20TokenCaller: ERC20TokenCaller{contract: contract}, ERC20TokenTransactor: ERC20TokenTransactor{contract: contract}, ERC20TokenFilterer: ERC20TokenFilterer{contract: contract}}, nil
}

// ERC20Token is an auto generated Go binding around an Ethereum contract.
type ERC20Token struct {
	ERC20TokenCaller     // Read-only binding to the contract
	ERC20TokenTransactor // Write-only binding to the contract
	ERC20TokenFilterer   // Log filterer for contract events
}

// ERC20TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TokenSession struct {
	Contract     *ERC20Token       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TokenCallerSession struct {
	Contract *ERC20TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TokenTransactorSession struct {
	Contract     *ERC20TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TokenRaw struct {
	Contract *ERC20Token // Generic contract binding to access the raw methods on
}

// ERC20TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TokenCallerRaw struct {
	Contract *ERC20TokenCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TokenTransactorRaw struct {
	Contract *ERC20TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Token creates a new instance of ERC20Token, bound to a specific deployed contract.
func NewERC20Token(address common.Address, backend bind.ContractBackend) (*ERC20Token, error) {
	contract, err := bindERC20Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Token{ERC20TokenCaller: ERC20TokenCaller{contract: contract}, ERC20TokenTransactor: ERC20TokenTransactor{contract: contract}, ERC20TokenFilterer: ERC20TokenFilterer{contract: contract}}, nil
}

// NewERC20TokenCaller creates a new read-only instance of ERC20Token, bound to a specific deployed contract.
func NewERC20TokenCaller(address common.Address, caller bind.ContractCaller) (*ERC20TokenCaller, error) {
	contract, err := bindERC20Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenCaller{contract: contract}, nil
}

// NewERC20TokenTransactor creates a new write-only instance of ERC20Token, bound to a specific deployed contract.
func NewERC20TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TokenTransactor, error) {
	contract, err := bindERC20Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenTransactor{contract: contract}, nil
}

// NewERC20TokenFilterer creates a new log filterer instance of ERC20Token, bound to a specific deployed contract.
func NewERC20TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TokenFilterer, error) {
	contract, err := bindERC20Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenFilterer{contract: contract}, nil
}

// bindERC20Token binds a generic wrapper to an already deployed contract.
func bindERC20Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Token *ERC20TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Token.Contract.ERC20TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Token *ERC20TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Token.Contract.ERC20TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Token *ERC20TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Token.Contract.ERC20TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Token *ERC20TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Token *ERC20TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Token *ERC20TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Token *ERC20TokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Token.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Token *ERC20TokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Token.Contract.Allowance(&_ERC20Token.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Token *ERC20TokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Token.Contract.Allowance(&_ERC20Token.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC20Token *ERC20TokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Token.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC20Token *ERC20TokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC20Token.Contract.BalanceOf(&_ERC20Token.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC20Token *ERC20TokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC20Token.Contract.BalanceOf(&_ERC20Token.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Token *ERC20TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Token.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Token *ERC20TokenSession) Decimals() (uint8, error) {
	return _ERC20Token.Contract.Decimals(&_ERC20Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Token *ERC20TokenCallerSession) Decimals() (uint8, error) {
	return _ERC20Token.Contract.Decimals(&_ERC20Token.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ERC20Token *ERC20TokenCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Token.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ERC20Token *ERC20TokenSession) Minter() (common.Address, error) {
	return _ERC20Token.Contract.Minter(&_ERC20Token.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ERC20Token *ERC20TokenCallerSession) Minter() (common.Address, error) {
	return _ERC20Token.Contract.Minter(&_ERC20Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Token *ERC20TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Token *ERC20TokenSession) Name() (string, error) {
	return _ERC20Token.Contract.Name(&_ERC20Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Token *ERC20TokenCallerSession) Name() (string, error) {
	return _ERC20Token.Contract.Name(&_ERC20Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Token *ERC20TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Token *ERC20TokenSession) Symbol() (string, error) {
	return _ERC20Token.Contract.Symbol(&_ERC20Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Token *ERC20TokenCallerSession) Symbol() (string, error) {
	return _ERC20Token.Contract.Symbol(&_ERC20Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Token *ERC20TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Token *ERC20TokenSession) TotalSupply() (*big.Int, error) {
	return _ERC20Token.Contract.TotalSupply(&_ERC20Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Token *ERC20TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Token.Contract.TotalSupply(&_ERC20Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.Approve(&_ERC20Token.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.Approve(&_ERC20Token.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.Burn(&_ERC20Token.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.Burn(&_ERC20Token.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenTransactor) BurnFrom(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.contract.Transact(opts, "burnFrom", from, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenSession) BurnFrom(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.BurnFrom(&_ERC20Token.TransactOpts, from, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenTransactorSession) BurnFrom(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.BurnFrom(&_ERC20Token.TransactOpts, from, value)
}

// SuperBurnFrom is a paid mutator transaction binding the contract method 0x624c4d72.
//
// Solidity: function superBurnFrom(address from, uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenTransactor) SuperBurnFrom(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.contract.Transact(opts, "superBurnFrom", from, value)
}

// SuperBurnFrom is a paid mutator transaction binding the contract method 0x624c4d72.
//
// Solidity: function superBurnFrom(address from, uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenSession) SuperBurnFrom(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.SuperBurnFrom(&_ERC20Token.TransactOpts, from, value)
}

// SuperBurnFrom is a paid mutator transaction binding the contract method 0x624c4d72.
//
// Solidity: function superBurnFrom(address from, uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenTransactorSession) SuperBurnFrom(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.SuperBurnFrom(&_ERC20Token.TransactOpts, from, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.Transfer(&_ERC20Token.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.Transfer(&_ERC20Token.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.TransferFrom(&_ERC20Token.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20Token *ERC20TokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Token.Contract.TransferFrom(&_ERC20Token.TransactOpts, from, to, value)
}

// ERC20TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Token contract.
type ERC20TokenApprovalIterator struct {
	Event *ERC20TokenApproval // Event containing the contract specifics and raw log

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
func (it *ERC20TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenApproval)
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
		it.Event = new(ERC20TokenApproval)
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
func (it *ERC20TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenApproval represents a Approval event raised by the ERC20Token contract.
type ERC20TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Token *ERC20TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenApprovalIterator{contract: _ERC20Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Token *ERC20TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenApproval)
				if err := _ERC20Token.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20Token *ERC20TokenFilterer) ParseApproval(log types.Log) (*ERC20TokenApproval, error) {
	event := new(ERC20TokenApproval)
	if err := _ERC20Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the ERC20Token contract.
type ERC20TokenBurnedIterator struct {
	Event *ERC20TokenBurned // Event containing the contract specifics and raw log

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
func (it *ERC20TokenBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenBurned)
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
		it.Event = new(ERC20TokenBurned)
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
func (it *ERC20TokenBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenBurned represents a Burned event raised by the ERC20Token contract.
type ERC20TokenBurned struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed from, uint256 value)
func (_ERC20Token *ERC20TokenFilterer) FilterBurned(opts *bind.FilterOpts, from []common.Address) (*ERC20TokenBurnedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ERC20Token.contract.FilterLogs(opts, "Burned", fromRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenBurnedIterator{contract: _ERC20Token.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed from, uint256 value)
func (_ERC20Token *ERC20TokenFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *ERC20TokenBurned, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ERC20Token.contract.WatchLogs(opts, "Burned", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenBurned)
				if err := _ERC20Token.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed from, uint256 value)
func (_ERC20Token *ERC20TokenFilterer) ParseBurned(log types.Log) (*ERC20TokenBurned, error) {
	event := new(ERC20TokenBurned)
	if err := _ERC20Token.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Token contract.
type ERC20TokenTransferIterator struct {
	Event *ERC20TokenTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenTransfer)
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
		it.Event = new(ERC20TokenTransfer)
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
func (it *ERC20TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenTransfer represents a Transfer event raised by the ERC20Token contract.
type ERC20TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Token *ERC20TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenTransferIterator{contract: _ERC20Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Token *ERC20TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenTransfer)
				if err := _ERC20Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20Token *ERC20TokenFilterer) ParseTransfer(log types.Log) (*ERC20TokenTransfer, error) {
	event := new(ERC20TokenTransfer)
	if err := _ERC20Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}


