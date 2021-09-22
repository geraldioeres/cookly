package models

import "time"

type CartProducts struct {
	ID            int
	UserCartID    int
	ProductID     int
	TotalPrice    int
	TotalQuantity int
	CreatedAt     time.Time
	ModifiedAt    time.Time
}
