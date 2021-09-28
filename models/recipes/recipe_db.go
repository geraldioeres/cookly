package recipes

import (
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	ID               int            `gorm:"primaryKey" json:"id"`
	Title            string         `json:"title"`
	Description      string         `json:"description"`
	Rating           int            `json:"rating"`
	UserID           int            `json:"userId"`
	RecipeCategoryID int            `json:"recipeCategoryId"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `json:"deletedAt"`
}
