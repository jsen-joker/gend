# gend-embed

gend rpc版

protoc --go_out=plugins=grpc:. GendRpc.proto

go build -o gend-grpc-server main.go