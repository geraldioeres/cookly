package recipes

import (
	"cookly/business/recipes"
	"cookly/drivers/databases/categories"
	"cookly/drivers/databases/users"
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	ID               int     `gorm:"primaryKey" json:"id"`
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	Rating           float32 `json:"rating"`
	UserID           int     `json:"userId"`
	User             users.Users
	RecipeCategoryID int `json:"recipeCategoryId"`
	RecipeCategory   categories.Category
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `json:"deletedAt"`
}

func (record *Recipe) ToDomain() recipes.Domain {
	return recipes.Domain{
		ID:               record.ID,
		Title:            record.Title,
		Description:      record.Description,
		Rating:           record.Rating,
		UserID:           record.UserID,
		RecipeCategoryID: record.RecipeCategoryID,
		CreatedAt:        record.CreatedAt,
		UpdatedAt:        record.UpdatedAt,
		DeletedAt:        record.DeletedAt,
	}
}

func FromDomain(domain recipes.Domain) Recipe {
	return Recipe{
		ID:               domain.ID,
		Title:            domain.Title,
		Description:      domain.Description,
		Rating:           domain.Rating,
		UserID:           domain.UserID,
		RecipeCategoryID: domain.RecipeCategoryID,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
		DeletedAt:        domain.DeletedAt,
	}
}
