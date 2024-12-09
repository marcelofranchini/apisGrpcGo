package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"productsApi/src/pb/products"
	"productsApi/src/repository"

	"google.golang.org/grpc"
)

type server struct {
	products.ProductServiceServer
	productRepository *repository.ProductRepository
}

func (s server) Create(ctx context.Context, product *products.Product) (*products.Product, error) {

	newProduct, err := s.productRepository.Create(*product)
	if err != nil {
		return product, err
	}

	return &newProduct, nil

}

func (s server) FindAll(ctx context.Context, product *products.Product) (*products.ProductList, error) {
	productList, err := s.productRepository.FindAll()
	fmt.Println("findd")
	return &productList, err

}
func main() {
	fmt.Println("Start ProductsApi")
	srv := server{productRepository: &repository.ProductRepository{}}

	listener, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Fatalln("erro on create listener. error:  ", err)
	}

	s := grpc.NewServer()

	products.RegisterProductServiceServer(s, &srv)

	if err := s.Serve(listener); err != nil {
		log.Fatalln("error on server error: ", err)
	}

}
