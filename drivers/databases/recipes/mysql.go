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

	result := repository.Conn.Preload(clause.Associations).Preload("RecipeIngredient."+clause.Associations).Find(&recipeByID, id)
	if result.Error != nil {
		return recipes.Domain{}, result.Error
	}

	return recipeByID.ToDomain(), nil
}

func (repository *mysqlRecipeRepository) GetAll(ctx context.Context) ([]recipes.Domain, error) {
	getAll := []Recipe{}

	result := repository.Conn.Preload(clause.Associations).Preload("RecipeIngredient." + clause.Associations).Find(&getAll)
	if result.Error != nil {
		return []recipes.Domain{}, result.Error
	}

	recipes := []recipes.Domain{}
	for _, recipe := range getAll {
		recipes = append(recipes, recipe.ToDomain())
	}

	return recipes, nil
}

func (repository *mysqlRecipeRepository) Update(ctx context.Context, recipeDomain *recipes.Domain) (recipes.Domain, error) {
	update := FromDomain(*recipeDomain)

	result := repository.Conn.Updates(update)
	if result.Error != nil {
		return recipes.Domain{}, result.Error
	}

	err := repository.Conn.Preload(clause.Associations).First(&update, update.ID)

	if err != nil {
		return recipes.Domain{}, result.Error
	}

	return update.ToDomain(), nil
}

func (repository *mysqlRecipeRepository) Search(ctx context.Context, title string) ([]recipes.Domain, error) {
	search := []Recipe{}
	result := repository.Conn.Where("title like ?", "%"+title+"%").Preload(clause.Associations).Preload("RecipeIngredient." + clause.Associations).Find(&search)
	if result.Error != nil {
		return []recipes.Domain{}, result.Error
	}

	var recipesDomain []recipes.Domain
	for _, get := range search {
		recipesDomain = append(recipesDomain, get.ToDomain())
	}

	return recipesDomain, nil
}
