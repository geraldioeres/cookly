package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type UseCase interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
}