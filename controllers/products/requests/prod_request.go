package requests

import "cookly/business/products"

type Product struct {
	Name string `json:"name"`
}

func (request *Product) ToDomain() *products.Domain {
	return &products.Domain{
		Name: request.Name,
	}
}
