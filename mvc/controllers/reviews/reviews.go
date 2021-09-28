package reviews

import (
	"cookly/configs"
	"cookly/models/responses"
	"cookly/models/reviews"
	"net/http"

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
