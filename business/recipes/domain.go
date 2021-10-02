package recipes

import (
	"context"
	"cookly/business/categories"
	"cookly/business/ingredients"
	"cookly/business/steps"
	"cookly/business/users"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID               int
	Title            string
	Description      string
	Rating           float32
	UserID           int
	User             users.Domain
	RecipeCategoryID int
	RecipeCategory   categories.Domain
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	RecipeIngredient []ingredients.IngredientDomain
	Step             []steps.StepDomain
}

type UseCase interface {
	Create(ctx context.Context, data *Domain) (Domain, error)
	RecipeByID(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Create(ctx context.Context, data *Domain) (Domain, error)
	RecipeByID(ctx context.Context, id int) (Domain, error)
}
