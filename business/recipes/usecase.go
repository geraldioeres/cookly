package recipes

import (
	"context"
	"time"
)

type RecipeUseCase struct {
	recipeRepo     Repository
	contextTimeout time.Duration
}

func NewRecipeUseCase(rr Repository, timeout time.Duration) UseCase {
	return &RecipeUseCase{
		recipeRepo:     rr,
		contextTimeout: timeout,
	}
}

func (ruc *RecipeUseCase) Create(ctx context.Context, recipeDomain *Domain) error {
	err := ruc.recipeRepo.Create(ctx, recipeDomain)
	if err != nil {
		return err
	}

	return nil
}
