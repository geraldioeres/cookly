package steps

import (
	"cookly/business/steps"
	"time"
)

type Step struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	RecipeID    int       
	Order       int       
	Instruction string    
	CreatedAt   time.Time 
	UpdatedAt   time.Time 
}

func (record *Step) ToDomain() steps.StepDomain {
	return steps.StepDomain{
		ID:          record.ID,
		RecipeID:    record.RecipeID,
		Order:       record.Order,
		Instruction: record.Instruction,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
	}
}

func FromDomain(domain *steps.StepDomain) *Step {
	return &Step{
		ID:          domain.ID,
		RecipeID:    domain.RecipeID,
		Order:       domain.Order,
		Instruction: domain.Instruction,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
