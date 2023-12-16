// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	service "github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/service"
)

// IdentityServiceClient is an autogenerated mock type for the IdentityServiceClient type
type IdentityServiceClient struct {
	mock.Mock
}

type IdentityServiceClient_UserInfo struct {
	*mock.Call
}

func (_m IdentityServiceClient_UserInfo) Return(_a0 *service.UserInfoResponse, _a1 error) *IdentityServiceClient_UserInfo {
	return &IdentityServiceClient_UserInfo{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *IdentityServiceClient) OnUserInfo(ctx context.Context, in *service.UserInfoRequest, opts ...grpc.CallOption) *IdentityServiceClient_UserInfo {
	c_call := _m.On("UserInfo", ctx, in, opts)
	return &IdentityServiceClient_UserInfo{Call: c_call}
}

func (_m *IdentityServiceClient) OnUserInfoMatch(matchers ...interface{}) *IdentityServiceClient_UserInfo {
	c_call := _m.On("UserInfo", matchers...)
	return &IdentityServiceClient_UserInfo{Call: c_call}
}

// UserInfo provides a mock function with given fields: ctx, in, opts
func (_m *IdentityServiceClient) UserInfo(ctx context.Context, in *service.UserInfoRequest, opts ...grpc.CallOption) (*service.UserInfoResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *service.UserInfoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *service.UserInfoRequest, ...grpc.CallOption) *service.UserInfoResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.UserInfoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.UserInfoRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
