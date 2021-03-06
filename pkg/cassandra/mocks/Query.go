// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package mocks

import cassandra "github.com/uber/jaeger/pkg/cassandra"
import mock "github.com/stretchr/testify/mock"

// Query is an autogenerated mock type for the Query type
type Query struct {
	mock.Mock
}

// Bind provides a mock function with given fields: v
func (_m *Query) Bind(v ...interface{}) cassandra.Query {
	ret := _m.Called(v)

	var r0 cassandra.Query
	if rf, ok := ret.Get(0).(func(...interface{}) cassandra.Query); ok {
		r0 = rf(v...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cassandra.Query)
		}
	}

	return r0
}

// Consistency provides a mock function with given fields: level
func (_m *Query) Consistency(level cassandra.Consistency) cassandra.Query {
	ret := _m.Called(level)

	var r0 cassandra.Query
	if rf, ok := ret.Get(0).(func(cassandra.Consistency) cassandra.Query); ok {
		r0 = rf(level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cassandra.Query)
		}
	}

	return r0
}

// PageSize provides a mock function with given fields: n
func (_m *Query) PageSize(n int) cassandra.Query {
	ret := _m.Called(n)

	var r0 cassandra.Query
	if rf, ok := ret.Get(0).(func(int) cassandra.Query); ok {
		r0 = rf(n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cassandra.Query)
		}
	}

	return r0
}

// Exec provides a mock function with given fields:
func (_m *Query) Exec() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Iter provides a mock function with given fields:
func (_m *Query) Iter() cassandra.Iterator {
	ret := _m.Called()

	var r0 cassandra.Iterator
	if rf, ok := ret.Get(0).(func() cassandra.Iterator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cassandra.Iterator)
		}
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *Query) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ScanCAS provides a mock function with given fields: dest
func (_m *Query) ScanCAS(dest ...interface{}) (bool, error) {
	ret := _m.Called(dest)

	var r0 bool
	if rf, ok := ret.Get(0).(func(...interface{}) bool); ok {
		r0 = rf(dest...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...interface{}) error); ok {
		r1 = rf(dest...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ cassandra.Query = (*Query)(nil)
