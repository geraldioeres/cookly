package recipecategory

import (
	"cookly/mvc/models/recipe_category"
	"cookly/mvc/configs"
	"cookly/mvc/models/responses"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateCategoryController(c echo.Context) error {
	var typeInput recipecategory.CreateType
	c.Bind(&typeInput)

	var recipe_cat recipecategory.RecipeCategory
	recipe_cat.Name = typeInput.Name

	result := configs.DB.Create(&recipe_cat)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create recipe category",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully created recipe category",
		Data:    recipe_cat,
	})
}

func GetAllCategories(c echo.Context) error {
	recipesCat := []recipecategory.RecipeCategory{}

	result := configs.DB.Find(&recipesCat)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error when retrieve recipe categories from database",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success get recipe categories data",
		Data:    recipesCat,
	})
}
