package requests

import "cookly/business/steps"

type StepRequest struct {
	Order       int    `json:"order"`
	Instruction string `json:"instruction"`
}

func (request *StepRequest) ToDomain() *steps.StepDomain {
	return &steps.StepDomain{
		Order:       request.Order,
		Instruction: request.Instruction,
	}
}
