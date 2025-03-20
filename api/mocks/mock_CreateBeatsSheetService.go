// Code generated by mockery v2.53.1. DO NOT EDIT.

package apimocks

import (
	context "github.com/a-novel-kit/context"
	mock "github.com/stretchr/testify/mock"

	models "github.com/a-novel/story-schematics/models"

	services "github.com/a-novel/story-schematics/internal/services"
)

// MockCreateBeatsSheetService is an autogenerated mock type for the CreateBeatsSheetService type
type MockCreateBeatsSheetService struct {
	mock.Mock
}

type MockCreateBeatsSheetService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCreateBeatsSheetService) EXPECT() *MockCreateBeatsSheetService_Expecter {
	return &MockCreateBeatsSheetService_Expecter{mock: &_m.Mock}
}

// CreateBeatsSheet provides a mock function with given fields: ctx, request
func (_m *MockCreateBeatsSheetService) CreateBeatsSheet(ctx context.Context, request services.CreateBeatsSheetRequest) (*models.BeatsSheet, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for CreateBeatsSheet")
	}

	var r0 *models.BeatsSheet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, services.CreateBeatsSheetRequest) (*models.BeatsSheet, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, services.CreateBeatsSheetRequest) *models.BeatsSheet); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.BeatsSheet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, services.CreateBeatsSheetRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCreateBeatsSheetService_CreateBeatsSheet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateBeatsSheet'
type MockCreateBeatsSheetService_CreateBeatsSheet_Call struct {
	*mock.Call
}

// CreateBeatsSheet is a helper method to define mock.On call
//   - ctx context.Context
//   - request services.CreateBeatsSheetRequest
func (_e *MockCreateBeatsSheetService_Expecter) CreateBeatsSheet(ctx interface{}, request interface{}) *MockCreateBeatsSheetService_CreateBeatsSheet_Call {
	return &MockCreateBeatsSheetService_CreateBeatsSheet_Call{Call: _e.mock.On("CreateBeatsSheet", ctx, request)}
}

func (_c *MockCreateBeatsSheetService_CreateBeatsSheet_Call) Run(run func(ctx context.Context, request services.CreateBeatsSheetRequest)) *MockCreateBeatsSheetService_CreateBeatsSheet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(services.CreateBeatsSheetRequest))
	})
	return _c
}

func (_c *MockCreateBeatsSheetService_CreateBeatsSheet_Call) Return(_a0 *models.BeatsSheet, _a1 error) *MockCreateBeatsSheetService_CreateBeatsSheet_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCreateBeatsSheetService_CreateBeatsSheet_Call) RunAndReturn(run func(context.Context, services.CreateBeatsSheetRequest) (*models.BeatsSheet, error)) *MockCreateBeatsSheetService_CreateBeatsSheet_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCreateBeatsSheetService creates a new instance of MockCreateBeatsSheetService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCreateBeatsSheetService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCreateBeatsSheetService {
	mock := &MockCreateBeatsSheetService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
