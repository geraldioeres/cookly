package reviews

import (
	"context"
	"cookly/business/reviews"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mysqlReviewRepository struct {
	Conn *gorm.DB
}

func NewMysqlReviewRepository(conn *gorm.DB) reviews.Repository {
	return &mysqlReviewRepository{
		Conn: conn,
	}
}

func (repository *mysqlReviewRepository) Create(ctx context.Context, reviewDomain *reviews.Domain) (reviews.Domain, error) {
	create := FromDomain(*reviewDomain)

	result := repository.Conn.Create(&create)
	if result.Error != nil {
		return reviews.Domain{}, result.Error
	}

	err := repository.Conn.Preload(clause.Associations).First(&create, create.ID).Error
	if err != nil {
		return reviews.Domain{}, err
	}

	return create.ToDomain(), nil
}

func (repository *mysqlReviewRepository) GetReviewsByRecipeID(ctx context.Context, id int) ([]reviews.Domain, error) {
	reviewsByRecipeID := []Review{}

	result := repository.Conn.Where("recipe_id = ?", id).Find(&reviewsByRecipeID)
	if result.Error != nil {
		return []reviews.Domain{}, result.Error
	}

	reviews := []reviews.Domain{}
	for _, review := range reviewsByRecipeID {
		reviews = append(reviews, review.ToDomain())
	}

	return reviews, nil
}
