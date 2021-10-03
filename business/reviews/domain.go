package reviews

import (
	"context"
	"cookly/business/recipes"
	"cookly/business/users"
	"time"
)

type Domain struct {
	ID       int
	UserID   int
	User     users.Domain
	RecipeID int
	Recipe   recipes.Domain
	Rating   float64
	Comment  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UseCase interface {
	Create(ctx context.Context, data *Domain) (Domain, error)
	GetReviewsByRecipeID(ctx context.Context, recipeId int) ([]Domain, error)
}

type Repository interface {
	Create(ctx context.Context, data *Domain) (Domain, error)
	GetReviewsByRecipeID(ctx context.Context, recipeId int) ([]Domain, error)
}
