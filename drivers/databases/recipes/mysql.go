package recipes

import (
	"context"
	"cookly/business/recipes"

	"gorm.io/gorm"
)

type mysqlRecipeRepository struct {
	Conn *gorm.DB
}

func NewMysqlRecipeRepository(conn *gorm.DB) recipes.Repository {
	return &mysqlRecipeRepository{
		Conn: conn,
	}
}

func (repository *mysqlRecipeRepository) Create(ctx context.Context, recipeDomain *recipes.Domain) error {
	create := FromDomain(*recipeDomain)

	result := repository.Conn.Create(&create)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
