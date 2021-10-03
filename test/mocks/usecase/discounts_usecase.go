// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	checkout "viniti.us/hashout/models/checkout"

	mock "github.com/stretchr/testify/mock"
)

// DiscountUseCase is an autogenerated mock type for the Service type
type DiscountUseCase struct {
	mock.Mock
}

// CalculateDiscounts provides a mock function with given fields: items
func (_m *DiscountUseCase) CalculateDiscounts(items []checkout.Item) ([]checkout.Item, error) {
	ret := _m.Called(items)

	var r0 []checkout.Item
	if rf, ok := ret.Get(0).(func([]checkout.Item) []checkout.Item); ok {
		r0 = rf(items)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]checkout.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]checkout.Item) error); ok {
		r1 = rf(items)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
