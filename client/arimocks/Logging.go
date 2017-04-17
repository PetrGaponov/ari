package arimocks

import ari "github.com/CyCoreSystems/ari"
import mock "github.com/stretchr/testify/mock"

// Logging is an autogenerated mock type for the Logging type
type Logging struct {
	mock.Mock
}

// Create provides a mock function with given fields: name, level
func (_m *Logging) Create(name string, level string) error {
	ret := _m.Called(name, level)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(name, level)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Data provides a mock function with given fields: name
func (_m *Logging) Data(name string) (*ari.LogData, error) {
	ret := _m.Called(name)

	var r0 *ari.LogData
	if rf, ok := ret.Get(0).(func(string) *ari.LogData); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.LogData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: name
func (_m *Logging) Delete(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: name
func (_m *Logging) Get(name string) ari.LogHandle {
	ret := _m.Called(name)

	var r0 ari.LogHandle
	if rf, ok := ret.Get(0).(func(string) ari.LogHandle); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.LogHandle)
		}
	}

	return r0
}

// List provides a mock function with given fields:
func (_m *Logging) List() ([]ari.LogHandle, error) {
	ret := _m.Called()

	var r0 []ari.LogHandle
	if rf, ok := ret.Get(0).(func() []ari.LogHandle); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ari.LogHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Rotate provides a mock function with given fields: name
func (_m *Logging) Rotate(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}