package sample

import (
	"context"
	"go-grpc/internal/rpc"
	"log"
)

type SampleServiceServer struct {
	rpc.UnimplementedSampleServiceServer
}

func (s *SampleServiceServer) CreateSample(ctx context.Context , in *rpc.NewSample) (*rpc.Sample , error)  {
	log.Printf("Received: %v", in.GetName())

	return &rpc.Sample{
		Id: int32(123),
		Name: "hey",
	} , nil
}
