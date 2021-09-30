package recipecategory

import (
	"time"

	"gorm.io/gorm"
)

type RecipeCategory struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
