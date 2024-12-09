package repository

import (
	"fmt"
	"os"
	"productsApi/src/pb/products"

	"google.golang.org/protobuf/proto"
)

type ProductRepository struct{}

const filename string = "products.txt"

func (pr *ProductRepository) loadData() (products.ProductList, error) {
	productList := products.ProductList{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return productList, fmt.Errorf("read file error: %w", err)
	}

	err = proto.Unmarshal(data, &productList)
	if err != nil {
		return productList, fmt.Errorf("unmarshal error: %w", err)
	}

	return productList, nil

}

func (pr *ProductRepository) saveData(productList products.ProductList) error {

	data, err := proto.Marshal(&productList)
	if err != nil {
		return fmt.Errorf("Marshal error: %w", err)
	}

	err = os.WriteFile(filename, data, 0664)
	if err != nil {
		return fmt.Errorf("WriteFile error: %w", err)
	}

	return nil
}

func (pr *ProductRepository) Create(products products.Product) (products.Product, error) {
	productList, err := pr.loadData()
	if err != nil {
		return products, fmt.Errorf("WriteFile error: %w", err)
	}

	products.Id = int32(len(productList.Products) + 1)
	productList.Products = append(productList.Products, &products)
	err = pr.saveData(productList)
	if err != nil {
		return products, fmt.Errorf("WriteFile error: %w", err)
	}

	return products, nil
}

func (pr *ProductRepository) FindAll() (products.ProductList, error) {
	return pr.loadData()
}
