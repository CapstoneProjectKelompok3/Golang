// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	vehicles "project-capston/features/vehicles"

	mock "github.com/stretchr/testify/mock"
)

// VehicleData is an autogenerated mock type for the VehicleDataInterface type
type VehicleData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *VehicleData) Delete(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: input
func (_m *VehicleData) Insert(input vehicles.VehicleEntity) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(vehicles.VehicleEntity) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectAll provides a mock function with given fields:
func (_m *VehicleData) SelectAll() ([]vehicles.VehicleEntity, error) {
	ret := _m.Called()

	var r0 []vehicles.VehicleEntity
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]vehicles.VehicleEntity, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []vehicles.VehicleEntity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]vehicles.VehicleEntity)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectById provides a mock function with given fields: id
func (_m *VehicleData) SelectById(id uint) (vehicles.VehicleEntity, error) {
	ret := _m.Called(id)

	var r0 vehicles.VehicleEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (vehicles.VehicleEntity, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) vehicles.VehicleEntity); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(vehicles.VehicleEntity)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: input, id
func (_m *VehicleData) Update(input vehicles.VehicleEntity, id uint) error {
	ret := _m.Called(input, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(vehicles.VehicleEntity, uint) error); ok {
		r0 = rf(input, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewVehicleData interface {
	mock.TestingT
	Cleanup(func())
}

// NewVehicleData creates a new instance of VehicleData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewVehicleData(t mockConstructorTestingTNewVehicleData) *VehicleData {
	mock := &VehicleData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
