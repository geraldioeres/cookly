package responses

import (
	"cookly/business/ingredients"
	"cookly/business/products"
)

type IngredientResponse struct {
	Product products.Domain `json:"product"`
	Amount  string          `json:"amount"`
}

func FromIngredientDomain(domain ingredients.IngredientDomain) IngredientResponse {
	return IngredientResponse{
		Product: domain.Product,
		Amount:  domain.Amount,
	}
}
