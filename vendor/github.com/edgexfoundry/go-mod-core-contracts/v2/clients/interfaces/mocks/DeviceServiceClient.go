// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"

	errors "github.com/edgexfoundry/go-mod-core-contracts/v2/errors"

	mock "github.com/stretchr/testify/mock"

	requests "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"

	responses "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/responses"
)

// DeviceServiceClient is an autogenerated mock type for the DeviceServiceClient type
type DeviceServiceClient struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, reqs
func (_m *DeviceServiceClient) Add(ctx context.Context, reqs []requests.AddDeviceServiceRequest) ([]common.BaseWithIdResponse, errors.EdgeX) {
	ret := _m.Called(ctx, reqs)

	var r0 []common.BaseWithIdResponse
	if rf, ok := ret.Get(0).(func(context.Context, []requests.AddDeviceServiceRequest) []common.BaseWithIdResponse); ok {
		r0 = rf(ctx, reqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.BaseWithIdResponse)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, []requests.AddDeviceServiceRequest) errors.EdgeX); ok {
		r1 = rf(ctx, reqs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllDeviceServices provides a mock function with given fields: ctx, labels, offset, limit
func (_m *DeviceServiceClient) AllDeviceServices(ctx context.Context, labels []string, offset int, limit int) (responses.MultiDeviceServicesResponse, errors.EdgeX) {
	ret := _m.Called(ctx, labels, offset, limit)

	var r0 responses.MultiDeviceServicesResponse
	if rf, ok := ret.Get(0).(func(context.Context, []string, int, int) responses.MultiDeviceServicesResponse); ok {
		r0 = rf(ctx, labels, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiDeviceServicesResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, []string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, labels, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeleteByName provides a mock function with given fields: ctx, name
func (_m *DeviceServiceClient) DeleteByName(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX) {
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

// DeviceServiceByName provides a mock function with given fields: ctx, name
func (_m *DeviceServiceClient) DeviceServiceByName(ctx context.Context, name string) (responses.DeviceServiceResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name)

	var r0 responses.DeviceServiceResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) responses.DeviceServiceResponse); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(responses.DeviceServiceResponse)
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

// Update provides a mock function with given fields: ctx, reqs
func (_m *DeviceServiceClient) Update(ctx context.Context, reqs []requests.UpdateDeviceServiceRequest) ([]common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, reqs)

	var r0 []common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, []requests.UpdateDeviceServiceRequest) []common.BaseResponse); ok {
		r0 = rf(ctx, reqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.BaseResponse)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, []requests.UpdateDeviceServiceRequest) errors.EdgeX); ok {
		r1 = rf(ctx, reqs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}
