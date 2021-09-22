package models

import "time"

type ProductTypes struct {
	ID         int
	Name       string
	CreatedAt  time.Time
	ModifiedAt time.Time
}
