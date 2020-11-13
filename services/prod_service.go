package services

import (
	"context"
	"google.golang.org/grpc/metadata"
	"log"
)

type ProdService struct{}

func (this *ProdService) GetProdStock(ctx context.Context, in *ProdRequest) (*ProdResponse, error) {
	log.Print(in)
	m ,_ := metadata.FromIncomingContext(ctx)
	log.Println(m["tracing"])
	return &ProdResponse{ProdStock: 20}, nil
}

func (this *ProdService) GetProdStockSize(ctx context.Context, in *QuerySize) (*ProdResponseList, error) {
	log.Println(in)
	return &ProdResponseList{Prodres: []*ProdResponse{&ProdResponse{ProdStock: 20}, &ProdResponse{ProdStock: 21}}}, nil
}
