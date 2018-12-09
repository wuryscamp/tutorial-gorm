package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/wuryscamp/tutorial-gorm/products/domain"
)

// ProductRepositoryGorm struct
type ProductRepositoryGorm struct {
	db *gorm.DB
}

// NewProductRepositoryGorm function
func NewProductRepositoryGorm(db *gorm.DB) *ProductRepositoryGorm {
	return &ProductRepositoryGorm{db: db}
}

// Save function
func (r *ProductRepositoryGorm) Save(product *domain.Product) Output {
	err := r.db.Save(product).Error
	if err != nil {
		return Output{Error: err}
	}

	return Output{Result: product}
}

// Delete function
func (r *ProductRepositoryGorm) Delete(product *domain.Product) Output {
	err := r.db.Delete(product).Error
	if err != nil {
		return Output{Error: err}
	}

	return Output{Result: product}
}

// FindByID function
func (r *ProductRepositoryGorm) FindByID(id string) Output {
	var product domain.Product

	err := r.db.Where(&domain.Product{ID: id}).Take(&product).Error
	if err != nil {
		return Output{Error: err}
	}

	return Output{Result: &product}
}

// FindAll function
func (r *ProductRepositoryGorm) FindAll() Output {
	var products domain.Products

	err := r.db.Find(&products).Error
	if err != nil {
		return Output{Error: err}
	}

	return Output{Result: products}
}
