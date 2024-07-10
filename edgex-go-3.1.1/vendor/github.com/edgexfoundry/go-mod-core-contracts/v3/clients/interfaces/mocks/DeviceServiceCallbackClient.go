// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"

	errors "github.com/edgexfoundry/go-mod-core-contracts/v3/errors"

	mock "github.com/stretchr/testify/mock"

	requests "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/requests"
)

// DeviceServiceCallbackClient is an autogenerated mock type for the DeviceServiceCallbackClient type
type DeviceServiceCallbackClient struct {
	mock.Mock
}

// AddDeviceCallback provides a mock function with given fields: ctx, request
func (_m *DeviceServiceCallbackClient) AddDeviceCallback(ctx context.Context, request requests.AddDeviceRequest) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, request)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, requests.AddDeviceRequest) common.BaseResponse); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, requests.AddDeviceRequest) errors.EdgeX); ok {
		r1 = rf(ctx, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AddProvisionWatcherCallback provides a mock function with given fields: ctx, request
func (_m *DeviceServiceCallbackClient) AddProvisionWatcherCallback(ctx context.Context, request requests.AddProvisionWatcherRequest) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, request)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, requests.AddProvisionWatcherRequest) common.BaseResponse); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, requests.AddProvisionWatcherRequest) errors.EdgeX); ok {
		r1 = rf(ctx, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeleteDeviceCallback provides a mock function with given fields: ctx, name
func (_m *DeviceServiceCallbackClient) DeleteDeviceCallback(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) common.BaseResponse); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeleteProvisionWatcherCallback provides a mock function with given fields: ctx, name
func (_m *DeviceServiceCallbackClient) DeleteProvisionWatcherCallback(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) common.BaseResponse); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// UpdateDeviceCallback provides a mock function with given fields: ctx, request
func (_m *DeviceServiceCallbackClient) UpdateDeviceCallback(ctx context.Context, request requests.UpdateDeviceRequest) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, request)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, requests.UpdateDeviceRequest) common.BaseResponse); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, requests.UpdateDeviceRequest) errors.EdgeX); ok {
		r1 = rf(ctx, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// UpdateDeviceProfileCallback provides a mock function with given fields: ctx, request
func (_m *DeviceServiceCallbackClient) UpdateDeviceProfileCallback(ctx context.Context, request requests.DeviceProfileRequest) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, request)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, requests.DeviceProfileRequest) common.BaseResponse); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, requests.DeviceProfileRequest) errors.EdgeX); ok {
		r1 = rf(ctx, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// UpdateDeviceServiceCallback provides a mock function with given fields: ctx, request
func (_m *DeviceServiceCallbackClient) UpdateDeviceServiceCallback(ctx context.Context, request requests.UpdateDeviceServiceRequest) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, request)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, requests.UpdateDeviceServiceRequest) common.BaseResponse); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, requests.UpdateDeviceServiceRequest) errors.EdgeX); ok {
		r1 = rf(ctx, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// UpdateProvisionWatcherCallback provides a mock function with given fields: ctx, request
func (_m *DeviceServiceCallbackClient) UpdateProvisionWatcherCallback(ctx context.Context, request requests.UpdateProvisionWatcherRequest) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, request)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, requests.UpdateProvisionWatcherRequest) common.BaseResponse); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, requests.UpdateProvisionWatcherRequest) errors.EdgeX); ok {
		r1 = rf(ctx, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ValidateDeviceCallback provides a mock function with given fields: ctx, request
func (_m *DeviceServiceCallbackClient) ValidateDeviceCallback(ctx context.Context, request requests.AddDeviceRequest) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, request)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, requests.AddDeviceRequest) common.BaseResponse); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, requests.AddDeviceRequest) errors.EdgeX); ok {
		r1 = rf(ctx, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewDeviceServiceCallbackClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewDeviceServiceCallbackClient creates a new instance of DeviceServiceCallbackClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDeviceServiceCallbackClient(t mockConstructorTestingTNewDeviceServiceCallbackClient) *DeviceServiceCallbackClient {
	mock := &DeviceServiceCallbackClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
