package requests

import "cookly/business/recipes"

type Recipe struct {
	ID               int     `json:"id"`
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	Rating           float32 `json:"rating"`
	UserID           int     `json:"userId"`
	RecipeCategoryID int     `json:"recipeCategoryId"`
	// RecipeIngredient []RecipeIngredient
	// Step             []steps.Step
}

func (request *Recipe) ToDomain() *recipes.Domain {
	return &recipes.Domain{
		ID:               request.ID,
		Title:            request.Title,
		Description:      request.Description,
		Rating:           request.Rating,
		UserID:           request.UserID,
		RecipeCategoryID: request.RecipeCategoryID,
	}
}
