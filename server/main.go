package main

import (
	"grpctest/server/interfaces/student"
	pb "grpctest/server/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

// type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer lis.Close()

	s := grpc.NewServer()
	pb.RegisterCRUDServer(s, student.NewStudentInterfaces())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Server Started")
}
