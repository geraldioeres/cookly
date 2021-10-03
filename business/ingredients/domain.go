package ingredients

import (
	"cookly/business/products"
	"time"
)

type IngredientDomain struct {
	ID        int
	RecipeID  int
	ProductID int
	Product   products.Domain
	Amount    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
