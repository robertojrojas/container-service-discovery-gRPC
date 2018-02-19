First [install protoc](https://github.com/google/protobuf/blob/master/README.md)

To generate the Go gRPC sources:

`protoc -I ./pb ./pb/file/service.proto --go_out=plugins=grpc:.`
