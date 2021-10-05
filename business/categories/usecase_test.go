package categories_test

import (
	"context"
	"cookly/business/categories"
	_mockCategoriesRepository "cookly/business/categories/mocks"
	"errors"
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
	t.Run("Test Case 1 | Valid test", func(t *testing.T) {
		domain := categories.Domain{
			Name: "Side-dish",
		}
		catRepository.On("Create", mock.Anything, mock.Anything).Return(nil).Once()

		err := catUseCase.Create(context.Background(), &domain)

		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Empty Name", func(t *testing.T) {
		errRepository := errors.New("category name is empty")
		catRepository.On("Create", mock.Anything, mock.Anything).Return(errRepository).Once()

		err := catUseCase.Create(context.Background(), &categories.Domain{})

		assert.Equal(t, errRepository, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1 | Valid Test", func(t *testing.T) {
		domain := categories.Domain{
			ID: 1,
			Name: "Soups",
		}
		catRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := catUseCase.Update(context.Background(), &domain, 1)

		assert.Nil(t, err)
	})
}
