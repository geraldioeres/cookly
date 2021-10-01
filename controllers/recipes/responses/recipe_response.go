package responses

import (
	"cookly/business/recipes"
	"time"

	"gorm.io/gorm"
)

type RecipeResponse struct {
	ID               int            `json:"id"`
	Title            string         `json:"title"`
	Description      string         `json:"description"`
	Rating           float32        `json:"rating"`
	UserID           int            `json:"userId"`
	RecipeCategoryID int            `json:"recipeCategoryId"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `json:"-"`
	// RecipeIngredient []RecipeIngredient
	// Step             []steps.Step
}

func FromDomain(domain recipes.Domain) RecipeResponse {
	return RecipeResponse{
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
