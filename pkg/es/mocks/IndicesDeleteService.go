// Copyright (c) The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
//
// Run 'make generate-mocks' to regenerate.

// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	elastic "github.com/olivere/elastic"

	mock "github.com/stretchr/testify/mock"
)

// IndicesDeleteService is an autogenerated mock type for the IndicesDeleteService type
type IndicesDeleteService struct {
	mock.Mock
}

// Do provides a mock function with given fields: ctx
func (_m *IndicesDeleteService) Do(ctx context.Context) (*elastic.IndicesDeleteResponse, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Do")
	}

	var r0 *elastic.IndicesDeleteResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*elastic.IndicesDeleteResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *elastic.IndicesDeleteResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastic.IndicesDeleteResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIndicesDeleteService creates a new instance of IndicesDeleteService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIndicesDeleteService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IndicesDeleteService {
	mock := &IndicesDeleteService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}