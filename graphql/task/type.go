package task

import (
	"github.com/graphql-go/graphql"
	"time"
	"web-svc/db/repository"
)

type ResponseItem struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
}

var Type = graphql.NewObject(graphql.ObjectConfig{
	Name: "Task",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"due_date": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var StatusEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "TaskStatus",
	Values: graphql.EnumValueConfigMap{
		"on_progress": &graphql.EnumValueConfig{
			Value: repository.StatusOnProgress,
		},
		"done": &graphql.EnumValueConfig{
			Value: repository.StatusDone,
		},
		"not_started": &graphql.EnumValueConfig{
			Value: repository.StatusNotStarted,
		},
		"closed": &graphql.EnumValueConfig{
			Value: repository.StatusClosed,
		},
	},
})
