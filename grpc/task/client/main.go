package client

import (
	"google.golang.org/grpc"
	"log"
	"web-svc/grpc/task/server"
)

func Service() server.TaskManagerClient {
	con, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("cannot connect to grpc server: %v", err)
	}

	return server.NewTaskManagerClient(con)
}
