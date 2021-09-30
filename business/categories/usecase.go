package categories

import (
	"context"
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
	err := cr.catRepo.Create(ctx, categoryDomain)
	if err != nil {
		return err
	}

	return nil
}
