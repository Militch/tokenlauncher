package tokenlauncher

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
	"tokenlauncher/api"
	"tokenlauncher/launcher"
)

func TestNewRPCServer(t *testing.T) {
	starter,err := NewRPCServerStarter()
	if err != nil {
		t.Fatal(err)
	}
	err = starter.RegisterName("ERC20Token",
		new(api.ERC20TokenAPIHandler))
	if err != nil {
		t.Fatal(err)
	}
	err = starter.RegisterName("CrowdSale",
		new(api.CrowdSaleAPIHandler))
	if err != nil {
		t.Fatal(err)
	}

	err = starter.Run(":9005")
	if err != nil {
		t.Fatal(err)
	}
}

func TestName(t *testing.T) {

	privateKey,err := crypto.HexToECDSA("b4d4c0ec95f939f0311eaace6614e9170b4b9b33e5fe57b9b8bdfe743c466a6b")
	if err != nil {
		t.Error(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Error("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	t.Logf("address: %s", address)
}

func TestIm(t *testing.T) {
	keyFilePath := "./keystore/UTC--2021-03-12T12-52-15.284179379Z--ffd15b670f81779fd65065784afb32fa41822292"
	password := "123"
	key, err := launcher.ImportKeyStoreByFilepath(keyFilePath, password)
	if err != nil{
		t.Error("cannot import key store from file name")
	}
	if key != nil {
		address := key.Address
		privateKeyBytes := crypto.FromECDSA(key.PrivateKey)
		t.Logf("address: %s", address.Hex())
		t.Logf("privateKey: %s", hexutil.Encode(privateKeyBytes)[2:])
	}


}