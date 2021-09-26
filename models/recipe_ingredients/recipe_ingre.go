package recipeingredients

import (
	"time"

	"gorm.io/gorm"
)

type RecipeIngredient struct {
	ID           int            `json:"id"`
	RecipeID     int            `json:"recipeId"`
	IngredientID int            `json:"ingredientId"`
	Amount       string         `json:"amount"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt"`
}
