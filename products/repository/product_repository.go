package repository

import (
	"github.com/wuryscamp/tutorial-gorm/products/domain"
)

// Output struct
type Output struct {
	Result interface{}
	Error  error
}

// ProductRepository interface
type ProductRepository interface {
	Save(*domain.Product) Output
	Delete(*domain.Product) Output
	FindByID(string) Output
	FindAll() Output
}
