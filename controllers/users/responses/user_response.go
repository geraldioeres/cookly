package responses

import (
	"cookly/business/users"
	"time"
)

type UserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		ID:       domain.ID,
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
	}
}