package reviews

import (
	"context"
	"cookly/business/recipes"
	"fmt"
	"time"
)

type ReviewUseCase struct {
	reviewRepo     Repository
	recipeRepo     recipes.Repository
	contextTimeout time.Duration
}

func NewReviewUseCase(rvr Repository, rcr recipes.Repository, timeout time.Duration) UseCase {
	return &ReviewUseCase{
		reviewRepo:     rvr,
		recipeRepo:     rcr,
		contextTimeout: timeout,
	}
}

func (rvuc *ReviewUseCase) Create(ctx context.Context, reviewDomain *Domain) (Domain, error) {
	// Get the number of reviews based on the recipe id
	numberOfReviews, err := rvuc.reviewRepo.CountReviews(reviewDomain.RecipeID)
	if err != nil {
		return Domain{}, nil
	}
	fmt.Println("number is ", numberOfReviews)

	// Get current rating value of the recipe based on recipe id
	recipes, err := rvuc.recipeRepo.RecipeByID(ctx, reviewDomain.RecipeID)
	if err != nil {
		return Domain{}, nil
	}

	// Rating calculation
	sum := (float64(numberOfReviews) * recipes.Rating) + reviewDomain.Rating
	newRating := sum / (float64(numberOfReviews) + 1)
	fmt.Println("sum is", sum)
	fmt.Println("new rating is", newRating)

	err = rvuc.reviewRepo.UpdateRecipeRating(reviewDomain.RecipeID, newRating)
	if err != nil {
		return Domain{}, err
	}

	result, err := rvuc.reviewRepo.Create(ctx, reviewDomain)
	if err != nil {
		return Domain{}, nil
	}

	return result, nil
}

func (rvuc *ReviewUseCase) GetReviewsByRecipeID(ctx context.Context, recipeId int) ([]Domain, error) {
	result, err := rvuc.reviewRepo.GetReviewsByRecipeID(ctx, recipeId)
	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}
