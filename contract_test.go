package tokenlauncher

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"testing"
)

func TestCreateERC20TokenContract(t *testing.T) {
	//totalSupply := *uint256.NewUInt256ByUInt32(uint32(2000000000))
	//token := &ERC20Token{
	//	Name: "FixToken",
	//	Symbol: "FIX",
	//	TotalSupply: totalSupply,
	//	Decimals: uint8(12),
	//}
	client, err := ethclient.Dial("http://localhost:8101")
	if err != nil {
		log.Fatal(err)
	}
	_ = client
	//contract.NewERC20TokenContract(token)
}

func TestNewERC20TokenContract(t *testing.T) {

}