// Code generated by mockery v2.53.1. DO NOT EDIT.

package servicesmocks

import (
	context "github.com/a-novel-kit/context"
	dao "github.com/a-novel/story-schematics/internal/dao"
	mock "github.com/stretchr/testify/mock"

	models "github.com/a-novel/story-schematics/models"
)

// MockCreateStoryPlanSource is an autogenerated mock type for the CreateStoryPlanSource type
type MockCreateStoryPlanSource struct {
	mock.Mock
}

type MockCreateStoryPlanSource_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCreateStoryPlanSource) EXPECT() *MockCreateStoryPlanSource_Expecter {
	return &MockCreateStoryPlanSource_Expecter{mock: &_m.Mock}
}

// InsertStoryPlan provides a mock function with given fields: ctx, data
func (_m *MockCreateStoryPlanSource) InsertStoryPlan(ctx context.Context, data dao.InsertStoryPlanData) (*dao.StoryPlanEntity, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for InsertStoryPlan")
	}

	var r0 *dao.StoryPlanEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dao.InsertStoryPlanData) (*dao.StoryPlanEntity, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dao.InsertStoryPlanData) *dao.StoryPlanEntity); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dao.StoryPlanEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dao.InsertStoryPlanData) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCreateStoryPlanSource_InsertStoryPlan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertStoryPlan'
type MockCreateStoryPlanSource_InsertStoryPlan_Call struct {
	*mock.Call
}

// InsertStoryPlan is a helper method to define mock.On call
//   - ctx context.Context
//   - data dao.InsertStoryPlanData
func (_e *MockCreateStoryPlanSource_Expecter) InsertStoryPlan(ctx interface{}, data interface{}) *MockCreateStoryPlanSource_InsertStoryPlan_Call {
	return &MockCreateStoryPlanSource_InsertStoryPlan_Call{Call: _e.mock.On("InsertStoryPlan", ctx, data)}
}

func (_c *MockCreateStoryPlanSource_InsertStoryPlan_Call) Run(run func(ctx context.Context, data dao.InsertStoryPlanData)) *MockCreateStoryPlanSource_InsertStoryPlan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dao.InsertStoryPlanData))
	})
	return _c
}

func (_c *MockCreateStoryPlanSource_InsertStoryPlan_Call) Return(_a0 *dao.StoryPlanEntity, _a1 error) *MockCreateStoryPlanSource_InsertStoryPlan_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCreateStoryPlanSource_InsertStoryPlan_Call) RunAndReturn(run func(context.Context, dao.InsertStoryPlanData) (*dao.StoryPlanEntity, error)) *MockCreateStoryPlanSource_InsertStoryPlan_Call {
	_c.Call.Return(run)
	return _c
}

// SelectSlugIteration provides a mock function with given fields: ctx, data
func (_m *MockCreateStoryPlanSource) SelectSlugIteration(ctx context.Context, data dao.SelectSlugIterationData) (models.Slug, int, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for SelectSlugIteration")
	}

	var r0 models.Slug
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, dao.SelectSlugIterationData) (models.Slug, int, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dao.SelectSlugIterationData) models.Slug); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(models.Slug)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dao.SelectSlugIterationData) int); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, dao.SelectSlugIterationData) error); ok {
		r2 = rf(ctx, data)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockCreateStoryPlanSource_SelectSlugIteration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectSlugIteration'
type MockCreateStoryPlanSource_SelectSlugIteration_Call struct {
	*mock.Call
}

// SelectSlugIteration is a helper method to define mock.On call
//   - ctx context.Context
//   - data dao.SelectSlugIterationData
func (_e *MockCreateStoryPlanSource_Expecter) SelectSlugIteration(ctx interface{}, data interface{}) *MockCreateStoryPlanSource_SelectSlugIteration_Call {
	return &MockCreateStoryPlanSource_SelectSlugIteration_Call{Call: _e.mock.On("SelectSlugIteration", ctx, data)}
}

func (_c *MockCreateStoryPlanSource_SelectSlugIteration_Call) Run(run func(ctx context.Context, data dao.SelectSlugIterationData)) *MockCreateStoryPlanSource_SelectSlugIteration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dao.SelectSlugIterationData))
	})
	return _c
}

func (_c *MockCreateStoryPlanSource_SelectSlugIteration_Call) Return(_a0 models.Slug, _a1 int, _a2 error) *MockCreateStoryPlanSource_SelectSlugIteration_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockCreateStoryPlanSource_SelectSlugIteration_Call) RunAndReturn(run func(context.Context, dao.SelectSlugIterationData) (models.Slug, int, error)) *MockCreateStoryPlanSource_SelectSlugIteration_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCreateStoryPlanSource creates a new instance of MockCreateStoryPlanSource. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCreateStoryPlanSource(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCreateStoryPlanSource {
	mock := &MockCreateStoryPlanSource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
