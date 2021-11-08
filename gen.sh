protoc --proto_path=./assets/protos \
  --go_out=./internal/rpc \
  --go_opt=paths=source_relative \
  --go-grpc_out=./internal/rpc \
  --go-grpc_opt=paths=source_relative \
  ./assets/protos/*.proto