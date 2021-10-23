package get_some_data

import (
	"context"
	server2 "github.com/RedMonster727/grpc-sandbox/internal/generated/server/proto"
	"strconv"
)

type handler struct {
	someData int64
}

func NewHandler() *handler {
	return &handler{
		someData: 5,
	}
}

func (h *handler) GetSomeData(_ context.Context, in *server2.GetSomeDataIn) (*server2.GetSomeDataOut, error) {
	result := make(map[string]*server2.OutNestedMessage)
	for _, el := range in.GetIn().GetValue() {
		result[strconv.FormatInt(el, 10)] = &server2.OutNestedMessage{Value: []int64{h.someData}}
	}
	return &server2.GetSomeDataOut{Result: result}, nil
}