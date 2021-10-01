package categories

import (
	"cookly/business/categories"
	"cookly/controllers"
	"cookly/controllers/categories/requests"
	"cookly/controllers/categories/responses"
	"net/http"
	"strconv"

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

func (categoryController *CategoryController) GetAll(c echo.Context) error {
	categories := []responses.CategoryResponse{}
	ctx := c.Request().Context()

	result, err := categoryController.CatUseCase.GetAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	for _, res := range result {
		categories = append(categories, responses.FromDomain(res))
	}

	return controllers.NewSuccessResponse(c, categories)
}

func (categoryController *CategoryController) Update(c echo.Context) error {
	update := requests.Category{}
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	c.Bind(&update)
	err := categoryController.CatUseCase.Update(ctx, update.ToDomain(), id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, "Successfully updated category")
}
