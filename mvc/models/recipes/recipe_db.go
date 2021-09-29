package recipes

import (
	recipecategory "cookly/mvc/models/recipe_category"
	"cookly/mvc/models/users"
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	ID               int                           `gorm:"primaryKey" json:"id"`
	Title            string                        `json:"title"`
	Description      string                        `json:"description"`
	Rating           float32                       `json:"rating"`
	UserID           int                           `json:"userId"`
	User             users.User                    `json:"user"`
	RecipeCategoryID int                           `json:"recipeCategoryId"`
	RecipeCategory   recipecategory.RecipeCategory `json:"recipeCategory"`
	CreatedAt        time.Time                     `json:"createdAt"`
	UpdatedAt        time.Time                     `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt                `json:"deletedAt"`
	//RecipeIngredient []recipeingredients.RecipeIngredient `json:"ingredients" gorm:"foreignKey:RecipeID"`
	// Step             []steps.Step                         `json:"cookingSteps" gorm:"foreignKey:RecipeID"`
	// Review []reviews.Review `json:"reviews" gorm:"foreignKey:RecipeID"`
}


