package users

import (
	"context"
	"cookly/app/middleware"
	"cookly/business"
	"cookly/helpers/encrypt"
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
	checkUser, err := uc.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return Domain{}, err
	}

	if !encrypt.HashValidation(password, checkUser.Password) {
		return Domain{}, business.ErrorInvalidPassword
	}

	checkUser.Token = uc.jwtAuth.GenerateToken(checkUser.ID)
	return checkUser, nil
}

func (uc *UserUseCase) Register(ctx context.Context, userDomain *Domain) error {
	if userDomain.Password == "" {
		return business.ErrorInvalidPassword
	}
	
	var err error
	userDomain.Password, err = encrypt.Hash(userDomain.Password)
	if err != nil {
		return err
	}

	err = uc.userRepo.Register(ctx, userDomain)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCase) GetUserByID(ctx context.Context, id int) (Domain, error) {
	result, err := uc.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
