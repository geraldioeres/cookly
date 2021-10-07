package recipes

import "cookly/mvc/models/steps"

type CreateRecipe struct {
	Title            string             `json:"title"`
	Description      string             `json:"description"`
	UserID           int                `json:"userId"`
	RecipeCategoryID int                `json:"recipeCategoryId"`
	RecipeIngredient []RecipeIngredient `json:"ingredients"`
	Step             []steps.Step       `json:"cookingSteps"`
}
