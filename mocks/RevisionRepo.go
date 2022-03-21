// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/model"
)

// RevisionRepo is an autogenerated mock type for the RevisionRepo type
type RevisionRepo struct {
	mock.Mock
}

type RevisionRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *RevisionRepo) EXPECT() *RevisionRepo_Expecter {
	return &RevisionRepo_Expecter{mock: &_m.Mock}
}

// CountPersonRelated provides a mock function with given fields: ctx, personID
func (_m *RevisionRepo) CountPersonRelated(ctx context.Context, personID uint32) (int64, error) {
	ret := _m.Called(ctx, personID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, uint32) int64); ok {
		r0 = rf(ctx, personID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, personID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevisionRepo_CountPersonRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountPersonRelated'
type RevisionRepo_CountPersonRelated_Call struct {
	*mock.Call
}

// CountPersonRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - personID uint32
func (_e *RevisionRepo_Expecter) CountPersonRelated(ctx interface{}, personID interface{}) *RevisionRepo_CountPersonRelated_Call {
	return &RevisionRepo_CountPersonRelated_Call{Call: _e.mock.On("CountPersonRelated", ctx, personID)}
}

func (_c *RevisionRepo_CountPersonRelated_Call) Run(run func(ctx context.Context, personID uint32)) *RevisionRepo_CountPersonRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *RevisionRepo_CountPersonRelated_Call) Return(_a0 int64, _a1 error) *RevisionRepo_CountPersonRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPersonRelated provides a mock function with given fields: ctx, id
func (_m *RevisionRepo) GetPersonRelated(ctx context.Context, id uint32) (model.Revision, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Revision
	if rf, ok := ret.Get(0).(func(context.Context, uint32) model.Revision); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Revision)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevisionRepo_GetPersonRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonRelated'
type RevisionRepo_GetPersonRelated_Call struct {
	*mock.Call
}

// GetPersonRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint32
func (_e *RevisionRepo_Expecter) GetPersonRelated(ctx interface{}, id interface{}) *RevisionRepo_GetPersonRelated_Call {
	return &RevisionRepo_GetPersonRelated_Call{Call: _e.mock.On("GetPersonRelated", ctx, id)}
}

func (_c *RevisionRepo_GetPersonRelated_Call) Run(run func(ctx context.Context, id uint32)) *RevisionRepo_GetPersonRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *RevisionRepo_GetPersonRelated_Call) Return(_a0 model.Revision, _a1 error) *RevisionRepo_GetPersonRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListPersonRelated provides a mock function with given fields: ctx, personID, limit, offset
func (_m *RevisionRepo) ListPersonRelated(ctx context.Context, personID uint32, limit int, offset int) ([]model.Revision, error) {
	ret := _m.Called(ctx, personID, limit, offset)

	var r0 []model.Revision
	if rf, ok := ret.Get(0).(func(context.Context, uint32, int, int) []model.Revision); ok {
		r0 = rf(ctx, personID, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Revision)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, int, int) error); ok {
		r1 = rf(ctx, personID, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevisionRepo_ListPersonRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListPersonRelated'
type RevisionRepo_ListPersonRelated_Call struct {
	*mock.Call
}

// ListPersonRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - personID uint32
//  - limit int
//  - offset int
func (_e *RevisionRepo_Expecter) ListPersonRelated(ctx interface{}, personID interface{}, limit interface{}, offset interface{}) *RevisionRepo_ListPersonRelated_Call {
	return &RevisionRepo_ListPersonRelated_Call{Call: _e.mock.On("ListPersonRelated", ctx, personID, limit, offset)}
}

func (_c *RevisionRepo_ListPersonRelated_Call) Run(run func(ctx context.Context, personID uint32, limit int, offset int)) *RevisionRepo_ListPersonRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *RevisionRepo_ListPersonRelated_Call) Return(_a0 []model.Revision, _a1 error) *RevisionRepo_ListPersonRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}