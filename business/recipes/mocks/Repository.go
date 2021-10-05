// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	recipes "cookly/business/recipes"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, data
func (_m *Repository) Create(ctx context.Context, data *recipes.Domain) (recipes.Domain, error) {
	ret := _m.Called(ctx, data)

	var r0 recipes.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *recipes.Domain) recipes.Domain); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(recipes.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *recipes.Domain) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repository) GetAll(ctx context.Context) ([]recipes.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []recipes.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []recipes.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]recipes.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RecipeByID provides a mock function with given fields: ctx, id
func (_m *Repository) RecipeByID(ctx context.Context, id int) (recipes.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 recipes.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) recipes.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(recipes.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, data
func (_m *Repository) Update(ctx context.Context, data *recipes.Domain) (recipes.Domain, error) {
	ret := _m.Called(ctx, data)

	var r0 recipes.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *recipes.Domain) recipes.Domain); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(recipes.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *recipes.Domain) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}