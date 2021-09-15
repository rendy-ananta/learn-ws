package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	taskServer "web-svc/grpc/task/server"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("cannot listening to the address: %v", err)
	}

	grpcServer := grpc.NewServer()

	taskServer.RegisterTaskManagerServer(grpcServer, taskServer.TaskManagerServerImpl{})

	fmt.Println("starting grpc server on :8080")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("cannot serve task grpc server: %v", err)
	}
}
