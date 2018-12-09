package main

import (
	"fmt"
	"os"

	"github.com/wuryscamp/tutorial-gorm/database"
	"github.com/wuryscamp/tutorial-gorm/products/domain"
	"github.com/wuryscamp/tutorial-gorm/products/repository"
)

func main() {
	fmt.Println("TUTORIAL GORM")

	db, err := database.GetGormConn("localhost", "wurianto", "tutorial_gorm", "123456", 5432)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.DB().Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	productRepository := repository.NewProductRepositoryGorm(db)

	//insertProduct(productRepository)
	//updateProduct(productRepository)
	//deleteProduct(productRepository)

	//getProductByID(productRepository)
	getProducts(productRepository)
}

func insertProduct(r repository.ProductRepository) {

	p := &domain.Product{
		ID:       "002",
		Name:     "Nokia 6",
		Quantity: 3,
		Version:  1,
	}

	output := r.Save(p)

	if output.Error != nil {
		fmt.Println(output.Error.Error())
		os.Exit(1)
	}

	product, ok := output.Result.(*domain.Product)
	if !ok {
		fmt.Println("result is not a product")
	}

	fmt.Println(product)
}

func updateProduct(r repository.ProductRepository) {
	output := r.FindByID("001")

	if output.Error != nil {
		fmt.Println(output.Error.Error())
		os.Exit(1)
	}

	product, ok := output.Result.(*domain.Product)
	if !ok {
		fmt.Println("result is not a product")
	}

	product.Name = "Samsung Galaxy S2"
	product.Quantity = 4
	product.Version++

	outputUpdate := r.Save(product)

	if outputUpdate.Error != nil {
		fmt.Println(outputUpdate.Error.Error())
	}

	productUpdated, ok := outputUpdate.Result.(*domain.Product)

	if !ok {
		fmt.Println("result is not a product")
	}

	fmt.Println(productUpdated)

}

func deleteProduct(r repository.ProductRepository) {
	output := r.FindByID("001")

	if output.Error != nil {
		fmt.Println(output.Error.Error())
		os.Exit(1)
	}

	product, ok := output.Result.(*domain.Product)
	if !ok {
		fmt.Println("result is not a product")
	}

	outputDelete := r.Delete(product)
	if outputDelete.Error != nil {
		fmt.Println(outputDelete.Error.Error())
		os.Exit(1)
	}

	productDeleted, ok := outputDelete.Result.(*domain.Product)
	if !ok {
		fmt.Println("result is not a product")
	}

	fmt.Println(productDeleted)
}

func getProductByID(r repository.ProductRepository) {
	output := r.FindByID("002")

	if output.Error != nil {
		fmt.Println(output.Error.Error())
		os.Exit(1)
	}

	product, ok := output.Result.(*domain.Product)
	if !ok {
		fmt.Println("result is not a product")
	}

	fmt.Println(product)
}

func getProducts(r repository.ProductRepository) {
	output := r.FindAll()

	if output.Error != nil {
		fmt.Println(output.Error.Error())
		os.Exit(1)
	}

	products, ok := output.Result.(domain.Products)
	if !ok {
		fmt.Println("result is not list of product")
	}

	for _, v := range products {
		fmt.Println(v)
	}
}
