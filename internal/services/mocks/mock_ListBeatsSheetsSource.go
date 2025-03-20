// Code generated by mockery v2.53.1. DO NOT EDIT.

package servicesmocks

import (
	context "github.com/a-novel-kit/context"
	dao "github.com/a-novel/story-schematics/internal/dao"
	mock "github.com/stretchr/testify/mock"
)

// MockListBeatsSheetsSource is an autogenerated mock type for the ListBeatsSheetsSource type
type MockListBeatsSheetsSource struct {
	mock.Mock
}

type MockListBeatsSheetsSource_Expecter struct {
	mock *mock.Mock
}

func (_m *MockListBeatsSheetsSource) EXPECT() *MockListBeatsSheetsSource_Expecter {
	return &MockListBeatsSheetsSource_Expecter{mock: &_m.Mock}
}

// ListBeatsSheets provides a mock function with given fields: ctx, data
func (_m *MockListBeatsSheetsSource) ListBeatsSheets(ctx context.Context, data dao.ListBeatsSheetsData) ([]*dao.BeatsSheetPreviewEntity, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for ListBeatsSheets")
	}

	var r0 []*dao.BeatsSheetPreviewEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dao.ListBeatsSheetsData) ([]*dao.BeatsSheetPreviewEntity, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dao.ListBeatsSheetsData) []*dao.BeatsSheetPreviewEntity); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dao.BeatsSheetPreviewEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dao.ListBeatsSheetsData) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockListBeatsSheetsSource_ListBeatsSheets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListBeatsSheets'
type MockListBeatsSheetsSource_ListBeatsSheets_Call struct {
	*mock.Call
}

// ListBeatsSheets is a helper method to define mock.On call
//   - ctx context.Context
//   - data dao.ListBeatsSheetsData
func (_e *MockListBeatsSheetsSource_Expecter) ListBeatsSheets(ctx interface{}, data interface{}) *MockListBeatsSheetsSource_ListBeatsSheets_Call {
	return &MockListBeatsSheetsSource_ListBeatsSheets_Call{Call: _e.mock.On("ListBeatsSheets", ctx, data)}
}

func (_c *MockListBeatsSheetsSource_ListBeatsSheets_Call) Run(run func(ctx context.Context, data dao.ListBeatsSheetsData)) *MockListBeatsSheetsSource_ListBeatsSheets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dao.ListBeatsSheetsData))
	})
	return _c
}

func (_c *MockListBeatsSheetsSource_ListBeatsSheets_Call) Return(_a0 []*dao.BeatsSheetPreviewEntity, _a1 error) *MockListBeatsSheetsSource_ListBeatsSheets_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockListBeatsSheetsSource_ListBeatsSheets_Call) RunAndReturn(run func(context.Context, dao.ListBeatsSheetsData) ([]*dao.BeatsSheetPreviewEntity, error)) *MockListBeatsSheetsSource_ListBeatsSheets_Call {
	_c.Call.Return(run)
	return _c
}

// SelectLogline provides a mock function with given fields: ctx, data
func (_m *MockListBeatsSheetsSource) SelectLogline(ctx context.Context, data dao.SelectLoglineData) (*dao.LoglineEntity, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for SelectLogline")
	}

	var r0 *dao.LoglineEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dao.SelectLoglineData) (*dao.LoglineEntity, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dao.SelectLoglineData) *dao.LoglineEntity); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dao.LoglineEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dao.SelectLoglineData) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockListBeatsSheetsSource_SelectLogline_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectLogline'
type MockListBeatsSheetsSource_SelectLogline_Call struct {
	*mock.Call
}

// SelectLogline is a helper method to define mock.On call
//   - ctx context.Context
//   - data dao.SelectLoglineData
func (_e *MockListBeatsSheetsSource_Expecter) SelectLogline(ctx interface{}, data interface{}) *MockListBeatsSheetsSource_SelectLogline_Call {
	return &MockListBeatsSheetsSource_SelectLogline_Call{Call: _e.mock.On("SelectLogline", ctx, data)}
}

func (_c *MockListBeatsSheetsSource_SelectLogline_Call) Run(run func(ctx context.Context, data dao.SelectLoglineData)) *MockListBeatsSheetsSource_SelectLogline_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dao.SelectLoglineData))
	})
	return _c
}

func (_c *MockListBeatsSheetsSource_SelectLogline_Call) Return(_a0 *dao.LoglineEntity, _a1 error) *MockListBeatsSheetsSource_SelectLogline_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockListBeatsSheetsSource_SelectLogline_Call) RunAndReturn(run func(context.Context, dao.SelectLoglineData) (*dao.LoglineEntity, error)) *MockListBeatsSheetsSource_SelectLogline_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockListBeatsSheetsSource creates a new instance of MockListBeatsSheetsSource. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockListBeatsSheetsSource(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockListBeatsSheetsSource {
	mock := &MockListBeatsSheetsSource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
