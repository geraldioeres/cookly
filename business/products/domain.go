package products

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type UseCase interface {
	Create(ctx context.Context, data *Domain) error
	GetAll(ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, data *Domain, id int) error
}

type Repository interface {
	Create(ctx context.Context, data *Domain) error
	GetAll(ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, data *Domain, id int) error
}
