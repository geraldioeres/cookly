package recipes

import (
	"time"

	"gorm.io/gorm"
)

type RecipeType struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
