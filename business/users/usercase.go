package users

import (
	"context"
	"time"
)

type UserUseCase struct {
	userRepo       Repository
	contextTimeout time.Duration
}

func NewUserUseCase(repo Repository, timeout time.Duration) UseCase {
	return &UserUseCase{
		userRepo:       repo,
		contextTimeout: timeout,
	}
}

func (uc *UserUseCase) Login(ctx context.Context, email string, password string) (Domain, error) {
	result, err := uc.userRepo.Login(ctx, email, password)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}
