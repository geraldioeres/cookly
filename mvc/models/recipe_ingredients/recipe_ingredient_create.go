package recipeingredients

type CreateRecipeIngredient struct {
	PoductID int    `json:"ProductId"`
	Amount   string `json:"amount"`
}
