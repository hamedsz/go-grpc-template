package main

import (
	internalRpc "go-grpc/internal/rpc"
	"go-grpc/internal/services/sample"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	var port = ":6565"
	lis , err := net.Listen("tcp" , port)
	if err != nil{
		log.Fatalf("failed to listen: %v" , err)
	}

	s := grpc.NewServer()
	internalRpc.RegisterSampleServiceServer(s , &sample.SampleServiceServer{})

	log.Printf("server listening at %v" , lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc: %v" , err)
	}
}
