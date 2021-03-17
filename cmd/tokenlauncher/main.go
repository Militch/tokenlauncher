package main

import (
	"flag"
	"fmt"
	"os"
	"tokenlauncher"
	"tokenlauncher/api"
)

var (
	h bool
	v bool
	port string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.StringVar(&port, "p", "9017", "set rpc listen port `port`")
	flag.Usage = usage
}


func main() {

	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(0)
	}

	if v {
		fmt.Printf("tokenlauncher version: 0.1.0")
		os.Exit(0)
	}
	starter,err := tokenlauncher.NewRPCServerStarter()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	err = starter.RegisterName("ERC20Token",
		new(api.ERC20TokenAPIHandler))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	err = starter.RegisterName("CrowdSale",
		new(api.CrowdSaleAPIHandler))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = starter.Run(":" + port)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}


func usage() {
	fmt.Printf(`tokenlauncher version: 0.1.0
Usage: tokenlauncher [-hv] [-c filename]

Options:
`)
	flag.PrintDefaults()
}