package pb

/**
使用方法：
1、下载对应系统的 protoc，并且设置PATH
https://github.com/protocolbuffers/protobuf/releases

2、下载 proto-gen-to
go get -u google.golang.org/protobuf

3、下载 grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway

4、分别点击两个generate,生成对应的pb文件
*/

// 生成RPC pb.go
//go:generate protoc -I . --go_out=plugins=grpc:../services *.proto

// 生成gateway gw.go
//go:generate protoc --grpc-gateway_out=logtostderr=true:../services *.proto

// 此文件夹下面的 google/protobuf 为 protoc 带的文件，放到此处是复制过来 为了生成方便
// google/api|rpc 是 gateway 项目下的thirty_part目录下带的
