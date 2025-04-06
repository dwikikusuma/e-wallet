package cmd

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
	"wallet/helpers"
)

func ServeGRPC() {
	list, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "50051"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Register your gRPC service here
	logrus.Info("gRPC server started on port " + helpers.GetEnv("GRPC_PORT", "50051"))
	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
