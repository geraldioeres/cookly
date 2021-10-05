package categories

import (
	"context"
	"errors"
	"time"
)

type CategoryUseCase struct {
	catRepo        Repository
	contextTimeout time.Duration
}

func NewCategoryUseCase(cr Repository, timeout time.Duration) UseCase {
	return &CategoryUseCase{
		catRepo:        cr,
		contextTimeout: timeout,
	}
}

func (cr *CategoryUseCase) Create(ctx context.Context, categoryDomain *Domain) error {
	if categoryDomain.Name == "" {
		return errors.New("category name is empty")
	}

	err := cr.catRepo.Create(ctx, categoryDomain)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CategoryUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	result, err := cr.catRepo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (cr *CategoryUseCase) Update(ctx context.Context, categoryDomain *Domain, id int) error {
	err := cr.catRepo.Update(ctx, categoryDomain, id)
	if err != nil {
		return err
	}

	return nil
}
