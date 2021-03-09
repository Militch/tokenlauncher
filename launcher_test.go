package tokenlauncher

import (
	"fmt"
	"math/big"
	"testing"
	"time"
	"tokenlauncher/model"
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
	token := &model.ERC20TokenOpts{
		Name: "UniToken",
		Symbol: "UNI",
		InitialSupply: "0x1",
		Decimals: uint8(12),
	}
	cAddr, err := launcher.DeployERC20Token(token)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Contract address: %s", cAddr.Hex())
	targetFunds := *uint256.NewUInt256ByUInt32(100)
	price := *uint256.NewUInt256("0xde0b6b3a7640000")
	currentTime := time.Now().Unix()
	startTime := *uint256.NewUInt256ByUInt64(uint64(currentTime))
	endTimeAt := currentTime + (1 * 24 * 60 * 60)
	endTime := *uint256.NewUInt256ByUInt64(uint64(endTimeAt))
	crowdSaleOpts := &model.CrowdSaleOpts{
		TokenAddress: *cAddr,
		TargetFunds: targetFunds,
		Price: price,
		StartTime: startTime,
		EndTime: endTime,
	}
	cAddr, err = launcher.DeployCrowdSale(crowdSaleOpts)
	if err != nil {
		t.Error(err)
	}
	if err != nil {
		t.Error(err)
	}
	t.Logf("Contract address: %s", cAddr.Hex())
}

func TestImportKeyStoreByFilepath(t *testing.T) {
	keyFile := "D:\\workspace\\fixtoken\\data0\\keystore\\UTC--2021-03-05T09-35-51.816544100Z--60d1148b3b2ab38a5937dc30244a3b4c5ec6da52"
	password := "123"
	key, err := ImportKeyStoreByFilepath(keyFile, password)
	_ = key
	_ = err
}

func TestName(t *testing.T) {
	currentTime := time.Now().Unix()
	startTime := *uint256.NewUInt256ByUInt64(uint64(currentTime))
	endTimeAt := currentTime + (1 * 24 * 60 * 60)
	endTime := *uint256.NewUInt256ByUInt64(uint64(endTimeAt))
	fmt.Printf("currentTime: %d\n", currentTime)
	fmt.Printf("endTimeAt: %d\n", endTimeAt)
	fmt.Printf("startTime: %d\n", startTime.ToBigInt())
	fmt.Printf("endTime: %d\n", endTime.ToBigInt())
}