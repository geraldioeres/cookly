package recipes

import (
	"context"
	"cookly/business/recipes"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mysqlRecipeRepository struct {
	Conn *gorm.DB
}

func NewMysqlRecipeRepository(conn *gorm.DB) recipes.Repository {
	return &mysqlRecipeRepository{
		Conn: conn,
	}
}

func (repository *mysqlRecipeRepository) Create(ctx context.Context, recipeDomain *recipes.Domain) (recipes.Domain, error) {
	create := FromDomain(*recipeDomain)

	result := repository.Conn.Create(&create)
	if result.Error != nil {
		return recipes.Domain{}, result.Error
	}

	err := repository.Conn.Preload(clause.Associations).Preload("RecipeIngredient."+clause.Associations).First(&create, create.ID).Error
	if err != nil {
		return recipes.Domain{}, result.Error
	}

	return create.ToDomain(), nil
}

func (repository *mysqlRecipeRepository) RecipeByID(ctx context.Context, id int) (recipes.Domain, error) {
	recipeByID := Recipe{}

	result := repository.Conn.Preload(clause.Associations).Find(&recipeByID, id)
	if result.Error != nil {
		return recipes.Domain{}, result.Error
	}

	return recipeByID.ToDomain(), nil
}
