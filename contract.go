package tokenlauncher

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"io/ioutil"
	"os"
	"strings"
)

type Contract struct {
	ByteCode []byte
	ABI abi.ABI
	Address string
	Data []byte
}

func readStringByFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
func NewERC20TokenContract(token *ERC20Token) (*Contract, error) {
	abiFilename := "./build/ERC20Token.abi"
	binFilename := "./build/ERC20Token.bin"
	abiString, err := readStringByFile(abiFilename)
	if err != nil {
		return nil, err
	}
	binString, err := readStringByFile(binFilename)
	if err != nil {
		return nil, err
	}
	parsed, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return nil, err
	}
	byteCode := common.FromHex(binString)
	initialSupply := token.InitialSupply.ToBigInt()
	input, err := parsed.Pack("", token.Name,
		token.Symbol, initialSupply, token.Decimals)
	if err != nil {
		return nil, err
	}
	return &Contract{
		ABI: parsed,
		ByteCode: byteCode,
		Data: append(byteCode, input...),
	}, nil
}



func (c *Contract) BuildTransaction() *types.Transaction {


	return types.NewTx(&types.LegacyTx{

	})
}

func (c *Contract) BuildTxInput() *types.Transaction {


	return types.NewTx(&types.LegacyTx{

	})
}