// Code generated by mockery v2.53.0. DO NOT EDIT.

package mocks

import (
	da "github.com/rollkit/rollkit/core/da"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// DA is an autogenerated mock type for the DA type
type DA struct {
	mock.Mock
}

// Commit provides a mock function with given fields: ctx, blobs, namespace
func (_m *DA) Commit(ctx context.Context, blobs [][]byte, namespace []byte) ([][]byte, error) {
	ret := _m.Called(ctx, blobs, namespace)

	if len(ret) == 0 {
		panic("no return value specified for Commit")
	}

	var r0 [][]byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, [][]byte, []byte) ([][]byte, error)); ok {
		return rf(ctx, blobs, namespace)
	}
	if rf, ok := ret.Get(0).(func(context.Context, [][]byte, []byte) [][]byte); ok {
		r0 = rf(ctx, blobs, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, [][]byte, []byte) error); ok {
		r1 = rf(ctx, blobs, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GasMultiplier provides a mock function with given fields: ctx
func (_m *DA) GasMultiplier(ctx context.Context) (float64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GasMultiplier")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (float64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) float64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GasPrice provides a mock function with given fields: ctx
func (_m *DA) GasPrice(ctx context.Context) (float64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GasPrice")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (float64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) float64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, ids, namespace
func (_m *DA) Get(ctx context.Context, ids [][]byte, namespace []byte) ([][]byte, error) {
	ret := _m.Called(ctx, ids, namespace)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 [][]byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, [][]byte, []byte) ([][]byte, error)); ok {
		return rf(ctx, ids, namespace)
	}
	if rf, ok := ret.Get(0).(func(context.Context, [][]byte, []byte) [][]byte); ok {
		r0 = rf(ctx, ids, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, [][]byte, []byte) error); ok {
		r1 = rf(ctx, ids, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIDs provides a mock function with given fields: ctx, height, namespace
func (_m *DA) GetIDs(ctx context.Context, height uint64, namespace []byte) (*da.GetIDsResult, error) {
	ret := _m.Called(ctx, height, namespace)

	if len(ret) == 0 {
		panic("no return value specified for GetIDs")
	}

	var r0 *da.GetIDsResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, []byte) (*da.GetIDsResult, error)); ok {
		return rf(ctx, height, namespace)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, []byte) *da.GetIDsResult); ok {
		r0 = rf(ctx, height, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*da.GetIDsResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, []byte) error); ok {
		r1 = rf(ctx, height, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProofs provides a mock function with given fields: ctx, ids, namespace
func (_m *DA) GetProofs(ctx context.Context, ids [][]byte, namespace []byte) ([][]byte, error) {
	ret := _m.Called(ctx, ids, namespace)

	if len(ret) == 0 {
		panic("no return value specified for GetProofs")
	}

	var r0 [][]byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, [][]byte, []byte) ([][]byte, error)); ok {
		return rf(ctx, ids, namespace)
	}
	if rf, ok := ret.Get(0).(func(context.Context, [][]byte, []byte) [][]byte); ok {
		r0 = rf(ctx, ids, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, [][]byte, []byte) error); ok {
		r1 = rf(ctx, ids, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MaxBlobSize provides a mock function with given fields: ctx
func (_m *DA) MaxBlobSize(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for MaxBlobSize")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Submit provides a mock function with given fields: ctx, blobs, gasPrice, namespace
func (_m *DA) Submit(ctx context.Context, blobs [][]byte, gasPrice float64, namespace []byte) ([][]byte, error) {
	ret := _m.Called(ctx, blobs, gasPrice, namespace)

	if len(ret) == 0 {
		panic("no return value specified for Submit")
	}

	var r0 [][]byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, [][]byte, float64, []byte) ([][]byte, error)); ok {
		return rf(ctx, blobs, gasPrice, namespace)
	}
	if rf, ok := ret.Get(0).(func(context.Context, [][]byte, float64, []byte) [][]byte); ok {
		r0 = rf(ctx, blobs, gasPrice, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, [][]byte, float64, []byte) error); ok {
		r1 = rf(ctx, blobs, gasPrice, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubmitWithOptions provides a mock function with given fields: ctx, blobs, gasPrice, namespace, options
func (_m *DA) SubmitWithOptions(ctx context.Context, blobs [][]byte, gasPrice float64, namespace []byte, options []byte) ([][]byte, error) {
	ret := _m.Called(ctx, blobs, gasPrice, namespace, options)

	if len(ret) == 0 {
		panic("no return value specified for SubmitWithOptions")
	}

	var r0 [][]byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, [][]byte, float64, []byte, []byte) ([][]byte, error)); ok {
		return rf(ctx, blobs, gasPrice, namespace, options)
	}
	if rf, ok := ret.Get(0).(func(context.Context, [][]byte, float64, []byte, []byte) [][]byte); ok {
		r0 = rf(ctx, blobs, gasPrice, namespace, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, [][]byte, float64, []byte, []byte) error); ok {
		r1 = rf(ctx, blobs, gasPrice, namespace, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validate provides a mock function with given fields: ctx, ids, proofs, namespace
func (_m *DA) Validate(ctx context.Context, ids [][]byte, proofs [][]byte, namespace []byte) ([]bool, error) {
	ret := _m.Called(ctx, ids, proofs, namespace)

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 []bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, [][]byte, [][]byte, []byte) ([]bool, error)); ok {
		return rf(ctx, ids, proofs, namespace)
	}
	if rf, ok := ret.Get(0).(func(context.Context, [][]byte, [][]byte, []byte) []bool); ok {
		r0 = rf(ctx, ids, proofs, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bool)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, [][]byte, [][]byte, []byte) error); ok {
		r1 = rf(ctx, ids, proofs, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDA creates a new instance of DA. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDA(t interface {
	mock.TestingT
	Cleanup(func())
}) *DA {
	mock := &DA{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
