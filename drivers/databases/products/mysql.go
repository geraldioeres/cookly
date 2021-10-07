package products

import (
	"context"
	"cookly/business/products"

	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	Conn *gorm.DB
}

func NewMysqlProductRepository(conn *gorm.DB) products.Repository {
	return &mysqlProductRepository{
		Conn: conn,
	}
}

func (repository *mysqlProductRepository) Create(ctx context.Context, productDomain *products.Domain) error {
create := FromDomain(*productDomain)

result := repository.Conn.Create(&create)
if result.Error != nil {
	return result.Error
}

	return nil
}

func (repository *mysqlProductRepository) GetAll(ctx context.Context) ([]products.Domain, error) {
	var getAll []Product

	result := repository.Conn.Find(&getAll)
	if result.Error != nil {
		return []products.Domain{}, result.Error
	}

	var products []products.Domain
	for _, get := range getAll {
		products = append(products, get.ToDomain())
	}

	return products, nil
}

func (repository *mysqlProductRepository) Update(ctx context.Context, productDomain *products.Domain, id int) error {
	update := FromDomain(*productDomain)

	result := repository.Conn.Where("id = ?", id).Updates(&update)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

