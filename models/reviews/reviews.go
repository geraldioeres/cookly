package reviews

type Review struct {
	ID           int    `json:"id"`
	UserID       int    `json:"userId"`
	RecipeID     int    `json:"recipeId"`
	Rating       int    `json:"rating"`
	RecipeReview string `json:"review"`
}
