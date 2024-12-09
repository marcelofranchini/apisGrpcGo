package main

import (
	"context"
	"fmt"
	"log"
	"productsApiClient/src/pb/products"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connect, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("error on get client error: ", err)
	}

	defer connect.Close()

	createProduct(connect)

	finAllProducts(connect)
}

func finAllProducts(connect *grpc.ClientConn) {
	productClient := products.NewProductServiceClient(connect)

	productList, err := productClient.FindAll(context.Background(), &products.Product{})
	if err != nil {
		log.Fatalln("error on get client error: ", err)
	}
	fmt.Printf("products: %+v\n", productList)
}

func createProduct(connect *grpc.ClientConn) {

	newProduct := &products.Product{
		Name:        "TesteMarcelo",
		Description: "Produto Teste",
		Price:       288,
		Quantity:    8,
	}
	productClient := products.NewProductServiceClient(connect)

	productCreate, err := productClient.Create(context.Background(), newProduct)
	if err != nil {
		log.Fatalln("error on get client error: ", err)
	}

	fmt.Printf("products: %+v\n", productCreate)

}
