package requests

import "cookly/business/reviews"

type Review struct {
	UserID   int     `json:"userId"`
	RecipeID int     `json:"recipeId"`
	Rating   float64 `json:"rating"`
	Comment  string  `json:"comment"`
}

func (request *Review) ToDomain() *reviews.Domain {
	return &reviews.Domain{
		UserID:   request.UserID,
		RecipeID: request.RecipeID,
		Rating:   request.Rating,
		Comment:  request.Comment,
	}
}
