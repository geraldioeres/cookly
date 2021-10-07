package recipeingredients

import "cookly/mvc/models/products"

type RecipeIngredient struct {
	ID       int `gorm:"primaryKey" json:"id"`
	PoductID int `json:"ProductId"`
	Product  products.Product
	Amount   string `json:"amount"`
}
