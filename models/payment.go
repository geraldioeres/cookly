package models

import "time"

type PaymentMethods struct {
	ID         int
	Name       string
	CreatedAt  time.Time
	ModifiedAt time.Time
}
