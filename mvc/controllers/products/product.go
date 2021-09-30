package products

import (
	"cookly/mvc/configs"
	"cookly/mvc/models/products"
	"cookly/mvc/models/responses"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateProduct(c echo.Context) error {
	var productInput products.CreateProduct
	c.Bind(&productInput)

	var productDB products.Product
	productDB.Name = productInput.Name

	result := configs.DB.Create(&productDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create product",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully created prodcut",
		Data:    productDB,
	})
}

func GetAllProducts(c echo.Context) error {
	products := []products.Product{}

	result := configs.DB.Find(&products)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error when retrieve products from database",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success get products data",
		Data:    products,
	})
}
