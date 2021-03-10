package launcher

import (
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math/big"
	"tokenlauncher/contract"
	"tokenlauncher/uint256"
)

type Launcher struct {
	CLI *ethclient.Client
	key *keystore.Key
	chainId *big.Int

}

type config struct {
	RpcUrl      string `json:"rpcUrl"`
	ChainId int64 `json:"chainID"`
	KeyFilePath string `json:"keyFilePath"`
	KeyPassword string `json:"keyPassword"`
}

func ImportKeyStoreByFilepath(keyFilepath string, password string) (key *keystore.Key, err error) {
	jsonBytes, err := ioutil.ReadFile(keyFilepath)
	if err != nil {
		return nil, err
	}
	return keystore.DecryptKey(jsonBytes, password)
}

func NewLauncherDefault() (*Launcher, error)  {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return nil, err
	}
	var c *config
	err = json.Unmarshal(data, &c)
	if err != nil{
		return nil, err
	}
	return NewLauncher(c.RpcUrl, big.NewInt(c.ChainId), c.KeyFilePath, c.KeyPassword)
}

// 创建启动器
func NewLauncher(rpcUrl string, chainId *big.Int, keyFilepath string, password string) (*Launcher, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	key, err := ImportKeyStoreByFilepath(keyFilepath, password)
	if err != nil {
		return nil, err
	}

	return &Launcher{
		CLI: client,
		key: key,
		chainId: chainId,
	}, nil
}

func (launcher *Launcher) DefaultTxOpts()  (*bind.TransactOpts, error)  {
	privateKey := launcher.key.PrivateKey
	return bind.NewKeyedTransactorWithChainID(
		privateKey,  launcher.chainId)
}

// 部署合约
func (launcher *Launcher) DeployERC20Token(opts *ERC20TokenOpts) (*Contract, error) {
	address := launcher.key.Address
	txOpts, err := launcher.DefaultTxOpts()
	if err != nil {
		return nil, err
	}
	initialSupply := uint256.NewUInt256(opts.InitialSupply).ToBigInt()
	address, tx, _, err := contract.DeployERC20Token(
		txOpts, launcher.CLI,opts.Name, opts.Symbol,
		initialSupply, opts.Decimals)
	if err != nil {
		return nil, err
	}
	txdata := "0x" + hex.EncodeToString(tx.Data())
	txhash := tx.Hash().Hex()
	con := &Contract{
		Address: address.Hex(),
		TxHash: txhash,
		ABI: contract.ERC20TokenABI,
		Data: txdata,
	}
	return con,nil
}

func (launcher *Launcher) LoadERC20TokenByContractAddr(addr string) (*contract.ERC20Token, error) {
	contractAddr := common.HexToAddress(addr)
	return contract.NewERC20Token(contractAddr,launcher.CLI)
}
func (launcher *Launcher) LoadCrowdSaleByContractAddr(addr string) (*contract.CrowdSale, error) {
	contractAddr := common.HexToAddress(addr)
	return contract.NewCrowdSale(contractAddr,launcher.CLI)
}
func (launcher *Launcher) DeployCrowdSale(opts *CrowdSaleOpts)  (*Contract, error){
	address := launcher.key.Address
	privateKey := launcher.key.PrivateKey
	txOpts, err := bind.NewKeyedTransactorWithChainID(
		privateKey, launcher.chainId)
	if err != nil {
		return nil, err
	}
	address, tx, _, err := contract.DeployCrowdSale(
		txOpts, launcher.CLI,opts.TokenAddress, opts.TargetFunds.ToBigInt(),
		opts.Price.ToBigInt(), opts.StartTime.ToBigInt(), opts.EndTime.ToBigInt())
	if err != nil {
		return nil, err
	}
	txdata := "0x" + hex.EncodeToString(tx.Data())
	txhash := tx.Hash().Hex()
	// 合约地址
	con := &Contract{
		Address: address.Hex(),
		TxHash: txhash,
		ABI: contract.ERC20TokenABI,
		Data: txdata,
	}
	return con,nil
}