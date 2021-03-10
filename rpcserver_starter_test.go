package tokenlauncher

import (
	"testing"
	"tokenlauncher/api"
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
	err = starter.RegisterName("Contract",
		new(api.ContractAPIHandler))
	if err != nil {
		t.Fatal(err)
	}
	err = starter.RegisterName("CrowdSale",
		new(api.ContractAPIHandler))
	if err != nil {
		t.Fatal(err)
	}

	err = starter.Run(":9005")
	if err != nil {
		t.Fatal(err)
	}
}