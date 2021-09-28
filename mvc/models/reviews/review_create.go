package reviews

type CreateReview struct {
	UserID       int    `json:"userId"`
	RecipeID     int    `json:"recipeId"`
	Rating       int    `json:"rating"`
	RecipeReview string `json:"review"`
}
