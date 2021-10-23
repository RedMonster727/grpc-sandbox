package main

import (
	"context"
	"log"
	"time"

	server2 "github.com/RedMonster727/grpc-sandbox/internal/generated/server/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8891"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := server2.NewTestServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	i := 0
	for {
		if i > 100 {
			break
		}
		r, err := c.GetSomeData(ctx, &server2.GetSomeDataIn{In: &server2.SomeNestedMessage{Value: []int64{1, 2, 3}}})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetResult())
		i++
	}

}
