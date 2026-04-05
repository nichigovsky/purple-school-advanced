package products

import (
	"github.com/lib/pq"
)

type ProductCreateRequest struct {
	Name        string `json:"name" validate:"required"`
    Description string `json:"description" validate:"required"`
	Images 		pq.StringArray `json:"images"`
}

type ProductUpdateRequest struct {
	Name        string `json:"name" validate:"required"`
    Description string `json:"description"`
	Images 		pq.StringArray `json:"images"`
}

type ProductResponse struct {
    ID          uint           `json:"ID"`
    Name        string         `json:"name"`
    Description string         `json:"description"`
    Images      pq.StringArray `json:"images"`
}