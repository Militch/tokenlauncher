package tokenlauncher

import "testing"

func TestNewRPCServer(t *testing.T) {
	starter,err := NewRPCServerStarter()
	if err != nil {
		t.Fatal(err)
	}
	err = starter.Bootstrap("127.0.0.1:9005")
	if err != nil {
		t.Fatal(err)
	}
}