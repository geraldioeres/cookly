package requests

import "cookly/business/users"

type User struct {
	Name     string
	Email    string
	Password string
}

func (request *User) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}
