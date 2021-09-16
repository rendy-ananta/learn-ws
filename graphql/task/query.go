package task

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/graphql-go/graphql"
	taskServer "web-svc/grpc/task/server"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "TaskQuery",
	Fields: graphql.Fields{
		"getList": &graphql.Field{
			Type: graphql.NewList(Type),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				tasks, err := service.GetAll(context.Background(), &empty.Empty{})

				if err != nil {
					return nil, fmt.Errorf("error grpc client to get taks list: %v", err)
				}

				list := make([]ResponseItem, 0)

				for _, item := range tasks.List {
					list = append(list, ToResponseItem(item))
				}

				return list, nil
			},
		},
		"find": &graphql.Field{
			Type: Type,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(string)

				task, err := service.Find(context.Background(), &taskServer.TaskIdRequest{Id: id})

				if err != nil {
					return nil, fmt.Errorf("error grpc client to find task: %v", err)
				}

				return ToResponseItem(task), nil
			},
		},
	},
})
