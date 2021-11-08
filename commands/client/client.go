package main

import (
	"context"
	"go-grpc/internal/rpc"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:6565"
)

func main() {
	conn , err := grpc.Dial(address , grpc.WithInsecure() , grpc.WithBlock())
	if err != nil{
		log.Fatalf("did not connect: %v" , err)
	}
	defer conn.Close()

	c := rpc.NewSampleServiceClient(conn)
	ctx  , cancel := context.WithTimeout(context.Background() , time.Second)
	defer cancel()

	result , err := c.CreateSample(ctx , &rpc.NewSample{
		Name: "hello world",
	})
	if err != nil {
		log.Fatalf("could not create sample: %v" , err)
	}
	log.Printf("user detail: %v" , result)
}
