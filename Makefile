gen:
	protoc --go_out=internal/generated/server --go_opt=paths=source_relative \
        --go-grpc_out=internal/generated/server --go-grpc_opt=paths=source_relative \
        proto/server.proto
