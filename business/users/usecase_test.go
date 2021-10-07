package users_test

import (
	"context"
	"cookly/app/middleware"
	"cookly/business"
	"cookly/business/users"
	_mockUserRepository "cookly/business/users/mocks"
	"cookly/helpers/encrypt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRepository _mockUserRepository.Repository
	userUseCase    users.UseCase
	jwtAuth        *middleware.ConfigJWT
)

func setup() {
	jwtAuth = &middleware.ConfigJWT{SecretJWT: "usertest111", ExiresDuration: 1}
	userUseCase = users.NewUserUseCase(&userRepository, 1, jwtAuth)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestRegister(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Name:     "gerald",
			Email:    "gerald@gmail.com",
			Password: "12345",
		}
		userRepository.On("Register", mock.Anything, mock.Anything).Return(nil).Once()

		err := userUseCase.Register(context.Background(), &domain)

		assert.Nil(t, err)
	})

	t.Run("Test 2 | Password Error", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Name:     "gerald",
			Email:    "gerald@gmail.com",
			Password: "",
		}
		userRepository.On("Register", mock.Anything, mock.Anything).Return(nil).Once()

		err := userUseCase.Register(context.Background(), &domain)

		assert.Equal(t, err, business.ErrorInvalidPassword)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		hashedPassword, _ := encrypt.Hash("12345")
		domain := users.Domain{
			ID:       1,
			Name:     "gerald",
			Email:    "gerald@gmail.com",
			Password: hashedPassword,
		}
		userRepository.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(domain, nil).Once()

		_, err := userUseCase.Login(context.Background(), "gerald@gmail.com", "12345")

		assert.Nil(t, err)
	})

	t.Run("Test 2 | Invalid Password", func(t *testing.T) {
		hashedPassword, _ := encrypt.Hash("54321")
		domain := users.Domain{
			ID:       1,
			Name:     "gerald",
			Email:    "gerald@gmail.com",
			Password: hashedPassword,
		}
		userRepository.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(domain, nil).Once()

		_, err := userUseCase.Login(context.Background(), "gerald@gmail.com", "12345")

		assert.Equal(t, err, business.ErrorInvalidPassword)
	})

	t.Run("Test 3 | User Not Found", func(t *testing.T) {
		errRepository := business.ErrorDataNotFound
		userRepository.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errRepository).Once()

		_, err := userUseCase.Login(context.Background(), "", "")

		assert.Error(t, err)
	})
}

func TestGetUserByID(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		domain := users.Domain{
			ID: 1,
			Name: "gerald",
			Email: "geral@gmail.com",
		}
		userRepository.On("GetUserByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := userUseCase.GetUserByID(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, domain.ID, result.ID)
	})

	t.Run("Test 2 | Data Not Found", func(t *testing.T) {
		errRepository := business.ErrorDataNotFound
		userRepository.On("GetUserByID", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errRepository).Once()

		result, err := userUseCase.GetUserByID(context.Background(), 0)

		assert.Equal(t, errRepository, err)
		assert.Equal(t, result, users.Domain{})
	})
}