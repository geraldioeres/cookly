package responses

import "cookly/business/users"

type LoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func LoginFromDomain(domain users.Domain) LoginResponse {
	return LoginResponse{
		ID:    domain.ID,
		Name:  domain.Name,
		Email: domain.Email,
		Token: domain.Token,
	}
}
