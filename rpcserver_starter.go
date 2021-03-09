package tokenlauncher

import (
	"log"
	"net"
	"net/rpc"
	"tokenlauncher/api"
)

type RPCServerStarter struct {
	rpcserver *rpc.Server
}

func NewRPCServerStarter() (*RPCServerStarter, error) {
	rpcserver := rpc.NewServer()
	err := rpcserver.Register(new(api.ERC20TokenAPIHandler))
	if err != nil { return nil, err }
	return &RPCServerStarter{
		rpcserver: rpcserver,
	}, nil
}

func (server *RPCServerStarter) Bootstrap(addr string) error {

	rpcserver := server.rpcserver
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("err: %v",err)
			continue
		}
		go rpcserver.ServeConn(conn)
	}
}