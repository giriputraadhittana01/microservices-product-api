package main

import (
	"currentcyproject/server"
	"net"
	"os"

	protos "currentcyproject/protos/currency"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	rpc := grpc.NewServer()
	cs := server.NewCurrency(log)

	protos.RegisterCurrencyServer(rpc, cs)

	reflection.Register(rpc)

	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Error("Unable to listen", "Error", err)
		os.Exit(1)
	}
	rpc.Serve(l)
}
