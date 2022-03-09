package main

import (
	"currentcyproject/server"
	"net"
	"os"

	protos "currentcyproject/protos/currency"

	"currentcyproject/data"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()
	rates, err := data.NewRates(log)
	if err != nil {
		log.Error("Unable to generate rates", "error", err)
		os.Exit(1)
	}
	rpc := grpc.NewServer()
	cs := server.NewCurrency(rates, log)

	protos.RegisterCurrencyServer(rpc, cs)

	reflection.Register(rpc)

	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Error("Unable to listen", "Error", err)
		os.Exit(1)
	}
	rpc.Serve(l)
}
