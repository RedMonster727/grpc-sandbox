package main

import (
	"context"
	server2 "github.com/RedMonster727/grpc-sandbox/internal/generated/server/proto"
	"github.com/RedMonster727/grpc-sandbox/internal/grpc/get_some_data"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	server2.UnimplementedTestServerServer
	getSomeDataHandler func(ctx context.Context, server3 *server2.GetSomeDataIn) (*server2.GetSomeDataOut, error)
}

func (s *server) GetSomeData(ctx context.Context, in *server2.GetSomeDataIn) (*server2.GetSomeDataOut, error) {
	return s.getSomeDataHandler(ctx, in)
}

func main() {
	lis, err := net.Listen("tcp", ":8891")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	handler := get_some_data.NewHandler().GetSomeData

	s := grpc.NewServer()
	server2.RegisterTestServerServer(s, &server{
		getSomeDataHandler: handler,
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


