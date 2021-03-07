package tokenlauncher

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math/big"
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

func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.TODO()
	}
	return ctx
}

func (launcher *Launcher) SuggestGasPrice() (*big.Int, error)  {
	contractTransactor := bind.ContractTransactor(launcher.cli)
	return contractTransactor.SuggestGasPrice(context.TODO())
}

func (launcher *Launcher) PendingNonce() (uint64, error)  {
	address := launcher.key.Address
	contractTransactor := bind.ContractTransactor(launcher.cli)
	return contractTransactor.PendingNonceAt(
		context.TODO(), address)
}

func (launcher *Launcher) EstimateGas(
	opts *bind.TransactOpts,
	contract *common.Address,
	input []byte) (uint64, error)  {
	address := launcher.key.Address
	contractTransactor := bind.ContractTransactor(launcher.cli)

	if contract != nil {
		code, err := contractTransactor.PendingCodeAt(
			ensureContext(opts.Context),address)
		if err != nil {
			return 0, err
		} else if len(code) == 0 {
			return 0, errors.New("no contract code at given address")
		}
	}
	msg := ethereum.CallMsg{
		From: opts.From,
		GasPrice: opts.GasPrice,
		Value: opts.Value,
		Data: input,
	}
	return contractTransactor.EstimateGas(ensureContext(opts.Context), msg)
}


// 部署合约
func (launcher *Launcher) DeployContract(contract *Contract) (*common.Address, error) {
	// common.Address
	contractTransactor := bind.ContractTransactor(launcher.cli)
	address := launcher.key.Address
	privateKey := launcher.key.PrivateKey
	txOpts, err := bind.NewKeyedTransactorWithChainID(
		privateKey, launcher.chainID)
	if err != nil {
		return nil, err
	}
	// 确认随机数
	nonce, err := launcher.PendingNonce()
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}
	txOpts.Nonce = new(big.Int).SetUint64(nonce)
	// 确认gas单价
	gasPrice, err := launcher.SuggestGasPrice()
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas needed: %v", err)
	}
	txOpts.GasPrice = gasPrice
	// 确认gas Limit
	gasLimit, err := launcher.EstimateGas(
		txOpts,nil, contract.Data)
	value := txOpts.Value
	if value == nil {
		value = new(big.Int)
	}
	rawTx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     contract.Data,
	})
	if txOpts.Signer == nil {
		return nil, errors.New("no signer to authorize the transaction with")
	}
	signedTx, err := txOpts.Signer(address, rawTx)
	if err != nil {
		return nil, err
	}
	err = contractTransactor.SendTransaction(
		ensureContext(txOpts.Context), signedTx)
	if err != nil {
		return nil, err
	}
	contractAddr := crypto.CreateAddress(txOpts.From, signedTx.Nonce())
	// 合约地址
	return &contractAddr,nil
}
// 加载合约
func (launcher *Launcher) LoadContract(address string) *Contract {

	return &Contract{

	}
}