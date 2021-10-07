package steps

import (
	"time"
)

type StepDomain struct {
	ID          int
	RecipeID    int
	Order       int
	Instruction string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
