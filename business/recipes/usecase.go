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

func (ruc *RecipeUseCase) Create(ctx context.Context, recipeDomain *Domain) (Domain, error) {
	result, err := ruc.recipeRepo.Create(ctx, recipeDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (ruc *RecipeUseCase) RecipeByID(ctx context.Context, id int) (Domain, error) {
	result, err := ruc.recipeRepo.RecipeByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (ruc *RecipeUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	result, err := ruc.recipeRepo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (ruc *RecipeUseCase) Update(ctx context.Context, recipeDomain *Domain) (*Domain, error) {
	_, err := ruc.recipeRepo.RecipeByID(ctx, recipeDomain.ID)
	if err != nil {
		return &Domain{}, err
	}

	result, err := ruc.recipeRepo.Update(ctx, recipeDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}
