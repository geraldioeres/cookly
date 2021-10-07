package recipes

import (
	ingredientDomain "cookly/business/ingredients"
	"cookly/business/recipes"
	stepDomain "cookly/business/steps"
	"cookly/drivers/databases/categories"
	ingreDb "cookly/drivers/databases/ingredients"
	stepDb "cookly/drivers/databases/steps"
	"cookly/drivers/databases/users"
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	ID               int `gorm:"primaryKey" json:"id"`
	Title            string
	Description      string
	Rating           float64
	UserID           int
	User             users.Users
	RecipeCategoryID int
	RecipeCategory   categories.Category
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	RecipeIngredient []ingreDb.Ingredient
	Step             []stepDb.Step
}

func (record *Recipe) ToDomain() recipes.Domain {
	var ingredients []ingredientDomain.IngredientDomain
	for _, ingre := range record.RecipeIngredient {
		ingredients = append(ingredients, ingre.ToDomain())
	}

	var steps []stepDomain.StepDomain
	for _, step := range record.Step {
		steps = append(steps, step.ToDomain())
	}

	return recipes.Domain{
		ID:               record.ID,
		Title:            record.Title,
		Description:      record.Description,
		Rating:           record.Rating,
		UserID:           record.UserID,
		User:             record.User.ToDomain(),
		RecipeCategoryID: record.RecipeCategoryID,
		RecipeCategory:   record.RecipeCategory.ToDomain(),
		CreatedAt:        record.CreatedAt,
		UpdatedAt:        record.UpdatedAt,
		DeletedAt:        record.DeletedAt,
		RecipeIngredient: ingredients,
		Step:             steps,
	}
}

func FromDomain(domain recipes.Domain) Recipe {
	var ingredients []ingreDb.Ingredient
	for _, ingre := range domain.RecipeIngredient {
		ingredients = append(ingredients, *ingreDb.FromDomain(&ingre))
	}

	var steps []stepDb.Step
	for _, step := range domain.Step {
		steps = append(steps, *stepDb.FromDomain(&step))
	}

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
		RecipeIngredient: ingredients,
		Step:             steps,
	}
}
