package business

import "errors"

var (
	ErrorInvalidPassword = errors.New("password cannot be null or empty")
	ErrorInvalidEmail = errors.New("email cannot be null or empty")
	ErrorDataNotFound = errors.New("data not found")
	ErrorInvalidRecipeID = errors.New("invalid recipe id")
	ErrorInvalidCategoryID = errors.New("invalid category id")
	ErrorCreateData = errors.New("error create data")
	ErrorEmptyName = errors.New("name cannot be empty")
)