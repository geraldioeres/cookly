package reviews

type CreateReview struct {
	UserID       int     `json:"userId"`
	RecipeID     int     `json:"recipeId"`
	Rating       float64 `json:"rating"`
	RecipeReview string  `json:"review"`
}
