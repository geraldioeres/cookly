package models

import "time"

type RecipeTypes struct {
	ID         int
	Name       string
	CreatedAt  time.Time
	ModifiedAt time.Time
}
