package tokenlauncher

import "github.com/ethereum/go-ethereum/core/types"

type Contract struct {
	Data string
	ABI string
	Address string
}


func NewERC20TokenContract(token *ERC20Token) *Contract {
	abiFile := "./contract/token.abi"
	//tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	//td := new types.TxData()
	//types.NewTx()
	return &Contract{
		ABI: abiFile,
	}
}

func (c *Contract) BuildTransaction() *types.Transaction {

	return types.NewTx(&types.LegacyTx{

	})
}