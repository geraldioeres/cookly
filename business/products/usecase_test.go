package products_test

import (
	"context"
	"cookly/business"
	"cookly/business/products"
	_mockProductsRepository "cookly/business/products/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository _mockProductsRepository.Repository

var productUseCase products.UseCase

func setup() {
	productUseCase = products.NewProductUseCase(&productRepository, 1)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestCreate(t *testing.T) {
	t.Run("Test Case 1 | Valid Test", func(t *testing.T) {
		domain := products.Domain{
			Name: "Onion",
		}
		productRepository.On("Create", mock.Anything, mock.Anything).Return(nil).Once()

		err := productUseCase.Create(context.Background(), &domain)

		assert.Nil(t, err)
	})

	t.Run("Test 2 | Error Create Product", func(t *testing.T) {
		errRepository := business.ErrorCreateData
		productRepository.On("Create", mock.Anything, mock.Anything).Return(errRepository).Once()

		err := productUseCase.Create(context.Background(), &products.Domain{})

		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		domain := products.Domain{
			ID: 1,
			Name: "Sugar",
		}
		productRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := productUseCase.Update(context.Background(), &domain, 1)

		assert.Nil(t, err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		productRepository.On("GetAll", mock.Anything).Return([]products.Domain{{ID: 1},{ID: 2},{ID: 3}}, nil).Once()

		result, err := productUseCase.GetAll(context.Background())

		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test 2 | Record Not Found", func(t *testing.T) {
		errRepository := business.ErrorDataNotFound
		productRepository.On("GetAll", mock.Anything).Return([]products.Domain{}, errRepository).Once()

		result, err := productUseCase.GetAll(context.Background())

		assert.Equal(t, errRepository, err)
		assert.Empty(t, result)
	})
}