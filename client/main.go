package main

import (
	"context"
	// f "fmt"
	"log"

	pb "grpctest/client/proto"

	"google.golang.org/grpc"
)

const (
	port = `127.0.0.1:50053`
)

func main() {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(`Did not connect to server: %v`, err)
	}
	defer conn.Close()

	c := pb.NewCRUDClient(conn)

	// create student

	data := &pb.Student{}
	data.Id = 6
	data.Name = "dori"
	data.Grade = 90

	c.Create(context.Background(), data)
	if err != nil {
		log.Fatalf("can't Create data", err)
	}
	log.Println("Create Success")

	// Read Last Student

	r, err := c.Read(context.Background(), &pb.StudentID{})
	if err != nil {
		log.Fatalf("can't get data", err)
	}
	log.Println("Get students: ", r.Id, r.Name)

	//delete Student

	c.Delete(context.Background(), &pb.StudentID{Id: 5})
	if err != nil {
		log.Fatalf("can't delete data", err)
	}
	log.Printf("delete success")

	// update student

	data := &pb.Student{}
	data.Id = 6
	data.Name = "duuurrrrii"
	data.Grade = 101

	c.Update(context.Background(), data)
	if err != nil {
		log.Fatalf("can't Edit data", err)
	}
	log.Printf("Edit success")

}
