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
