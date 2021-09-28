package recipeingredients

import (
	"time"

	"gorm.io/gorm"
)

type RecipeIngredient struct {
	ID        int            `json:"id"`
	RecipeID  int            `json:"recipeId"`
	PoductID  int            `json:"ProductId"`
	Amount    string         `json:"amount"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
