package tokenlauncher

import (
	"math/big"
	"testing"
	"tokenlauncher/uint256"
)

func TestNewLauncher(t *testing.T) {
	rcpUrl := "http://localhost:8101"
	keyFile := "D:\\workspace\\fixtoken\\data0\\keystore\\UTC--2021-03-05T09-35-51.816544100Z--60d1148b3b2ab38a5937dc30244a3b4c5ec6da52"
	password := "123"
	launcher, err := NewLauncher(rcpUrl,big.NewInt(10086),keyFile,password)
	if err != nil {
		t.Error(err)
	}
	_ = launcher
}

func TestLauncher_DeployContract(t *testing.T) {
	rcpUrl := "http://localhost:8101"
	keyFile := "D:\\workspace\\fixtoken\\data0\\keystore\\UTC--2021-03-05T09-35-51.816544100Z--60d1148b3b2ab38a5937dc30244a3b4c5ec6da52"
	password := "123"
	launcher, err := NewLauncher(rcpUrl,big.NewInt(10086),keyFile,password)
	if err != nil {
		t.Error(err)
	}
	totalSupply := *uint256.NewUInt256ByUInt32(uint32(1))
	token := &ERC20Token{
		Name: "UniToken",
		Symbol: "UNI",
		InitialSupply: totalSupply,
		Decimals: uint8(12),
	}
	c, err := NewERC20TokenContract(token)
	if err != nil {
		t.Error(err)
	}
	addr, err := launcher.DeployContract(c)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Contract address: %s", addr.Hex())
	t.Logf("tk = eth.contract(abi).at('%s')", addr.Hex())
}

func TestImportKeyStoreByFilepath(t *testing.T) {
	keyFile := "D:\\workspace\\fixtoken\\data0\\keystore\\UTC--2021-03-05T09-35-51.816544100Z--60d1148b3b2ab38a5937dc30244a3b4c5ec6da52"
	password := "123"
	key, err := ImportKeyStoreByFilepath(keyFile, password)
	_ = key
	_ = err
}