// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	reviews "cookly/business/reviews"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CountReviews provides a mock function with given fields: recipeId
func (_m *Repository) CountReviews(recipeId int) (int, error) {
	ret := _m.Called(recipeId)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(recipeId)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(recipeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, data
func (_m *Repository) Create(ctx context.Context, data *reviews.Domain) (reviews.Domain, error) {
	ret := _m.Called(ctx, data)

	var r0 reviews.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *reviews.Domain) reviews.Domain); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(reviews.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *reviews.Domain) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReviewsByRecipeID provides a mock function with given fields: ctx, recipeId
func (_m *Repository) GetReviewsByRecipeID(ctx context.Context, recipeId int) ([]reviews.Domain, error) {
	ret := _m.Called(ctx, recipeId)

	var r0 []reviews.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) []reviews.Domain); ok {
		r0 = rf(ctx, recipeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reviews.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, recipeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRecipeRating provides a mock function with given fields: recipeId, newRating
func (_m *Repository) UpdateRecipeRating(recipeId int, newRating float64) error {
	ret := _m.Called(recipeId, newRating)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, float64) error); ok {
		r0 = rf(recipeId, newRating)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
