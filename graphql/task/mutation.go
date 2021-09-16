package task

import (
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
	"web-svc/grpc/task/server"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "TaskMutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type: Type,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"status": &graphql.ArgumentConfig{
					Type: StatusEnum,
				},
				"due_date": &graphql.ArgumentConfig{
					Type: graphql.DateTime,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				task, err := service.Create(context.Background(), &server.CreateTaskRequest{
					Name:        p.Args["name"].(string),
					Description: p.Args["description"].(string),
					Status:      p.Args["status"].(string),
					DueDate:     timestamppb.New(p.Args["due_date"].(time.Time)),
				})

				if err != nil {
					return nil, fmt.Errorf("error grpc client to create a task: %v", err)
				}

				return ToResponseItem(task), nil
			},
		},
		"update": &graphql.Field{
			Type: Type,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"status": &graphql.ArgumentConfig{
					Type: StatusEnum,
				},
				"due_date": &graphql.ArgumentConfig{
					Type: graphql.DateTime,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				task, err := service.Update(context.Background(), &server.UpdateTaskRequest{
					Id:          p.Args["id"].(string),
					Name:        p.Args["name"].(string),
					Description: p.Args["description"].(string),
					Status:      p.Args["status"].(string),
					DueDate:     timestamppb.New(p.Args["due_date"].(time.Time)),
				})

				if err != nil {
					return nil, fmt.Errorf("error grpc client to update a task: %v", err)
				}

				return ToResponseItem(task), nil
			},
		},
		"updateStatus": &graphql.Field{
			Type: Type,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"status": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(StatusEnum),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				task, err := service.UpdateStatus(context.Background(), &server.TaskStatusRequest{
					Id:     p.Args["id"].(string),
					Status: p.Args["status"].(string),
				})

				if err != nil {
					return nil, fmt.Errorf("error grpc client to update a task status: %v", err)
				}

				return ToResponseItem(task), nil
			},
		},
		"delete": &graphql.Field{
			Type: Type,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				_, err := service.Delete(context.Background(), &server.TaskIdRequest{
					Id: p.Args["id"].(string),
				})

				if err != nil {
					return nil, fmt.Errorf("error grpc client to delete a task: %v", err)
				}

				return nil, nil
			},
		},
	},
})
