package reviews

import (
	"cookly/business/reviews"
	"cookly/controllers"
	"cookly/controllers/reviews/requests"
	"cookly/controllers/reviews/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReviewController struct {
	ReviewUseCase reviews.UseCase
}

func NewReviewController(reviewUseCase reviews.UseCase) *ReviewController {
	return &ReviewController{
		ReviewUseCase: reviewUseCase,
	}
}

func (reviewController *ReviewController) Create(c echo.Context) error {
	createReview := requests.Review{}

	if err := c.Bind(&createReview); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()

	result, err := reviewController.ReviewUseCase.Create(ctx, createReview.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(&result))
}

func (reviewController *ReviewController) GetReviewsByRecipeID(c echo.Context) error {
	reviews := []responses.ReviewResponse{}
	ctx := c.Request().Context()

	recipeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := reviewController.ReviewUseCase.GetReviewsByRecipeID(ctx, recipeId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	for _, review := range result {
		reviews = append(reviews, responses.FromDomain(&review))
	}

	return controllers.NewSuccessResponse(c, reviews)
}
