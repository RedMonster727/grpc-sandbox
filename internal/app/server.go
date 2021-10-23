package app

import (
	"context"

	server2 "github.com/RedMonster727/grpc-sandbox/internal/generated/server/proto"
	"github.com/RedMonster727/grpc-sandbox/internal/grpc/get_some_data"
)

type server struct {
	server2.UnimplementedTestServerServer
	getSomeDataHandler func(ctx context.Context, server3 *server2.GetSomeDataIn) (*server2.GetSomeDataOut, error)
}

// NewServer returns server instance and err
// in case if one of dependencies requires an error to process
func NewServer() (server2.TestServerServer, error) {
	getSomeDataHandler := get_some_data.NewHandler().GetSomeData

	s := &server{
		getSomeDataHandler: getSomeDataHandler,
	}
	return s, nil
}
