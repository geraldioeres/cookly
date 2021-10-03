package requests

import (
	"cookly/business/ingredients"
	"cookly/business/recipes"
	"cookly/business/steps"
	reqIngre "cookly/controllers/ingredients/requests"
	reqStep "cookly/controllers/steps/requests"
)

type Recipe struct {
	Title            string                       `json:"title"`
	Description      string                       `json:"description"`
	UserID           int                          `json:"userId"`
	RecipeCategoryID int                          `json:"recipeCategoryId"`
	RecipeIngredient []reqIngre.IngredientRequest `json:"ingredients"`
	Step             []reqStep.StepRequest        `json:"steps"`
}

func (request *Recipe) ToDomain() *recipes.Domain {
	var ingredients []ingredients.IngredientDomain
	for _, ingredient := range request.RecipeIngredient {
		ingredients = append(ingredients, *ingredient.ToDomain())
	}

	var steps []steps.StepDomain
	for _, step := range request.Step {
		steps = append(steps, *step.ToDomain())
	}

	return &recipes.Domain{
		Title:            request.Title,
		Description:      request.Description,
		UserID:           request.UserID,
		RecipeCategoryID: request.RecipeCategoryID,
		RecipeIngredient: ingredients,
		Step:             steps,
	}
}
