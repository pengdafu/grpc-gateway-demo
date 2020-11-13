package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	service "github.com/pengdafu/grpc-demo/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
)

func main() {
	gwmux := runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			return metadata.Pairs("tracing", request.Header.Get("tracing"))
		}),
	)

	opt := []grpc.DialOption{grpc.WithInsecure()}
	err := service.RegisterProdServiceHandlerFromEndpoint(context.Background(), gwmux, ":8081", opt)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/*any", func(c *gin.Context) {
		c.Request.Header.Set("tracing", "ing")
	},func(c *gin.Context) {
		gwmux.ServeHTTP(c.Writer, c.Request)
	})
	r.Run(":8080")
}
