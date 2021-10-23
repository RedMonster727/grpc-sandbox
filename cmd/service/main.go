package main

import (
	"fmt"
	"log"
	"net"

	"github.com/RedMonster727/grpc-sandbox/internal/app"
	server2 "github.com/RedMonster727/grpc-sandbox/internal/generated/server/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8891")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	grpcServer, err := app.NewServer()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to init application: %w", err))
	}

	server2.RegisterTestServerServer(s, grpcServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
