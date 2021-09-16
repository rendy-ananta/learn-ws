package main

import (
	"github.com/graphql-go/handler"
	"log"
	"net/http"
	"web-svc/graphql/task"
)

func main() {
	h := handler.New(&handler.Config{
		Schema: &task.Schema,
	})

	http.Handle("/graphql", h)

	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatalf("cannot start http server: %v", err)
	}
}
