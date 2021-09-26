package recipes

import (
	recipecategory "cookly/models/recipe_category"
	"cookly/models/users"
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	ID               int     `gorm:"primaryKey" json:"id"`
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	Rating           float64 `json:"rating"`
	UserID           int     `json:"userId"`
	User             users.User
	RecipeCategoryID int `json:"recipeCategoryId"`
	RecipeCategory   recipecategory.RecipeCategory
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `json:"deletedAt"`
}
