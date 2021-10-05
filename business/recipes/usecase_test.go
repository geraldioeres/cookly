package recipes_test

import (
	"context"
	"cookly/business"
	"cookly/business/ingredients"
	"cookly/business/recipes"
	_mockRecipeRepository "cookly/business/recipes/mocks"
	"cookly/business/steps"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var recipeRepository _mockRecipeRepository.Repository

var recipeUseCase recipes.UseCase

func setup() {
	recipeUseCase = recipes.NewRecipeUseCase(&recipeRepository, 1)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestCreate(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		domain := recipes.Domain{
			Title: "White Bean Chicken Soup",
			Description: "Modern take on classic chicken soup",
			UserID: 1,
			RecipeCategoryID: 1,
			RecipeIngredient: []ingredients.IngredientDomain{
				{
					ProductID: 1,
					Amount: "400 gr",
				},
				{
					ProductID: 2,
					Amount: "2",
				},
			},
			Step: []steps.StepDomain{
				{
					Order: 1,
					Instruction: "Boil the water for 3 minutes",
				},
				{
					Order: 2,
					Instruction: "Put the chicken into the boiled water",
				},
			},
		}
		recipeRepository.On("Create", mock.Anything, mock.Anything).Return(domain, nil).Once()

		result, err := recipeUseCase.Create(context.Background(), &domain)

		assert.Nil(t, err)
		assert.Equal(t, "White Bean Chicken Soup", result.Title)
	})
}

func TestRecipeByID(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		domain := recipes.Domain{
			ID: 1,
			Title: "Boiled Egg",
			Description: "Very easy daily recipe",
			UserID: 1,
			RecipeCategoryID: 1,
			RecipeIngredient: []ingredients.IngredientDomain{
				{
					ProductID: 1,
					Amount: "1",
				},
			},
			Step: []steps.StepDomain{
				{
					Order: 1,
					Instruction: "Boil the water for 3 minutes",
				},
			},
		}
		recipeRepository.On("RecipeByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := recipeUseCase.RecipeByID(context.Background(), 1)

		assert.Nil(t ,err)
		assert.Equal(t, domain.ID, result.ID)
	})

	t.Run("Test 2 | Data Not Found", func(t *testing.T) {
		errRepository := errors.New("Data Not Found")
		recipeRepository.On("RecipeByID", mock.Anything, mock.AnythingOfType("int")).Return(recipes.Domain{}, errRepository).Once()

		result, err := recipeUseCase.RecipeByID(context.Background(), 0)

		assert.Equal(t, errRepository, err)
		assert.Equal(t, result, recipes.Domain{})
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		recipeRepository.On("GetAll", mock.Anything).Return([]recipes.Domain{{ID: 1}, {ID: 2}, {ID: 3}}, nil).Once()

		result, err := recipeUseCase.GetAll(context.Background())

		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test 2 | Record Not Found", func(t *testing.T) {
		errRepository := business.ErrorDataNotFound
		recipeRepository.On("GetAll", mock.Anything).Return([]recipes.Domain{}, errRepository).Once()

		result, err := recipeUseCase.GetAll(context.Background())

		assert.Equal(t, errRepository, err)
		assert.Empty(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test 1 | Valid Test", func(t *testing.T) {
		recipeRepository.On("RecipeByID", mock.Anything, mock.AnythingOfType("int")).Return(recipes.Domain{}, nil).Once()
		var domain = recipes.Domain{
			ID: 1,
			Title: "Special Boiled Egg",
			Description: "easy daily recipe",
			UserID: 1,
			RecipeCategoryID: 1,
			RecipeIngredient: []ingredients.IngredientDomain{
				{
					ProductID: 1,
					Amount: "1",
				},
			},
			Step: []steps.StepDomain{
				{
					Order: 1,
					Instruction: "Boil the water for 3 minutes",
				},
			},
		}
		
		recipeRepository.On("Update", mock.Anything, mock.Anything).Return(domain, nil).Once()

		_, err := recipeUseCase.Update(context.Background(), &domain)
		assert.NoError(t, err)
	})
}