package steps

type Step struct {
	ID          int
	RecipeID    int
	Order       int `json:"order"`
	Instruction string `json:"instruction"`
}
