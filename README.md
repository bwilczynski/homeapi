# Home API

## Installation

```
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1\n
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Generate protobuf

```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    server/lights/lights.proto
```

## Start server

```
go run cmd/server/main.go
```

## Test services

Make sure to have grpcurl installed:

```
brew install grpcurl
```

List services:

```
grpcurl -plaintext localhost:8000 list
```

Invoke methods:

```
grpcurl -plaintext localhost:8000 LightService/List
```
