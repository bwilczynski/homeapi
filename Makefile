.PHONY: protogen
protogen:
	find proto/ -name *.proto -printf '%f\n' | xargs protoc \
		--proto_path=proto \
		--go_out=. --go_opt=module=github.com/bwilczynski/homeapi \
    	--go-grpc_out=. --go-grpc_opt=module=github.com/bwilczynski/homeapi \
    	lights.proto

.PHONY: start
start:
	go run cmd/server/main.go
