// Code generated by mockery v2.10.4. DO NOT EDIT.

// Copyright (c) 2022 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	es "github.com/jaegertracing/jaeger/pkg/es"
)

// IndexService is an autogenerated mock type for the IndexService type
type IndexService struct {
	mock.Mock
}

// Add provides a mock function with given fields:
func (_m *IndexService) Add() {
	_m.Called()
}

// BodyJson provides a mock function with given fields: body
func (_m *IndexService) BodyJson(body interface{}) es.IndexService {
	ret := _m.Called(body)

	var r0 es.IndexService
	if rf, ok := ret.Get(0).(func(interface{}) es.IndexService); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(es.IndexService)
		}
	}

	return r0
}

// Id provides a mock function with given fields: id
func (_m *IndexService) Id(id string) es.IndexService {
	ret := _m.Called(id)

	var r0 es.IndexService
	if rf, ok := ret.Get(0).(func(string) es.IndexService); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(es.IndexService)
		}
	}

	return r0
}

// Index provides a mock function with given fields: index
func (_m *IndexService) Index(index string) es.IndexService {
	ret := _m.Called(index)

	var r0 es.IndexService
	if rf, ok := ret.Get(0).(func(string) es.IndexService); ok {
		r0 = rf(index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(es.IndexService)
		}
	}

	return r0
}

// Type provides a mock function with given fields: typ
func (_m *IndexService) Type(typ string) es.IndexService {
	ret := _m.Called(typ)

	var r0 es.IndexService
	if rf, ok := ret.Get(0).(func(string) es.IndexService); ok {
		r0 = rf(typ)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(es.IndexService)
		}
	}

	return r0
}
