package main

import (
	"fmt"
	"myapp/products"
)

func main() {
	factory := products.Product{}

	product := factory.New()

	fmt.Println("My product was created at", product.CreatedAt.UTC() )
}