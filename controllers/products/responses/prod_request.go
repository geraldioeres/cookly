package responses

import (
	"cookly/business/products"
	"time"

	"gorm.io/gorm"
)

type ProductResponse struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func FromDomain(domain products.Domain) ProductResponse {
	return ProductResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}