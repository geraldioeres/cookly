package categories_test

import (
	"context"
	"cookly/business"
	"cookly/business/categories"
	_mockCategoriesRepository "cookly/business/categories/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var catRepository _mockCategoriesRepository.Repository

var catUseCase categories.UseCase

func setup() {
	catUseCase = categories.NewCategoryUseCase(&catRepository, 1)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestCreate(t *testing.T) {
	t.Run("Test Case 1 | Valid Test", func(t *testing.T) {
		domain := categories.Domain{
			Name: "Side-dish",
		}
		catRepository.On("Create", mock.Anything, mock.Anything).Return(nil).Once()

		err := catUseCase.Create(context.Background(), &domain)

		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Empty Name", func(t *testing.T) {
		errRepository := business.ErrorEmptyName
		catRepository.On("Create", mock.Anything, mock.Anything).Return(errRepository).Once()

		err := catUseCase.Create(context.Background(), &categories.Domain{})

		assert.Equal(t, errRepository, err)
	})

	t.Run("Test Case 3 | Error Create Categories", func(t *testing.T) {
		errRepository := business.ErrorCreateData
		catRepository.On("Create", mock.Anything, mock.Anything).Return(errRepository).Once()

		err := catUseCase.Create(context.Background(), &categories.Domain{})

		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1 | Valid Test", func(t *testing.T) {
		domain := categories.Domain{
			ID:   1,
			Name: "Soups",
		}
		catRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := catUseCase.Update(context.Background(), &domain, 1)

		assert.Nil(t, err)
	})

	t.Run("Test 2 | Invalid ID", func(t *testing.T) {
		errRepository := business.ErrorInvalidCategoryID
		catRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(errRepository).Once()

		err := catUseCase.Update(context.Background(), &categories.Domain{}, 0)

		assert.Error(t, err)

	})
}

func TestGetAll(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		catRepository.On("GetAll", mock.Anything).Return([]categories.Domain{{ID: 1}, {ID: 2}, {ID: 3}}, nil).Once()

		result, err := catUseCase.GetAll(context.Background())

		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test 2 | Record Not Found", func(t *testing.T) {
		errRepository := business.ErrorDataNotFound
		catRepository.On("GetAll", mock.Anything).Return([]categories.Domain{}, errRepository).Once()

		result, err := catUseCase.GetAll(context.Background())

		assert.Equal(t, errRepository, err)
		assert.Equal(t, []categories.Domain{}, result)
	})
}
