package tokenlauncher

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
)

type Launcher struct {
	cli *ethclient.Client
	account accounts.Account
}

func ImportKeyStoreByFilepath(keyFilepath string, password string) accounts.Account {
	jsonBytes, err := ioutil.ReadFile(keyFilepath)
	if err != nil {
		log.Fatal(err)
	}
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}
	return account
}

// 创建启动器
func NewLauncher(rpcUrl string, keyFilepath string, password string) *Launcher {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}
	account := ImportKeyStoreByFilepath(keyFilepath, password)
	return &Launcher{
		cli: client,
		account: account,
	}
}

// 部署合约
func (launcher *Launcher) DeployContract(contract *Contract) {
	tx := contract.BuildTransaction()
	err := launcher.cli.SendTransaction(context.Background(),tx)
	if err != nil {

	}
}
// 加载合约
func (launcher *Launcher) LoadContract(address string) *Contract {

	return &Contract{

	}
}