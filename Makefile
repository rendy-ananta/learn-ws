build:
	go build -o bin/grpc web-svc/cmd/grpc
	go build -o bin/graphql web-svc/cmd/graphql

test:
	go test web-svc/db/repository
