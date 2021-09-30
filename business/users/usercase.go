package users

import (
	"context"
	"cookly/app/middleware"
	"time"
)

type UserUseCase struct {
	userRepo       Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
}

func NewUserUseCase(repo Repository, timeout time.Duration, jwtauth *middleware.ConfigJWT) UseCase {
	return &UserUseCase{
		userRepo:       repo,
		contextTimeout: timeout,
		jwtAuth:        jwtauth,
	}
}

func (uc *UserUseCase) Login(ctx context.Context, email string, password string) (Domain, error) {
	result, err := uc.userRepo.Login(ctx, email, password)
	if err != nil {
		return Domain{}, err
	}
	result.Token = uc.jwtAuth.GenerateToken(result.ID)
	return result, nil
}
