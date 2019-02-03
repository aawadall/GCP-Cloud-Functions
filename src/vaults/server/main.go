package main

import (
	"context"
	"log"
	"net"

	pb "../vaults"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type server struct{}

func (s *server) Event(ctx context.Context, in *pb.EventRequest) (*pb.EventReply, error) {
	log.Printf("Event Received")
	return nil, &pb.EventReply{Message: "Acknowledged Event"}

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEventsServer(s, %sever{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
