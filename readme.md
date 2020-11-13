1、下载对应系统的 protoc，并且设置PATH
https://github.com/protocolbuffers/protobuf/releases

2、下载 proto-gen-to
go get -u google.golang.org/protobuf

3、下载 grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway

4、运行 [doc.go](./pb/doc.go)中的generate