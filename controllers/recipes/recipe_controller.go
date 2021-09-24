package recipes

import (
	"cookly/configs"
	"cookly/models/recipes"
	"cookly/models/responses"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetRecipesController(c echo.Context) error {
	recipes := []recipes.Recipe{}

	result := configs.DB.Find(&recipes)
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
	var createRecipe recipes.RecipeRegister
	c.Bind(&createRecipe)

	var recipeDB recipes.Recipe
	recipeDB.Title = createRecipe.Title
	recipeDB.Description = createRecipe.Description
	recipeDB.CreatedAt = time.Now()
	result := configs.DB.Create(&recipeDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create recipe",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully created recipe",
		Data:    recipeDB,
	})
}
