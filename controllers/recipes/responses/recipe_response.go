package responses

import (
	"cookly/business/recipes"
	resIngre "cookly/controllers/ingredients/responses"
	resStep "cookly/controllers/steps/responses"
	"time"
)

type RecipeResponse struct {
	ID               int                           `json:"id"`
	Title            string                        `json:"title"`
	Description      string                        `json:"description"`
	Rating           float32                       `json:"rating"`
	UserID           int                           `json:"userId"`
	PublishedBy      string                        `json:"publishedBy"`
	RecipeCategoryID int                           `json:"recipeCategoryId"`
	Category         string                        `json:"category"`
	CreatedAt        time.Time                     `json:"createdAt"`
	UpdatedAt        time.Time                     `json:"updatedAt"`
	RecipeIngredient []resIngre.IngredientResponse `json:"ingredients"`
	Step             []resStep.StepResponse        `json:"steps"`
}

func FromDomain(domain recipes.Domain) RecipeResponse {
	var ingredients []resIngre.IngredientResponse
	for _, getIngre := range domain.RecipeIngredient {
		ingredients = append(ingredients, resIngre.FromIngredientDomain(getIngre))
	}

	var steps []resStep.StepResponse
	for _, getStep := range domain.Step {
		steps = append(steps, resStep.FromStepDomain(getStep))
	}

	return RecipeResponse{
		ID:               domain.ID,
		Title:            domain.Title,
		Description:      domain.Description,
		Rating:           domain.Rating,
		UserID:           domain.UserID,
		PublishedBy:      domain.User.Name,
		RecipeCategoryID: domain.RecipeCategoryID,
		Category:         domain.RecipeCategory.Name,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
		RecipeIngredient: ingredients,
		Step:             steps,
	}
}
