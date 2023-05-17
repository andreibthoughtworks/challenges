// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	domain "polaris/internal/application/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Ads is an autogenerated mock type for the Ads type
type Ads struct {
	mock.Mock
}

type Ads_Expecter struct {
	mock *mock.Mock
}

func (_m *Ads) EXPECT() *Ads_Expecter {
	return &Ads_Expecter{mock: &_m.Mock}
}

// FindAll provides a mock function with given fields:
func (_m *Ads) FindAll() []domain.Ad {
	ret := _m.Called()

	var r0 []domain.Ad
	if rf, ok := ret.Get(0).(func() []domain.Ad); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Ad)
		}
	}

	return r0
}

// Ads_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type Ads_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
func (_e *Ads_Expecter) FindAll() *Ads_FindAll_Call {
	return &Ads_FindAll_Call{Call: _e.mock.On("FindAll")}
}

func (_c *Ads_FindAll_Call) Run(run func()) *Ads_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Ads_FindAll_Call) Return(_a0 []domain.Ad) *Ads_FindAll_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Ads_FindAll_Call) RunAndReturn(run func() []domain.Ad) *Ads_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindById provides a mock function with given fields: id
func (_m *Ads) FindById(id uuid.UUID) (*domain.Ad, error) {
	ret := _m.Called(id)

	var r0 *domain.Ad
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*domain.Ad, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *domain.Ad); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Ad)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ads_FindById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindById'
type Ads_FindById_Call struct {
	*mock.Call
}

// FindById is a helper method to define mock.On call
//   - id uuid.UUID
func (_e *Ads_Expecter) FindById(id interface{}) *Ads_FindById_Call {
	return &Ads_FindById_Call{Call: _e.mock.On("FindById", id)}
}

func (_c *Ads_FindById_Call) Run(run func(id uuid.UUID)) *Ads_FindById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *Ads_FindById_Call) Return(_a0 *domain.Ad, _a1 error) *Ads_FindById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Ads_FindById_Call) RunAndReturn(run func(uuid.UUID) (*domain.Ad, error)) *Ads_FindById_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: ad
func (_m *Ads) Save(ad domain.Ad) domain.Ad {
	ret := _m.Called(ad)

	var r0 domain.Ad
	if rf, ok := ret.Get(0).(func(domain.Ad) domain.Ad); ok {
		r0 = rf(ad)
	} else {
		r0 = ret.Get(0).(domain.Ad)
	}

	return r0
}

// Ads_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type Ads_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ad domain.Ad
func (_e *Ads_Expecter) Save(ad interface{}) *Ads_Save_Call {
	return &Ads_Save_Call{Call: _e.mock.On("Save", ad)}
}

func (_c *Ads_Save_Call) Run(run func(ad domain.Ad)) *Ads_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.Ad))
	})
	return _c
}

func (_c *Ads_Save_Call) Return(_a0 domain.Ad) *Ads_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Ads_Save_Call) RunAndReturn(run func(domain.Ad) domain.Ad) *Ads_Save_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewAds interface {
	mock.TestingT
	Cleanup(func())
}

// NewAds creates a new instance of Ads. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAds(t mockConstructorTestingTNewAds) *Ads {
	mock := &Ads{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}