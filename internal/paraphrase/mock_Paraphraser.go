// Code generated by mockery v2.40.1. DO NOT EDIT.

package paraphrase

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockParaphraser is an autogenerated mock type for the Paraphraser type
type MockParaphraser struct {
	mock.Mock
}

type MockParaphraser_Expecter struct {
	mock *mock.Mock
}

func (_m *MockParaphraser) EXPECT() *MockParaphraser_Expecter {
	return &MockParaphraser_Expecter{mock: &_m.Mock}
}

// Paraphrase provides a mock function with given fields: ctx, tone, text
func (_m *MockParaphraser) Paraphrase(ctx context.Context, tone string, text string) (string, error) {
	ret := _m.Called(ctx, tone, text)

	if len(ret) == 0 {
		panic("no return value specified for Paraphrase")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, tone, text)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, tone, text)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, tone, text)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockParaphraser_Paraphrase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Paraphrase'
type MockParaphraser_Paraphrase_Call struct {
	*mock.Call
}

// Paraphrase is a helper method to define mock.On call
//   - ctx context.Context
//   - tone string
//   - text string
func (_e *MockParaphraser_Expecter) Paraphrase(ctx interface{}, tone interface{}, text interface{}) *MockParaphraser_Paraphrase_Call {
	return &MockParaphraser_Paraphrase_Call{Call: _e.mock.On("Paraphrase", ctx, tone, text)}
}

func (_c *MockParaphraser_Paraphrase_Call) Run(run func(ctx context.Context, tone string, text string)) *MockParaphraser_Paraphrase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockParaphraser_Paraphrase_Call) Return(_a0 string, _a1 error) *MockParaphraser_Paraphrase_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockParaphraser_Paraphrase_Call) RunAndReturn(run func(context.Context, string, string) (string, error)) *MockParaphraser_Paraphrase_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockParaphraser creates a new instance of MockParaphraser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockParaphraser(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockParaphraser {
	mock := &MockParaphraser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
