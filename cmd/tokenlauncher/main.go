package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8101")
	if err != nil {
		log.Fatal(err)
	}
	_ = client
	fmt.Printf("abc \n")

	//client.
}