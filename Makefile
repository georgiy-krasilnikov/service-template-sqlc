SERVICE_NAME=grpc-service-template-sqlc
SERVICE_PORT=18000
PWD=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

up:
	docker build -t ${SERVICE_NAME} --build-arg SERVICE_NAME=${SERVICE_NAME} .
	docker run --network=st-network -p ${SERVICE_PORT}:${SERVICE_PORT} --env-file=.env ${SERVICE_NAME}

protoc:
	docker pull rvolosatovs/protoc:latest
	docker run --rm -v ${PWD}:/src -w /src rvolosatovs/protoc --proto_path=/src --go_out=. --go-grpc_out=require_unimplemented_servers=false:. proto/${SERVICE_NAME}.proto

lint:
	docker run --rm -v ${PWD}:/app -w /app golangci/golangci-lint:v1.50.1 golangci-lint run -v

test:
	go test ./...

sqlc:
	docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate