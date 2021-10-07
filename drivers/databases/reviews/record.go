package reviews

import (
	"cookly/business/reviews"
	"cookly/drivers/databases/recipes"
	"cookly/drivers/databases/users"
	"time"
)

type Review struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	User      users.Users
	RecipeID  int
	Recipe    recipes.Recipe
	Rating    float64
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (record *Review) ToDomain() reviews.Domain {
	return reviews.Domain{
		ID:        record.ID,
		UserID:    record.UserID,
		RecipeID:  record.RecipeID,
		Rating:    record.Rating,
		Comment:   record.Comment,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

func FromDomain(domain reviews.Domain) Review {
	return Review{
		ID:        domain.ID,
		UserID:    domain.UserID,
		RecipeID:  domain.RecipeID,
		Rating:    domain.Rating,
		Comment:   domain.Comment,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
