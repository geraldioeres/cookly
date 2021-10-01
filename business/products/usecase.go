package products

import (
	"context"
	"time"
)

type ProductUseCase struct {
	productRepo    Repository
	contextTimeout time.Duration
}

func NewProductUseCase(pr Repository, timeout time.Duration) UseCase {
	return &ProductUseCase{
		productRepo:    pr,
		contextTimeout: timeout,
	}
}

func (cr *ProductUseCase) Create(ctx context.Context, categoryDomain *Domain) error {
	err := cr.productRepo.Create(ctx, categoryDomain)
	if err != nil {
		return err
	}

	return nil
}

func (cr *ProductUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	result, err := cr.productRepo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return result, nil
}

func (cr *ProductUseCase) Update(ctx context.Context, categoryDomain *Domain, id int) error {
	err := cr.productRepo.Update(ctx, categoryDomain, id)
	if err != nil {
		return err
	}

	return nil
}
