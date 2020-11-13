package main

import (
	"context"
	"fmt"
	service "github.com/pengdafu/grpc-demo/services"
	"google.golang.org/grpc"
	"log"
)

func main() {

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal("启动失败", err)
	}
	defer conn.Close()

	c := service.NewProdServiceClient(conn)
	fmt.Println(c.GetProdStock(context.Background(), &service.ProdRequest{ProdId: 12}))
}
