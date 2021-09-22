package models

import "time"

type UserCart struct {
	ID              int
	UserID          int
	PaymentMethodID int
	CreatedAt       time.Time
	ModifiedAt      time.Time
}
