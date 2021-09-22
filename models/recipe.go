package models

import "time"

type Recipes struct {
	ID           int
	UserID       int
	RecipeTypeID int
	Name         string
	Blurb        string
	Rating       float32
	Ingredients  []string
	CreatedAt    time.Time
	ModifiedAt   time.Time
}
