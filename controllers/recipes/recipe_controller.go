package recipes

import (
	"cookly/configs"
	"cookly/models/recipes"
	"cookly/models/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetRecipeByID(c echo.Context) error {
	var recipeById recipes.Recipe

	id, _ := strconv.Atoi(c.Param("id"))
	result := configs.DB.First(&recipeById, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get recipe data",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success get recipe data",
		Data:    recipeById,
	})
}

func GetRecipesController(c echo.Context) error {
	recipes := []recipes.Recipe{}

	result := configs.DB.Preload("RecipeCategory").Preload("User").Find(&recipes)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error when retrieve recipe from database",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success get recipe data",
		Data:    recipes,
	})
}

func CreateRecipeController(c echo.Context) error {
	var createRecipe recipes.CreateRecipe
	c.Bind(&createRecipe)

	var recipes recipes.Recipe
	recipes.Title = createRecipe.Title
	recipes.Description = createRecipe.Description
	recipes.UserID = createRecipe.UserID
	recipes.RecipeCategoryID = createRecipe.RecipeCategoryID
	result := configs.DB.Create(&recipes)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create recipe",
			Data:    nil,
		})
	}

	join := configs.DB.Preload("RecipeCategory").Preload("User").Find(&recipes)
	if join.Error != nil {
		return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create recipe",
			Data:    nil,
		})
	}
	
	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully created recipe",
		Data:    recipes,
	})
}
