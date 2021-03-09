package tokenlauncher

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math/big"
	"tokenlauncher/contract"
	"tokenlauncher/model"
	"tokenlauncher/uint256"
)

type Launcher struct {
	cli *ethclient.Client
	key *keystore.Key
	chainID *big.Int
}

func ImportKeyStoreByFilepath(keyFilepath string, password string) (key *keystore.Key, err error) {
	jsonBytes, err := ioutil.ReadFile(keyFilepath)
	if err != nil {
		return nil, err
	}
	return keystore.DecryptKey(jsonBytes, password)
}

// 创建启动器
func NewLauncher(rpcUrl string, chainID *big.Int, keyFilepath string, password string) (*Launcher, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	key, err := ImportKeyStoreByFilepath(keyFilepath, password)
	if err != nil {
		return nil, err
	}
	return &Launcher{
		cli: client,
		key: key,
		chainID: chainID,
	}, nil
}


// 部署合约
func (launcher *Launcher) DeployERC20Token(opts *model.ERC20TokenOpts) (*common.Address, error) {
	address := launcher.key.Address
	privateKey := launcher.key.PrivateKey
	txOpts, err := bind.NewKeyedTransactorWithChainID(
		privateKey, launcher.chainID)
	if err != nil {
		return nil, err
	}
	initialSupply := uint256.NewUInt256(opts.InitialSupply).ToBigInt()
	address, _, _, err = contract.DeployERC20Token(
		txOpts, launcher.cli,opts.Name, opts.Symbol,
		initialSupply, opts.Decimals)
	if err != nil {
		return nil, err
	}
	// 合约地址
	return &address,nil
}

func (launcher *Launcher) DeployCrowdSale(opts *model.CrowdSaleOpts) (*common.Address, error) {
	address := launcher.key.Address
	privateKey := launcher.key.PrivateKey
	txOpts, err := bind.NewKeyedTransactorWithChainID(
		privateKey, launcher.chainID)
	if err != nil {
		return nil, err
	}
	address, _, _, err = contract.DeployCrowdSale(
		txOpts, launcher.cli,opts.TokenAddress, opts.TargetFunds.ToBigInt(),
		opts.Price.ToBigInt(), opts.StartTime.ToBigInt(), opts.EndTime.ToBigInt())
	if err != nil {
		return nil, err
	}
	// 合约地址
	return &address,nil
}