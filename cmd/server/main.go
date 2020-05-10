package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/billcunha/bridgekeeper/pkg/api"
	"github.com/billcunha/bridgekeeper/pkg/keeper"
)

func main() {
	// Create new gRPC server instance
	s := grpc.NewServer()
	srv := &keeper.GRPCServer{}

	// Register gRPC server
	api.RegisterKeeperServer(s, srv)

	// Listen on port 8080
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// Start gRPC server
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
