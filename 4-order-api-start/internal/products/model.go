package products

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
    gorm.Model
    Name        string `json:"name"`
    Description string `json:"description"`
    Images      pq.StringArray `json:"images" gorm:"type:text[]"`
}

func NewProduct(product ProductCreateRequest) *Product {
	newProduct := &Product{
		Name: product.Name,
		Description: product.Description,
		Images: product.Images,
	}

	return newProduct
}