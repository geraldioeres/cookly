package mysql

import (
	"fmt"
	"log"

	"cookly/drivers/databases/categories"
	"cookly/drivers/databases/ingredients"
	"cookly/drivers/databases/products"
	"cookly/drivers/databases/recipes"
	"cookly/drivers/databases/reviews"
	"cookly/drivers/databases/steps"
	"cookly/drivers/databases/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Database string
}

func (config *ConfigDB) InitialDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&users.Users{},
		&categories.Category{},
		&products.Product{},
		&recipes.Recipe{},
		&steps.Step{},
		&ingredients.Ingredient{},
		&reviews.Review{},
	)

	return db
}
