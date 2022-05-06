package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-server-client-alpha/protocol"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedDataInterfaceServer
}

func (s *Server) GetChunk(ctx context.Context, in *pb.ChunkRequest) (*pb.ChunkResponse, error) {
	log.Printf("Received %v", in.GetId())

	return &pb.ChunkResponse{Id: "12345", Data: []byte{10, 20, 30, 40}, Length: 4}, nil
}

func main() {
	port := 50051

	var lis net.Listener
	var err error

	if lis, err = net.Listen("tcp", fmt.Sprintf("localhost:%d", port)); err != nil {
		log.Fatalf("Error listening on port: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDataInterfaceServer(s, &Server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
