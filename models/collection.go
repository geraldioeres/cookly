package models

import "time"

type Collections struct {
	UserID     int
	RecipeID   int
	CreatedAt  time.Time
	ModifiedAt time.Time
}
