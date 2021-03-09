package main

import (
	"log"
	"tokenlauncher"
)

func main() {
	s,err := tokenlauncher.NewRPCServerStarter()
	if err != nil {
		log.Fatal(err)
	}
	err = s.Bootstrap(":9058")
	if err != nil {
		log.Fatal(err)
	}

	//client.
}