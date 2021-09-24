package recipes

import (
	"cookly/models/users"
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	Id           int    `gorm:"primaryKey" json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	RecipeTypeId int
	RecipeType   RecipeType
	UserId       int
	User         users.User
	Rating       float64        `json:"rating"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt"`
}
