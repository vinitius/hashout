// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	checkout "viniti.us/hashout/models/checkout"

	mock "github.com/stretchr/testify/mock"
)

// DiscountClient is an autogenerated mock type for the Client type
type DiscountClient struct {
	mock.Mock
}

// GetDiscount provides a mock function with given fields: item
func (_m *DiscountClient) GetDiscount(item *checkout.Item) error {
	ret := _m.Called(item)

	var r0 error
	if rf, ok := ret.Get(0).(func(*checkout.Item) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}