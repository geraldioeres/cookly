package recipes

type CreateRecipe struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	UserID           int    `json:"userId"`
	RecipeCategoryID int    `json:"recipeCategoryId"`
	//RecipeIngredient []recipeingredients.RecipeIngredient `json:"ingredients"`
}
