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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	recipe, err := recipeController.RecipeUseCase.RecipeByID(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(recipe))
}

func (recipeController *RecipeController) GetAll(c echo.Context) error {
	recipes := []responses.RecipeResponse{}
	ctx := c.Request().Context()

	result, err := recipeController.RecipeUseCase.GetAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	for _, recipe := range result {
		recipes = append(recipes, responses.FromDomain(recipe))
	}

	return controllers.NewSuccessResponse(c, recipes)
}

func (recipeController *RecipeController) Update(c echo.Context) error {
	recipeReq := requests.Recipe{}	
	ctx := c.Request().Context()
	Id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&recipeReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	recipeDomain := recipeReq.ToDomain()
	recipeDomain.ID = Id

	_, err := recipeController.RecipeUseCase.Update(ctx, recipeDomain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, "Succefully updated recipe")
}