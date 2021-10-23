package app

import (
	"context"

	server2 "github.com/RedMonster727/grpc-sandbox/internal/generated/server/proto"
)

func (s *server) GetSomeData(ctx context.Context, in *server2.GetSomeDataIn) (*server2.GetSomeDataOut, error) {
	return s.getSomeDataHandler(ctx, in)
}
