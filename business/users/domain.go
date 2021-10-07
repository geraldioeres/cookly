package users

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UseCase interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	Register(ctx context.Context, data *Domain) error
	GetUserByID(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	Register(ctx context.Context, data *Domain) error
	GetUserByID(ctx context.Context, id int) (Domain, error)
	GetUserByEmail(ctx context.Context, email string) (Domain, error)
}
