package recipes

import (
	"cookly/mvc/models/recipe_category"
	"cookly/mvc/models/users"
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	ID               int    `gorm:"primaryKey" json:"id"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	Rating           int    `json:"rating"`
	UserID           int    `json:"userId"`
	User             users.User
	RecipeCategoryID int `json:"recipeCategoryId"`
	RecipeCategory   recipecategory.RecipeCategory
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `json:"deletedAt"`
	// RecipeIngredient []recipeingredients.RecipeIngredient `json:"ingredients"`
	// Step             []steps.Step                         `json:"cookingSteps" gorm:"foreignKey:RecipeID"`
	// Review           []reviews.Review                     `json:"reviews"`
}
