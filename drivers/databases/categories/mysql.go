package categories

import (
	"context"
	"cookly/business/categories"

	"gorm.io/gorm"
)

type mysqlCategoryRepository struct {
	Conn *gorm.DB
}

func NewMysqlCategoryRepository(conn *gorm.DB) categories.Repository {
	return &mysqlCategoryRepository{
		Conn: conn,
	}
}

func (repository *mysqlCategoryRepository) Create(ctx context.Context, categoryDomain *categories.Domain) error {
	create := FromDomain(*categoryDomain)

	result := repository.Conn.Create(&create)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlCategoryRepository) GetAll(ctx context.Context) ([]categories.Domain, error) {
	var getAll []Category

	result := repository.Conn.Find(&getAll)
	if result.Error != nil {
		return []categories.Domain{}, result.Error
	}

	var categories []categories.Domain
	for _, get := range getAll {
		categories = append(categories, get.ToDomain())
	}

	return categories, nil
}