package responses

import (
	"cookly/business/ingredients"
	"cookly/controllers/products/responses"
)

type IngredientResponse struct {
	ID        int    `json:"id"`
	ProductID int    `json:"productId"`
	Product   responses.ProductResponse `json:"product"`
	Amount    string `json:"amount"`
}

func FromIngredientDomain(domain ingredients.IngredientDomain) IngredientResponse {
	return IngredientResponse{
		ID:        domain.ID,
		ProductID: domain.ProductID,
		Product: responses.ProductResponse(domain.Product),
		Amount:    domain.Amount,
	}
}
