package recipes

import (
	"cookly/business/recipes"
	"cookly/controllers"
	"cookly/controllers/recipes/requests"
	"cookly/controllers/recipes/responses"
	"net/http"
	"strconv"

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

	if err := c.Bind(&createRecipe); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()

	result, err := recipeController.RecipeUseCase.Create(ctx, createRecipe.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(result))
}

func (recipeController *RecipeController) RecipeByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))
	recipe, err := recipeController.RecipeUseCase.RecipeByID(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadGateway, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(recipe))
}
