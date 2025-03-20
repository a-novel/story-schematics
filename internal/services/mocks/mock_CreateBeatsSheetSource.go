// Code generated by mockery v2.53.1. DO NOT EDIT.

package servicesmocks

import (
	context "github.com/a-novel-kit/context"
	dao "github.com/a-novel/story-schematics/internal/dao"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockCreateBeatsSheetSource is an autogenerated mock type for the CreateBeatsSheetSource type
type MockCreateBeatsSheetSource struct {
	mock.Mock
}

type MockCreateBeatsSheetSource_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCreateBeatsSheetSource) EXPECT() *MockCreateBeatsSheetSource_Expecter {
	return &MockCreateBeatsSheetSource_Expecter{mock: &_m.Mock}
}

// InsertBeatsSheet provides a mock function with given fields: ctx, data
func (_m *MockCreateBeatsSheetSource) InsertBeatsSheet(ctx context.Context, data dao.InsertBeatsSheetData) (*dao.BeatsSheetEntity, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for InsertBeatsSheet")
	}

	var r0 *dao.BeatsSheetEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dao.InsertBeatsSheetData) (*dao.BeatsSheetEntity, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dao.InsertBeatsSheetData) *dao.BeatsSheetEntity); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dao.BeatsSheetEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dao.InsertBeatsSheetData) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCreateBeatsSheetSource_InsertBeatsSheet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertBeatsSheet'
type MockCreateBeatsSheetSource_InsertBeatsSheet_Call struct {
	*mock.Call
}

// InsertBeatsSheet is a helper method to define mock.On call
//   - ctx context.Context
//   - data dao.InsertBeatsSheetData
func (_e *MockCreateBeatsSheetSource_Expecter) InsertBeatsSheet(ctx interface{}, data interface{}) *MockCreateBeatsSheetSource_InsertBeatsSheet_Call {
	return &MockCreateBeatsSheetSource_InsertBeatsSheet_Call{Call: _e.mock.On("InsertBeatsSheet", ctx, data)}
}

func (_c *MockCreateBeatsSheetSource_InsertBeatsSheet_Call) Run(run func(ctx context.Context, data dao.InsertBeatsSheetData)) *MockCreateBeatsSheetSource_InsertBeatsSheet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dao.InsertBeatsSheetData))
	})
	return _c
}

func (_c *MockCreateBeatsSheetSource_InsertBeatsSheet_Call) Return(_a0 *dao.BeatsSheetEntity, _a1 error) *MockCreateBeatsSheetSource_InsertBeatsSheet_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCreateBeatsSheetSource_InsertBeatsSheet_Call) RunAndReturn(run func(context.Context, dao.InsertBeatsSheetData) (*dao.BeatsSheetEntity, error)) *MockCreateBeatsSheetSource_InsertBeatsSheet_Call {
	_c.Call.Return(run)
	return _c
}

// SelectLogline provides a mock function with given fields: ctx, data
func (_m *MockCreateBeatsSheetSource) SelectLogline(ctx context.Context, data dao.SelectLoglineData) (*dao.LoglineEntity, error) {
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

// MockCreateBeatsSheetSource_SelectLogline_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectLogline'
type MockCreateBeatsSheetSource_SelectLogline_Call struct {
	*mock.Call
}

// SelectLogline is a helper method to define mock.On call
//   - ctx context.Context
//   - data dao.SelectLoglineData
func (_e *MockCreateBeatsSheetSource_Expecter) SelectLogline(ctx interface{}, data interface{}) *MockCreateBeatsSheetSource_SelectLogline_Call {
	return &MockCreateBeatsSheetSource_SelectLogline_Call{Call: _e.mock.On("SelectLogline", ctx, data)}
}

func (_c *MockCreateBeatsSheetSource_SelectLogline_Call) Run(run func(ctx context.Context, data dao.SelectLoglineData)) *MockCreateBeatsSheetSource_SelectLogline_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dao.SelectLoglineData))
	})
	return _c
}

func (_c *MockCreateBeatsSheetSource_SelectLogline_Call) Return(_a0 *dao.LoglineEntity, _a1 error) *MockCreateBeatsSheetSource_SelectLogline_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCreateBeatsSheetSource_SelectLogline_Call) RunAndReturn(run func(context.Context, dao.SelectLoglineData) (*dao.LoglineEntity, error)) *MockCreateBeatsSheetSource_SelectLogline_Call {
	_c.Call.Return(run)
	return _c
}

// SelectStoryPlan provides a mock function with given fields: ctx, data
func (_m *MockCreateBeatsSheetSource) SelectStoryPlan(ctx context.Context, data uuid.UUID) (*dao.StoryPlanEntity, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for SelectStoryPlan")
	}

	var r0 *dao.StoryPlanEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*dao.StoryPlanEntity, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *dao.StoryPlanEntity); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dao.StoryPlanEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCreateBeatsSheetSource_SelectStoryPlan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectStoryPlan'
type MockCreateBeatsSheetSource_SelectStoryPlan_Call struct {
	*mock.Call
}

// SelectStoryPlan is a helper method to define mock.On call
//   - ctx context.Context
//   - data uuid.UUID
func (_e *MockCreateBeatsSheetSource_Expecter) SelectStoryPlan(ctx interface{}, data interface{}) *MockCreateBeatsSheetSource_SelectStoryPlan_Call {
	return &MockCreateBeatsSheetSource_SelectStoryPlan_Call{Call: _e.mock.On("SelectStoryPlan", ctx, data)}
}

func (_c *MockCreateBeatsSheetSource_SelectStoryPlan_Call) Run(run func(ctx context.Context, data uuid.UUID)) *MockCreateBeatsSheetSource_SelectStoryPlan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockCreateBeatsSheetSource_SelectStoryPlan_Call) Return(_a0 *dao.StoryPlanEntity, _a1 error) *MockCreateBeatsSheetSource_SelectStoryPlan_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCreateBeatsSheetSource_SelectStoryPlan_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*dao.StoryPlanEntity, error)) *MockCreateBeatsSheetSource_SelectStoryPlan_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCreateBeatsSheetSource creates a new instance of MockCreateBeatsSheetSource. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCreateBeatsSheetSource(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCreateBeatsSheetSource {
	mock := &MockCreateBeatsSheetSource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
