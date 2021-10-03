package reviews

import (
	"context"
	"time"
)

type ReviewUseCase struct {
	reviewRepo     Repository
	contextTimeout time.Duration
}

func NewReviewUseCase(rvr Repository, timeout time.Duration) UseCase {
	return &ReviewUseCase{
		reviewRepo:     rvr,
		contextTimeout: timeout,
	}
}

func (rvuc *ReviewUseCase) Create(ctx context.Context, reviewDomain *Domain) (Domain, error) {
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
