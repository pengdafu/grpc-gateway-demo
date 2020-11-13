package main

import (
	service "github.com/pengdafu/grpc-demo/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	//s := grpc.NewServer(grpc.Creds(helper.GetServerCreds()))
	s := grpc.NewServer()
	service.RegisterProdServiceServer(s, new(service.ProdService))

	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	s.Serve(lis)
}
