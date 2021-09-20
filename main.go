package main

import "time"

type Users struct {
	ID         int
	Name       string
	Email      string
	Password   string
	Created_At time.Time
	ModifiedAt time.Time
}

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

type RecipeTypes struct {
	ID         int
	Name       string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type Collections struct {
	UserID     int
	RecipeID   int
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type Reviews struct {
	ID int
	UserID int
	RecipeID int
	Rating int
	Description string
	CreatedAt time.Time
	ModifiedAt time.Time
}

func main() {

}
