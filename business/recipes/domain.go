package recipes

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID               int
	Title            string
	Description      string
	Rating           float32
	UserID           int
	RecipeCategoryID int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	// RecipeIngredient []RecipeIngredient
	// Step             []steps.Step
}

type UseCase interface {
	Create(ctx context.Context, data *Domain) error 
}

type Repository interface {
	Create(ctx context.Context, data *Domain) error 
}