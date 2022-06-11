package server

import (
	"net"
	"net/rpc"
)

func main() {
	// start the rpc server
	l, err := net.Listen("tcp", ":8081")
	logger.Errorf("Error while starting rpc server: %+v", err)
	go func() {
		for {
			rpc.Accept(l)
		}
	}()
	logger.Info("Listening on port 8081 for RPC requests")
}
