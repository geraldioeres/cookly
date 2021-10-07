package products

import (
	"cookly/business/products"
	"cookly/controllers"
	"cookly/controllers/products/requests"
	"cookly/controllers/products/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	ProductUseCase products.UseCase
}

func NewProductController(productUseCase products.UseCase) *ProductController {
	return &ProductController{
		ProductUseCase: productUseCase,
	}
}

func (productController *ProductController) Create(c echo.Context) error {
	createProduct := requests.Product{}
	c.Bind(&createProduct)

	ctx := c.Request().Context()

	err := productController.ProductUseCase.Create(ctx, createProduct.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Successfully create product")
}

func (productController *ProductController) GetAll(c echo.Context) error {
	products := []responses.ProductResponse{}
	ctx := c.Request().Context()

	result, err := productController.ProductUseCase.GetAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadGateway, err)
	}

	for _, res := range result {
		products = append(products, responses.FromDomain(res))
	}

	return controllers.NewSuccessResponse(c, products)
}

func (productController *ProductController) Update(c echo.Context) error {
	update := requests.Product{}
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	err := c.Bind(&update)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err = productController.ProductUseCase.Update(ctx, update.ToDomain(), id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, "Successfully updated product")
}
