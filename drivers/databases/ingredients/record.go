package ingredients

import (
	"cookly/business/ingredients"
	"cookly/drivers/databases/products"
	"time"
)

type Ingredient struct {
	ID        int `gorm:"primaryKey" json:"id"`
	RecipeID  int
	ProductID int
	Product   products.Product
	Amount    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (record *Ingredient) ToDomain() ingredients.IngredientDomain {
	return ingredients.IngredientDomain{
		ID:        record.ID,
		RecipeID:  record.RecipeID,
		ProductID: record.ProductID,
		Product:   record.Product.ToDomain(),
		Amount:    record.Amount,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

func FromDomain(domain *ingredients.IngredientDomain) *Ingredient {
	return &Ingredient{
		ID:        domain.ID,
		RecipeID:  domain.RecipeID,
		ProductID: domain.ProductID,
		Amount:    domain.Amount,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
