package responses

import (
	"cookly/business/categories"
	"time"

	"gorm.io/gorm"
)

type CategoryResponse struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func FromDomain(domain categories.Domain) CategoryResponse {
	return CategoryResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
