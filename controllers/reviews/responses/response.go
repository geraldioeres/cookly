package responses

import "cookly/business/reviews"

type ReviewResponse struct {
	ID       int     `json:"id"`
	UserID   int     `json:"userId"`
	RecipeID int     `json:"recipeId"`
	Rating   float64 `json:"rating"`
	Comment  string  `json:"comment"`
}

func FromDomain(domain *reviews.Domain) ReviewResponse {
	return ReviewResponse{
		ID:       domain.ID,
		UserID:   domain.UserID,
		RecipeID: domain.RecipeID,
		Rating:   domain.Rating,
		Comment:  domain.Comment,
	}
}
