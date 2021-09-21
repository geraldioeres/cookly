package models

import "time"

type Products struct {
	ID            int
	ProductTypeID int
	Name          string
	Description   string
	Price         int
	Status        string
	CreatedAt     time.Time
	ModifiedAt    time.Time
}
