package recipes

import (
	"cookly/business/recipes"
	"cookly/controllers"
	"cookly/controllers/recipes/requests"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RecipeController struct {
	RecipeUseCase recipes.UseCase
}

func NewRecipeController(recipeUseCase recipes.UseCase) *RecipeController {
	return &RecipeController{
		RecipeUseCase: recipeUseCase,
	}
}

func (recipeController *RecipeController) Create(c echo.Context) error {
	createRecipe := requests.Recipe{}
	c.Bind(&createRecipe)

	ctx := c.Request().Context()

	err := recipeController.RecipeUseCase.Create(ctx, createRecipe.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Successfully created recipe")
}