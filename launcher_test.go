package tokenlauncher

import (
	"testing"
	"tokenlauncher/uint256"
)

func TestNewLauncher(t *testing.T) {
	rcpUrl := "http://localhost:8101"
	keyFile := "D:\\workspace\\fixtoken\\data0\\keystore\\UTC--2021-03-05T09-35-51.816544100Z--60d1148b3b2ab38a5937dc30244a3b4c5ec6da52"
	password := "123"
	launcher := NewLauncher(rcpUrl,keyFile,password)
	_ = launcher
}

func TestLauncher_DeployContract(t *testing.T) {
	rcpUrl := "http://localhost:8101"
	keyFile := "D:\\workspace\\fixtoken\\data0\\keystore\\UTC--2021-03-05T09-35-51.816544100Z--60d1148b3b2ab38a5937dc30244a3b4c5ec6da52"
	password := "123"

	launcher := NewLauncher(rcpUrl,keyFile,password)

	totalSupply := *uint256.NewUInt256ByUInt32(uint32(2000000000))


	token := &ERC20Token{
		Name: "FixToken",
		Symbol: "FIX",
		TotalSupply: totalSupply,
		Decimals: uint8(12),
	}

	contract := NewERC20TokenContract(token)
	launcher.DeployContract(contract)

}