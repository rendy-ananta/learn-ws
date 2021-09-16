package task

import (
	"github.com/graphql-go/graphql"
	"web-svc/grpc/task/client"
	"web-svc/grpc/task/server"
)

var service = client.Service()

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    RootQuery,
	Mutation: RootMutation,
})

func ToResponseItem(task *server.Task) ResponseItem {
	return ResponseItem{
		Id:          task.Id,
		Name:        task.Name,
		CreatedAt:   task.CreatedAt.AsTime(),
		Status:      task.Status,
		Description: task.Description,
		DueDate:     task.DueDate.AsTime(),
	}
}
