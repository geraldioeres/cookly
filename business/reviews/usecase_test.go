package reviews_test

import (
	"context"
	"cookly/business/recipes"
	_mockRecipeRepository "cookly/business/recipes/mocks"
	"cookly/business/reviews"
	_mockReviewRepository "cookly/business/reviews/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	reviewRepository _mockReviewRepository.Repository
	recipeRepository _mockRecipeRepository.Repository
)
var reviewUseCase reviews.UseCase

var recipeDomain recipes.Domain

func setup() {
	reviewUseCase = reviews.NewReviewUseCase(&reviewRepository, &recipeRepository, 1)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestCreate(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		reviewRepository.On("CountReviews", mock.AnythingOfType("int")).Return(3, nil).Once()
		recipeRepository.On("RecipeByID", mock.Anything, mock.AnythingOfType("int")).Return(recipeDomain, nil).Once()
		reviewRepository.On("UpdateRecipeRating", mock.AnythingOfType("int"), mock.AnythingOfType("float64")).Return(nil)
		
		domain := reviews.Domain{
			UserID:   1,
			RecipeID: 1,
			Rating:   5.0,
			Comment:  "Easy to follow recipe",
		}
		reviewRepository.On("Create", mock.Anything, mock.Anything).Return(domain, nil).Once()

		result, err := reviewUseCase.Create(context.Background(), &domain)

		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})
}
