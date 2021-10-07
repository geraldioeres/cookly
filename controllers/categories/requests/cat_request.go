package requests

import "cookly/business/categories"

type Category struct {
	Name string
}

func (request *Category) ToDomain() *categories.Domain {
	return &categories.Domain{
		Name: request.Name,
	}
}