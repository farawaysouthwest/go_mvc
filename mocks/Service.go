// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	model "go_micro/model"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GetUser provides a mock function with given fields: id
func (_m *Service) GetUser(id int64) (model.User, error) {
	ret := _m.Called(id)

	var r0 model.User
	if rf, ok := ret.Get(0).(func(int64) model.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
