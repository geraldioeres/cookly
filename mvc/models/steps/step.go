package steps

import "time"

type Step struct {
	ID          int
	RecipeID    int
	Order       int `json:"order"`
	Instruction string `json:"instruction"`
	CreatedAt time.Time 
	UpdatedAt time.Time
}
