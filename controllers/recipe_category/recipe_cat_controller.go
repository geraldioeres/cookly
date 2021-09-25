package recipecategory

import (
	"cookly/configs"

	recipecategory "cookly/models/recipe_category"
	"cookly/models/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateCategoryController(c echo.Context) error {
	var typeInput recipecategory.CreateType
	c.Bind(&typeInput)

	var recipe_types recipecategory.RecipeCategory
	recipe_types.Name = typeInput.Name

	result := configs.DB.Create(&recipe_types)
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
		Data:    recipe_types,
	})
}
