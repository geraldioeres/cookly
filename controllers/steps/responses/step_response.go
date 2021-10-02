package responses

import (
	"cookly/business/steps"
)

type StepResponse struct {
	Order       int    `json:"order"`
	Instruction string `json:"instruction"`
}

func FromStepDomain(domain steps.StepDomain) StepResponse {
	return StepResponse{
		Order:       domain.Order,
		Instruction: domain.Instruction,
	}
}
