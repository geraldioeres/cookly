package configs

import (
	recipecategory "cookly/models/recipe_category"
	"cookly/models/recipes"
	"cookly/models/users"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:diosql@tcp(127.0.0.1:3306)/cookly?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database" + err.Error())
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&users.User{}, &recipes.Recipe{}, &recipecategory.RecipeCategory{})
}
