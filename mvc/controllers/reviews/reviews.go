package reviews

import (
	"cookly/mvc/configs"
	"cookly/mvc/models/responses"
	"cookly/mvc/models/reviews"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateReview(c echo.Context) error {
	var reviewInput reviews.CreateReview
	c.Bind(&reviewInput)

	var reviewDB reviews.Review
	reviewDB.UserID = reviewInput.UserID
	reviewDB.RecipeID = reviewInput.RecipeID
	reviewDB.Rating = reviewInput.Rating
	reviewDB.RecipeReview = reviewInput.RecipeReview

	result := configs.DB.Create(&reviewDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create review",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully created review",
		Data:    reviewDB,
	})
}

func GetReviewByRecipeID(c echo.Context) error {
	review := []reviews.Review{}

	id, _ := strconv.Atoi(c.Param("id"))
	result := configs.DB.Where("recipe_id = ?", id).Find(&review)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get reviews data",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success get reviews data",
		Data:    review,
	})
}
