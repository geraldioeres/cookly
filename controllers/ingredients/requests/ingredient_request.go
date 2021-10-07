package requests

import "cookly/business/ingredients"

type IngredientRequest struct {
	ProductID int    `json:"productId"`
	Amount    string `json:"amount"`
}

func (request *IngredientRequest) ToDomain() *ingredients.IngredientDomain {
	return &ingredients.IngredientDomain{
		ProductID: request.ProductID,
		Amount:    request.Amount,
	}
}
