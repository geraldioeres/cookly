package recipes

import (
	"cookly/configs"
	"cookly/models/recipes"
	"cookly/models/responses"
	"net/http"

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
