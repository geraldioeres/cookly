package categories

import (
	"cookly/business/categories"
	"cookly/controllers"
	"cookly/controllers/categories/requests"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	CatUseCase categories.UseCase
}

func NewCategoryController(catUseCase categories.UseCase) *CategoryController {
	return &CategoryController{
		CatUseCase: catUseCase,
	}
}

func (categoryController *CategoryController) Create(c echo.Context) error {
	createCategory := requests.Category{}
	c.Bind(&createCategory)

	ctx := c.Request().Context()

	err := categoryController.CatUseCase.Create(ctx, createCategory.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Successfully create category")
}
