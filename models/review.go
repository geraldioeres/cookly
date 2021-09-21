package models

import "time"

type Reviews struct {
	ID          int
	UserID      int
	RecipeID    int
	Rating      int
	Description string
	CreatedAt   time.Time
	ModifiedAt  time.Time
}
