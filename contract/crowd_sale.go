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

// CrowdSaleABI is the input ABI used to generate the binding from.
const CrowdSaleABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_targetFunds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"raise\",\"type\":\"uint256\"}],\"name\":\"Bought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"raise\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"raise\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundsRaisedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundsSoldTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSalesCompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isUncompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"setEndTime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"setStartTime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setTargetFunds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// CrowdSaleBin is the compiled bytecode used for deploying new contracts.
var CrowdSaleBin = "0x608060405260006001556000600255600060045560006005553480156200002557600080fd5b506040516200144c3803806200144c83398101604081905262000048916200024c565b600080546001600160a01b03199081163317909155600380546001600160a01b038881169190931617908190556040805163313ce56760e01b81529051919092169163313ce567916004808301926020929190829003018186803b158015620000b057600080fd5b505afa158015620000c5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000eb9190620002a1565b620000fb9060ff16600a62000318565b62000107908562000400565b600555600483905581620001715760405162461bcd60e51b815260206004820152602660248201527f737461727454696d65206d7573742062652067726561746572207468616e20746044820152656f207a65726f60d01b60648201526084015b60405180910390fd5b60008111620001cf5760405162461bcd60e51b8152602060048201526024808201527f656e6454696d65206d7573742062652067726561746572207468616e20746f206044820152637a65726f60e01b606482015260840162000168565b80821115620002395760405162461bcd60e51b815260206004820152602f60248201527f737461727454696d65206d757374206265206c657373207468616e206f72206560448201526e7175616c20746f20656e6454696d6560881b606482015260840162000168565b6006919091556007555062000438915050565b600080600080600060a0868803121562000264578081fd5b85516001600160a01b03811681146200027b578182fd5b602087015160408801516060890151608090990151929a91995097965090945092505050565b600060208284031215620002b3578081fd5b815160ff81168114620002c4578182fd5b9392505050565b80825b6001808611620002df57506200030f565b818704821115620002f457620002f462000422565b808616156200030257918102915b9490941c938002620002ce565b94509492505050565b6000620002c460001984846000826200033457506001620002c4565b816200034357506000620002c4565b81600181146200035c576002811462000367576200039b565b6001915050620002c4565b60ff8411156200037b576200037b62000422565b6001841b91508482111562000394576200039462000422565b50620002c4565b5060208310610133831016604e8410600b8410161715620003d3575081810a83811115620003cd57620003cd62000422565b620002c4565b620003e28484846001620002cb565b808604821115620003f757620003f762000422565b02949350505050565b60008160001904831182151516156200041d576200041d62000422565b500290565b634e487b7160e01b600052601160045260246000fd5b61100480620004486000396000f3fe60806040526004361061012e5760003560e01c806391b7f5ed116100ab578063db8c69ef1161006f578063db8c69ef146102ea578063e2481cce14610300578063eeeb365014610315578063f2fde38b1461032b578063f756017a1461034b578063fc0c546a1461036b5761013d565b806391b7f5ed14610272578063a035b1fe14610292578063a06c944a146102a8578063a6f2ae3a146102c2578063ccb98ffc146102ca5761013d565b80633e0a322d116100f25780633e0a322d146101da578063590e1ae3146101fa57806378e979251461020f5780637b352962146102255780638da5cb5b1461023a5761013d565b80631459e2ad1461014c5780631a39d8ef1461017557806322f3e2d41461018a5780633197cbb6146101af5780633ccfd60b146101c55761013d565b3661013d5761013b61038b565b005b34801561014957600080fd5b50005b34801561015857600080fd5b5061016260055481565b6040519081526020015b60405180910390f35b34801561018157600080fd5b506101626106e9565b34801561019657600080fd5b5061019f61076a565b604051901515815260200161016c565b3480156101bb57600080fd5b5061016260075481565b3480156101d157600080fd5b5061019f61079b565b3480156101e657600080fd5b5061019f6101f5366004610d84565b6108a1565b34801561020657600080fd5b5061019f61095a565b34801561021b57600080fd5b5061016260065481565b34801561023157600080fd5b5061019f610ab4565b34801561024657600080fd5b5060005461025a906001600160a01b031681565b6040516001600160a01b03909116815260200161016c565b34801561027e57600080fd5b5061019f61028d366004610d84565b610ad2565b34801561029e57600080fd5b5061016260045481565b3480156102b457600080fd5b50600554600154101561019f565b61013b61038b565b3480156102d657600080fd5b5061019f6102e5366004610d84565b610b06565b3480156102f657600080fd5b5061016260015481565b34801561030c57600080fd5b5061019f610bc4565b34801561032157600080fd5b5061016260025481565b34801561033757600080fd5b5061013b610346366004610d3d565b610be2565b34801561035757600080fd5b5061019f610366366004610d84565b610c90565b34801561037757600080fd5b5060035461025a906001600160a01b031681565b61039361076a565b6103e45760405162461bcd60e51b815260206004820152601860248201527f43726f77642073616c65206973206e6f7420616374697665000000000000000060448201526064015b60405180910390fd5b6003546040805163313ce56760e01b8152905134926000926001600160a01b039091169163313ce56791600480820192602092909190829003018186803b15801561042e57600080fd5b505afa158015610442573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104669190610db4565b60ff16905061048d61047982600a610ecb565b600454610487908590610cc4565b90610ce8565b6003546040516370a0823160e01b81523060048201529193506000916001600160a01b03909116906370a082319060240160206040518083038186803b1580156104d657600080fd5b505afa1580156104ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061050e9190610d9c565b9050600083116105605760405162461bcd60e51b815260206004820152601b60248201527f596f75206e65656420746f2073656e6420736f6d65206574686572000000000060448201526064016103db565b808311156105b05760405162461bcd60e51b815260206004820181905260248201527f4e6f7420656e6f75676820746f6b656e7320696e20746865207265736572766560448201526064016103db565b60035460405163a9059cbb60e01b8152336004820152602481018590526001600160a01b039091169063a9059cbb90604401602060405180830381600087803b1580156105fc57600080fd5b505af1158015610610573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106349190610d64565b506001546106429084610d21565b6001556002546106529034610d21565b6002553360009081526008602052604090205461066f9034610d21565b336000908152600860209081526040808320939093556009905220546106959084610d21565b3360008181526009602090815260409182902093909355805134815292830186905290917fa9a40dec7a304e5915d11358b968c1e8d365992abf20f82285d1df1b30c8e24c910160405180910390a2505050565b6003546040516370a0823160e01b81523060048201526000916001600160a01b0316906370a082319060240160206040518083038186803b15801561072d57600080fd5b505afa158015610741573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107659190610d9c565b905090565b6006546000904290811080159061078357506007548111155b8015610795575060055460015410155b155b91505090565b600080546001600160a01b031633146107c65760405162461bcd60e51b81526004016103db90610dd5565b6107ce610ab4565b61081a5760405162461bcd60e51b815260206004820152601a60248201527f43726f77642073616c65206973206e6f742066696e697368656400000000000060448201526064016103db565b600080546002546040516001600160a01b039092169281156108fc029290818181858888f19350505050158015610855573d6000803e3d6000fd5b506000546002546040519081526001600160a01b03909116907f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d59060200160405180910390a250600190565b600080546001600160a01b031633146108cc5760405162461bcd60e51b81526004016103db90610dd5565b600082116108ec5760405162461bcd60e51b81526004016103db90610e0c565b6007548211156109515760405162461bcd60e51b815260206004820152602a60248201527f74696d65206d757374206265206c657373207468616e206f7220657175616c20604482015269746f20656e6454696d6560b01b60648201526084016103db565b50600655600190565b6000610964610bc4565b6109b05760405162461bcd60e51b815260206004820152601960248201527f43726f77642073616c6520697320756e636f6d706c657465640000000000000060448201526064016103db565b3360009081526008602090815260408083205460099092529091205481610a275760405162461bcd60e51b815260206004820152602560248201527f596f75206e65656420746f2073656c6c206174206c6561737420736f6d6520746044820152646f6b656e7360d81b60648201526084016103db565b604051339083156108fc029084906000818181858888f19350505050158015610a54573d6000803e3d6000fd5b50336000818152600860209081526040808320839055600982528083209290925581518581529081018490527f2dc8e290002f06fc0085bbca9dfb8b415cf4d1178950c72ff9ee8f4d8878ee66910160405180910390a260019250505090565b60075460009042908111801561079557506005546001541015610795565b600080546001600160a01b03163314610afd5760405162461bcd60e51b81526004016103db90610dd5565b50600455600190565b600080546001600160a01b03163314610b315760405162461bcd60e51b81526004016103db90610dd5565b60008211610b515760405162461bcd60e51b81526004016103db90610e0c565b600654821015610bbb5760405162461bcd60e51b815260206004820152602f60248201527f74696d65206d7573742062652067726561746572207468616e206f722065717560448201526e616c20746f20737461727454696d6560881b60648201526084016103db565b50600755600190565b60075460009042908111801561079557506005546001541015610793565b6000546001600160a01b03163314610c0c5760405162461bcd60e51b81526004016103db90610dd5565b6001600160a01b038116610c6e5760405162461bcd60e51b8152602060048201526024808201527f43616e6e6f74207472616e7366657220746f20616c6c207a65726f206164647260448201526332b9b99760e11b60648201526084016103db565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b600080546001600160a01b03163314610cbb5760405162461bcd60e51b81526004016103db90610dd5565b50600555600190565b6000808211610cd257600080fd5b6000610cde8385610e65565b9150505b92915050565b600082610cf757506000610ce2565b6000610d038385610f99565b905082610d108583610e65565b14610d1a57600080fd5b9392505050565b600080610d2e8385610e4d565b905083811015610d1a57600080fd5b600060208284031215610d4e578081fd5b81356001600160a01b0381168114610d1a578182fd5b600060208284031215610d75578081fd5b81518015158114610d1a578182fd5b600060208284031215610d95578081fd5b5035919050565b600060208284031215610dad578081fd5b5051919050565b600060208284031215610dc5578081fd5b815160ff81168114610d1a578182fd5b60208082526019908201527f4f6e6c79206f776e65722063616e2063616c6c20746869732e00000000000000604082015260600190565b60208082526021908201527f74696d65206d7573742062652067726561746572207468616e20746f207a65726040820152606f60f81b606082015260800190565b60008219821115610e6057610e60610fb8565b500190565b600082610e8057634e487b7160e01b81526012600452602481fd5b500490565b80825b6001808611610e975750610ec2565b818704821115610ea957610ea9610fb8565b80861615610eb657918102915b9490941c938002610e88565b94509492505050565b6000610d1a6000198484600082610ee457506001610d1a565b81610ef157506000610d1a565b8160018114610f075760028114610f1157610f3e565b6001915050610d1a565b60ff841115610f2257610f22610fb8565b6001841b915084821115610f3857610f38610fb8565b50610d1a565b5060208310610133831016604e8410600b8410161715610f71575081810a83811115610f6c57610f6c610fb8565b610d1a565b610f7e8484846001610e85565b808604821115610f9057610f90610fb8565b02949350505050565b6000816000190483118215151615610fb357610fb3610fb8565b500290565b634e487b7160e01b600052601160045260246000fdfea26469706673582212204cbef87f4b6d6922268e4ba701e2c8b23f1044aa5b95a92cc9317a11c190d4cd64736f6c63430008020033"

// DeployCrowdSale deploys a new Ethereum contract, binding an instance of CrowdSale to it.
func DeployCrowdSale(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddress common.Address, _targetFunds *big.Int, _price *big.Int, _startTime *big.Int, _endTime *big.Int) (common.Address, *types.Transaction, *CrowdSale, error) {
	parsed, err := abi.JSON(strings.NewReader(CrowdSaleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CrowdSaleBin), backend, tokenAddress, _targetFunds, _price, _startTime, _endTime)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrowdSale{CrowdSaleCaller: CrowdSaleCaller{contract: contract}, CrowdSaleTransactor: CrowdSaleTransactor{contract: contract}, CrowdSaleFilterer: CrowdSaleFilterer{contract: contract}}, nil
}

// CrowdSale is an auto generated Go binding around an Ethereum contract.
type CrowdSale struct {
	CrowdSaleCaller     // Read-only binding to the contract
	CrowdSaleTransactor // Write-only binding to the contract
	CrowdSaleFilterer   // Log filterer for contract events
}

// CrowdSaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrowdSaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdSaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrowdSaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdSaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrowdSaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdSaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrowdSaleSession struct {
	Contract     *CrowdSale        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrowdSaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrowdSaleCallerSession struct {
	Contract *CrowdSaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CrowdSaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrowdSaleTransactorSession struct {
	Contract     *CrowdSaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CrowdSaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrowdSaleRaw struct {
	Contract *CrowdSale // Generic contract binding to access the raw methods on
}

// CrowdSaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrowdSaleCallerRaw struct {
	Contract *CrowdSaleCaller // Generic read-only contract binding to access the raw methods on
}

// CrowdSaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrowdSaleTransactorRaw struct {
	Contract *CrowdSaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrowdSale creates a new instance of CrowdSale, bound to a specific deployed contract.
func NewCrowdSale(address common.Address, backend bind.ContractBackend) (*CrowdSale, error) {
	contract, err := bindCrowdSale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrowdSale{CrowdSaleCaller: CrowdSaleCaller{contract: contract}, CrowdSaleTransactor: CrowdSaleTransactor{contract: contract}, CrowdSaleFilterer: CrowdSaleFilterer{contract: contract}}, nil
}

// NewCrowdSaleCaller creates a new read-only instance of CrowdSale, bound to a specific deployed contract.
func NewCrowdSaleCaller(address common.Address, caller bind.ContractCaller) (*CrowdSaleCaller, error) {
	contract, err := bindCrowdSale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdSaleCaller{contract: contract}, nil
}

// NewCrowdSaleTransactor creates a new write-only instance of CrowdSale, bound to a specific deployed contract.
func NewCrowdSaleTransactor(address common.Address, transactor bind.ContractTransactor) (*CrowdSaleTransactor, error) {
	contract, err := bindCrowdSale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdSaleTransactor{contract: contract}, nil
}

// NewCrowdSaleFilterer creates a new log filterer instance of CrowdSale, bound to a specific deployed contract.
func NewCrowdSaleFilterer(address common.Address, filterer bind.ContractFilterer) (*CrowdSaleFilterer, error) {
	contract, err := bindCrowdSale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrowdSaleFilterer{contract: contract}, nil
}

// bindCrowdSale binds a generic wrapper to an already deployed contract.
func bindCrowdSale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrowdSaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrowdSale *CrowdSaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrowdSale.Contract.CrowdSaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrowdSale *CrowdSaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdSale.Contract.CrowdSaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrowdSale *CrowdSaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrowdSale.Contract.CrowdSaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrowdSale *CrowdSaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrowdSale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrowdSale *CrowdSaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdSale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrowdSale *CrowdSaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrowdSale.Contract.contract.Transact(opts, method, params...)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_CrowdSale *CrowdSaleCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_CrowdSale *CrowdSaleSession) EndTime() (*big.Int, error) {
	return _CrowdSale.Contract.EndTime(&_CrowdSale.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_CrowdSale *CrowdSaleCallerSession) EndTime() (*big.Int, error) {
	return _CrowdSale.Contract.EndTime(&_CrowdSale.CallOpts)
}

// FundsRaisedTotal is a free data retrieval call binding the contract method 0xeeeb3650.
//
// Solidity: function fundsRaisedTotal() view returns(uint256)
func (_CrowdSale *CrowdSaleCaller) FundsRaisedTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "fundsRaisedTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundsRaisedTotal is a free data retrieval call binding the contract method 0xeeeb3650.
//
// Solidity: function fundsRaisedTotal() view returns(uint256)
func (_CrowdSale *CrowdSaleSession) FundsRaisedTotal() (*big.Int, error) {
	return _CrowdSale.Contract.FundsRaisedTotal(&_CrowdSale.CallOpts)
}

// FundsRaisedTotal is a free data retrieval call binding the contract method 0xeeeb3650.
//
// Solidity: function fundsRaisedTotal() view returns(uint256)
func (_CrowdSale *CrowdSaleCallerSession) FundsRaisedTotal() (*big.Int, error) {
	return _CrowdSale.Contract.FundsRaisedTotal(&_CrowdSale.CallOpts)
}

// FundsSoldTotal is a free data retrieval call binding the contract method 0xdb8c69ef.
//
// Solidity: function fundsSoldTotal() view returns(uint256)
func (_CrowdSale *CrowdSaleCaller) FundsSoldTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "fundsSoldTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundsSoldTotal is a free data retrieval call binding the contract method 0xdb8c69ef.
//
// Solidity: function fundsSoldTotal() view returns(uint256)
func (_CrowdSale *CrowdSaleSession) FundsSoldTotal() (*big.Int, error) {
	return _CrowdSale.Contract.FundsSoldTotal(&_CrowdSale.CallOpts)
}

// FundsSoldTotal is a free data retrieval call binding the contract method 0xdb8c69ef.
//
// Solidity: function fundsSoldTotal() view returns(uint256)
func (_CrowdSale *CrowdSaleCallerSession) FundsSoldTotal() (*big.Int, error) {
	return _CrowdSale.Contract.FundsSoldTotal(&_CrowdSale.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_CrowdSale *CrowdSaleCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_CrowdSale *CrowdSaleSession) IsActive() (bool, error) {
	return _CrowdSale.Contract.IsActive(&_CrowdSale.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_CrowdSale *CrowdSaleCallerSession) IsActive() (bool, error) {
	return _CrowdSale.Contract.IsActive(&_CrowdSale.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() view returns(bool)
func (_CrowdSale *CrowdSaleCaller) IsFinished(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "isFinished")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() view returns(bool)
func (_CrowdSale *CrowdSaleSession) IsFinished() (bool, error) {
	return _CrowdSale.Contract.IsFinished(&_CrowdSale.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() view returns(bool)
func (_CrowdSale *CrowdSaleCallerSession) IsFinished() (bool, error) {
	return _CrowdSale.Contract.IsFinished(&_CrowdSale.CallOpts)
}

// IsSalesCompleted is a free data retrieval call binding the contract method 0xa06c944a.
//
// Solidity: function isSalesCompleted() view returns(bool)
func (_CrowdSale *CrowdSaleCaller) IsSalesCompleted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "isSalesCompleted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSalesCompleted is a free data retrieval call binding the contract method 0xa06c944a.
//
// Solidity: function isSalesCompleted() view returns(bool)
func (_CrowdSale *CrowdSaleSession) IsSalesCompleted() (bool, error) {
	return _CrowdSale.Contract.IsSalesCompleted(&_CrowdSale.CallOpts)
}

// IsSalesCompleted is a free data retrieval call binding the contract method 0xa06c944a.
//
// Solidity: function isSalesCompleted() view returns(bool)
func (_CrowdSale *CrowdSaleCallerSession) IsSalesCompleted() (bool, error) {
	return _CrowdSale.Contract.IsSalesCompleted(&_CrowdSale.CallOpts)
}

// IsUncompleted is a free data retrieval call binding the contract method 0xe2481cce.
//
// Solidity: function isUncompleted() view returns(bool)
func (_CrowdSale *CrowdSaleCaller) IsUncompleted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "isUncompleted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUncompleted is a free data retrieval call binding the contract method 0xe2481cce.
//
// Solidity: function isUncompleted() view returns(bool)
func (_CrowdSale *CrowdSaleSession) IsUncompleted() (bool, error) {
	return _CrowdSale.Contract.IsUncompleted(&_CrowdSale.CallOpts)
}

// IsUncompleted is a free data retrieval call binding the contract method 0xe2481cce.
//
// Solidity: function isUncompleted() view returns(bool)
func (_CrowdSale *CrowdSaleCallerSession) IsUncompleted() (bool, error) {
	return _CrowdSale.Contract.IsUncompleted(&_CrowdSale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrowdSale *CrowdSaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrowdSale *CrowdSaleSession) Owner() (common.Address, error) {
	return _CrowdSale.Contract.Owner(&_CrowdSale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrowdSale *CrowdSaleCallerSession) Owner() (common.Address, error) {
	return _CrowdSale.Contract.Owner(&_CrowdSale.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_CrowdSale *CrowdSaleCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_CrowdSale *CrowdSaleSession) Price() (*big.Int, error) {
	return _CrowdSale.Contract.Price(&_CrowdSale.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_CrowdSale *CrowdSaleCallerSession) Price() (*big.Int, error) {
	return _CrowdSale.Contract.Price(&_CrowdSale.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_CrowdSale *CrowdSaleCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_CrowdSale *CrowdSaleSession) StartTime() (*big.Int, error) {
	return _CrowdSale.Contract.StartTime(&_CrowdSale.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_CrowdSale *CrowdSaleCallerSession) StartTime() (*big.Int, error) {
	return _CrowdSale.Contract.StartTime(&_CrowdSale.CallOpts)
}

// TargetFunds is a free data retrieval call binding the contract method 0x1459e2ad.
//
// Solidity: function targetFunds() view returns(uint256)
func (_CrowdSale *CrowdSaleCaller) TargetFunds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "targetFunds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetFunds is a free data retrieval call binding the contract method 0x1459e2ad.
//
// Solidity: function targetFunds() view returns(uint256)
func (_CrowdSale *CrowdSaleSession) TargetFunds() (*big.Int, error) {
	return _CrowdSale.Contract.TargetFunds(&_CrowdSale.CallOpts)
}

// TargetFunds is a free data retrieval call binding the contract method 0x1459e2ad.
//
// Solidity: function targetFunds() view returns(uint256)
func (_CrowdSale *CrowdSaleCallerSession) TargetFunds() (*big.Int, error) {
	return _CrowdSale.Contract.TargetFunds(&_CrowdSale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CrowdSale *CrowdSaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CrowdSale *CrowdSaleSession) Token() (common.Address, error) {
	return _CrowdSale.Contract.Token(&_CrowdSale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CrowdSale *CrowdSaleCallerSession) Token() (common.Address, error) {
	return _CrowdSale.Contract.Token(&_CrowdSale.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_CrowdSale *CrowdSaleCaller) TotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "totalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_CrowdSale *CrowdSaleSession) TotalAmount() (*big.Int, error) {
	return _CrowdSale.Contract.TotalAmount(&_CrowdSale.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_CrowdSale *CrowdSaleCallerSession) TotalAmount() (*big.Int, error) {
	return _CrowdSale.Contract.TotalAmount(&_CrowdSale.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_CrowdSale *CrowdSaleTransactor) Buy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdSale.contract.Transact(opts, "buy")
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_CrowdSale *CrowdSaleSession) Buy() (*types.Transaction, error) {
	return _CrowdSale.Contract.Buy(&_CrowdSale.TransactOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_CrowdSale *CrowdSaleTransactorSession) Buy() (*types.Transaction, error) {
	return _CrowdSale.Contract.Buy(&_CrowdSale.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns(bool)
func (_CrowdSale *CrowdSaleTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdSale.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns(bool)
func (_CrowdSale *CrowdSaleSession) Refund() (*types.Transaction, error) {
	return _CrowdSale.Contract.Refund(&_CrowdSale.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns(bool)
func (_CrowdSale *CrowdSaleTransactorSession) Refund() (*types.Transaction, error) {
	return _CrowdSale.Contract.Refund(&_CrowdSale.TransactOpts)
}

// SetEndTime is a paid mutator transaction binding the contract method 0xccb98ffc.
//
// Solidity: function setEndTime(uint256 _time) returns(bool)
func (_CrowdSale *CrowdSaleTransactor) SetEndTime(opts *bind.TransactOpts, _time *big.Int) (*types.Transaction, error) {
	return _CrowdSale.contract.Transact(opts, "setEndTime", _time)
}

// SetEndTime is a paid mutator transaction binding the contract method 0xccb98ffc.
//
// Solidity: function setEndTime(uint256 _time) returns(bool)
func (_CrowdSale *CrowdSaleSession) SetEndTime(_time *big.Int) (*types.Transaction, error) {
	return _CrowdSale.Contract.SetEndTime(&_CrowdSale.TransactOpts, _time)
}

// SetEndTime is a paid mutator transaction binding the contract method 0xccb98ffc.
//
// Solidity: function setEndTime(uint256 _time) returns(bool)
func (_CrowdSale *CrowdSaleTransactorSession) SetEndTime(_time *big.Int) (*types.Transaction, error) {
	return _CrowdSale.Contract.SetEndTime(&_CrowdSale.TransactOpts, _time)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns(bool)
func (_CrowdSale *CrowdSaleTransactor) SetPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _CrowdSale.contract.Transact(opts, "setPrice", _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns(bool)
func (_CrowdSale *CrowdSaleSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _CrowdSale.Contract.SetPrice(&_CrowdSale.TransactOpts, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns(bool)
func (_CrowdSale *CrowdSaleTransactorSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _CrowdSale.Contract.SetPrice(&_CrowdSale.TransactOpts, _price)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _time) returns(bool)
func (_CrowdSale *CrowdSaleTransactor) SetStartTime(opts *bind.TransactOpts, _time *big.Int) (*types.Transaction, error) {
	return _CrowdSale.contract.Transact(opts, "setStartTime", _time)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _time) returns(bool)
func (_CrowdSale *CrowdSaleSession) SetStartTime(_time *big.Int) (*types.Transaction, error) {
	return _CrowdSale.Contract.SetStartTime(&_CrowdSale.TransactOpts, _time)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _time) returns(bool)
func (_CrowdSale *CrowdSaleTransactorSession) SetStartTime(_time *big.Int) (*types.Transaction, error) {
	return _CrowdSale.Contract.SetStartTime(&_CrowdSale.TransactOpts, _time)
}

// SetTargetFunds is a paid mutator transaction binding the contract method 0xf756017a.
//
// Solidity: function setTargetFunds(uint256 _amount) returns(bool)
func (_CrowdSale *CrowdSaleTransactor) SetTargetFunds(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _CrowdSale.contract.Transact(opts, "setTargetFunds", _amount)
}

// SetTargetFunds is a paid mutator transaction binding the contract method 0xf756017a.
//
// Solidity: function setTargetFunds(uint256 _amount) returns(bool)
func (_CrowdSale *CrowdSaleSession) SetTargetFunds(_amount *big.Int) (*types.Transaction, error) {
	return _CrowdSale.Contract.SetTargetFunds(&_CrowdSale.TransactOpts, _amount)
}

// SetTargetFunds is a paid mutator transaction binding the contract method 0xf756017a.
//
// Solidity: function setTargetFunds(uint256 _amount) returns(bool)
func (_CrowdSale *CrowdSaleTransactorSession) SetTargetFunds(_amount *big.Int) (*types.Transaction, error) {
	return _CrowdSale.Contract.SetTargetFunds(&_CrowdSale.TransactOpts, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address addr) returns()
func (_CrowdSale *CrowdSaleTransactor) TransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CrowdSale.contract.Transact(opts, "transferOwnership", addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address addr) returns()
func (_CrowdSale *CrowdSaleSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _CrowdSale.Contract.TransferOwnership(&_CrowdSale.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address addr) returns()
func (_CrowdSale *CrowdSaleTransactorSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _CrowdSale.Contract.TransferOwnership(&_CrowdSale.TransactOpts, addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_CrowdSale *CrowdSaleTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdSale.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_CrowdSale *CrowdSaleSession) Withdraw() (*types.Transaction, error) {
	return _CrowdSale.Contract.Withdraw(&_CrowdSale.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_CrowdSale *CrowdSaleTransactorSession) Withdraw() (*types.Transaction, error) {
	return _CrowdSale.Contract.Withdraw(&_CrowdSale.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_CrowdSale *CrowdSaleTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _CrowdSale.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_CrowdSale *CrowdSaleSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CrowdSale.Contract.Fallback(&_CrowdSale.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_CrowdSale *CrowdSaleTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CrowdSale.Contract.Fallback(&_CrowdSale.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrowdSale *CrowdSaleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdSale.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrowdSale *CrowdSaleSession) Receive() (*types.Transaction, error) {
	return _CrowdSale.Contract.Receive(&_CrowdSale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrowdSale *CrowdSaleTransactorSession) Receive() (*types.Transaction, error) {
	return _CrowdSale.Contract.Receive(&_CrowdSale.TransactOpts)
}

// CrowdSaleBoughtIterator is returned from FilterBought and is used to iterate over the raw logs and unpacked data for Bought events raised by the CrowdSale contract.
type CrowdSaleBoughtIterator struct {
	Event *CrowdSaleBought // Event containing the contract specifics and raw log

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
func (it *CrowdSaleBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdSaleBought)
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
		it.Event = new(CrowdSaleBought)
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
func (it *CrowdSaleBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdSaleBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdSaleBought represents a Bought event raised by the CrowdSale contract.
type CrowdSaleBought struct {
	Owner  common.Address
	Amount *big.Int
	Raise  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBought is a free log retrieval operation binding the contract event 0xa9a40dec7a304e5915d11358b968c1e8d365992abf20f82285d1df1b30c8e24c.
//
// Solidity: event Bought(address indexed owner, uint256 amount, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) FilterBought(opts *bind.FilterOpts, owner []common.Address) (*CrowdSaleBoughtIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CrowdSale.contract.FilterLogs(opts, "Bought", ownerRule)
	if err != nil {
		return nil, err
	}
	return &CrowdSaleBoughtIterator{contract: _CrowdSale.contract, event: "Bought", logs: logs, sub: sub}, nil
}

// WatchBought is a free log subscription operation binding the contract event 0xa9a40dec7a304e5915d11358b968c1e8d365992abf20f82285d1df1b30c8e24c.
//
// Solidity: event Bought(address indexed owner, uint256 amount, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) WatchBought(opts *bind.WatchOpts, sink chan<- *CrowdSaleBought, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CrowdSale.contract.WatchLogs(opts, "Bought", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdSaleBought)
				if err := _CrowdSale.contract.UnpackLog(event, "Bought", log); err != nil {
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

// ParseBought is a log parse operation binding the contract event 0xa9a40dec7a304e5915d11358b968c1e8d365992abf20f82285d1df1b30c8e24c.
//
// Solidity: event Bought(address indexed owner, uint256 amount, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) ParseBought(log types.Log) (*CrowdSaleBought, error) {
	event := new(CrowdSaleBought)
	if err := _CrowdSale.contract.UnpackLog(event, "Bought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdSaleRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the CrowdSale contract.
type CrowdSaleRefundedIterator struct {
	Event *CrowdSaleRefunded // Event containing the contract specifics and raw log

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
func (it *CrowdSaleRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdSaleRefunded)
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
		it.Event = new(CrowdSaleRefunded)
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
func (it *CrowdSaleRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdSaleRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdSaleRefunded represents a Refunded event raised by the CrowdSale contract.
type CrowdSaleRefunded struct {
	From   common.Address
	Amount *big.Int
	Raise  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0x2dc8e290002f06fc0085bbca9dfb8b415cf4d1178950c72ff9ee8f4d8878ee66.
//
// Solidity: event Refunded(address indexed from, uint256 amount, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) FilterRefunded(opts *bind.FilterOpts, from []common.Address) (*CrowdSaleRefundedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _CrowdSale.contract.FilterLogs(opts, "Refunded", fromRule)
	if err != nil {
		return nil, err
	}
	return &CrowdSaleRefundedIterator{contract: _CrowdSale.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0x2dc8e290002f06fc0085bbca9dfb8b415cf4d1178950c72ff9ee8f4d8878ee66.
//
// Solidity: event Refunded(address indexed from, uint256 amount, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *CrowdSaleRefunded, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _CrowdSale.contract.WatchLogs(opts, "Refunded", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdSaleRefunded)
				if err := _CrowdSale.contract.UnpackLog(event, "Refunded", log); err != nil {
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

// ParseRefunded is a log parse operation binding the contract event 0x2dc8e290002f06fc0085bbca9dfb8b415cf4d1178950c72ff9ee8f4d8878ee66.
//
// Solidity: event Refunded(address indexed from, uint256 amount, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) ParseRefunded(log types.Log) (*CrowdSaleRefunded, error) {
	event := new(CrowdSaleRefunded)
	if err := _CrowdSale.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdSaleWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the CrowdSale contract.
type CrowdSaleWithdrawnIterator struct {
	Event *CrowdSaleWithdrawn // Event containing the contract specifics and raw log

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
func (it *CrowdSaleWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdSaleWithdrawn)
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
		it.Event = new(CrowdSaleWithdrawn)
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
func (it *CrowdSaleWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdSaleWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdSaleWithdrawn represents a Withdrawn event raised by the CrowdSale contract.
type CrowdSaleWithdrawn struct {
	From  common.Address
	Raise *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed from, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) FilterWithdrawn(opts *bind.FilterOpts, from []common.Address) (*CrowdSaleWithdrawnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _CrowdSale.contract.FilterLogs(opts, "Withdrawn", fromRule)
	if err != nil {
		return nil, err
	}
	return &CrowdSaleWithdrawnIterator{contract: _CrowdSale.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed from, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *CrowdSaleWithdrawn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _CrowdSale.contract.WatchLogs(opts, "Withdrawn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdSaleWithdrawn)
				if err := _CrowdSale.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed from, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) ParseWithdrawn(log types.Log) (*CrowdSaleWithdrawn, error) {
	event := new(CrowdSaleWithdrawn)
	if err := _CrowdSale.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
