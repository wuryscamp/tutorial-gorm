package domain

import "time"

// Product domain
type Product struct {
	ID        string `gorm:"primary_key"`
	Name      string
	Quantity  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   uint
}

// Products type list of Product
type Products []Product
