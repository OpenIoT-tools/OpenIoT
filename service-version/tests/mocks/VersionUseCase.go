// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	entity "github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"
	mock "github.com/stretchr/testify/mock"
)

// VersionUseCase is an autogenerated mock type for the VersionUseCase type
type VersionUseCase struct {
	mock.Mock
}

type VersionUseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *VersionUseCase) EXPECT() *VersionUseCase_Expecter {
	return &VersionUseCase_Expecter{mock: &_m.Mock}
}

// CreateVersion provides a mock function with given fields: version
func (_m *VersionUseCase) CreateVersion(version *entity.Version) (*entity.Version, error) {
	ret := _m.Called(version)

	if len(ret) == 0 {
		panic("no return value specified for CreateVersion")
	}

	var r0 *entity.Version
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Version) (*entity.Version, error)); ok {
		return rf(version)
	}
	if rf, ok := ret.Get(0).(func(*entity.Version) *entity.Version); ok {
		r0 = rf(version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Version)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.Version) error); ok {
		r1 = rf(version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VersionUseCase_CreateVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateVersion'
type VersionUseCase_CreateVersion_Call struct {
	*mock.Call
}

// CreateVersion is a helper method to define mock.On call
//   - version *entity.Version
func (_e *VersionUseCase_Expecter) CreateVersion(version interface{}) *VersionUseCase_CreateVersion_Call {
	return &VersionUseCase_CreateVersion_Call{Call: _e.mock.On("CreateVersion", version)}
}

func (_c *VersionUseCase_CreateVersion_Call) Run(run func(version *entity.Version)) *VersionUseCase_CreateVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Version))
	})
	return _c
}

func (_c *VersionUseCase_CreateVersion_Call) Return(_a0 *entity.Version, _a1 error) *VersionUseCase_CreateVersion_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VersionUseCase_CreateVersion_Call) RunAndReturn(run func(*entity.Version) (*entity.Version, error)) *VersionUseCase_CreateVersion_Call {
	_c.Call.Return(run)
	return _c
}

// ListVersions provides a mock function with given fields: categoryId
func (_m *VersionUseCase) ListVersions(categoryId string) ([]*entity.Version, error) {
	ret := _m.Called(categoryId)

	if len(ret) == 0 {
		panic("no return value specified for ListVersions")
	}

	var r0 []*entity.Version
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*entity.Version, error)); ok {
		return rf(categoryId)
	}
	if rf, ok := ret.Get(0).(func(string) []*entity.Version); ok {
		r0 = rf(categoryId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Version)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(categoryId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VersionUseCase_ListVersions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListVersions'
type VersionUseCase_ListVersions_Call struct {
	*mock.Call
}

// ListVersions is a helper method to define mock.On call
//   - categoryId string
func (_e *VersionUseCase_Expecter) ListVersions(categoryId interface{}) *VersionUseCase_ListVersions_Call {
	return &VersionUseCase_ListVersions_Call{Call: _e.mock.On("ListVersions", categoryId)}
}

func (_c *VersionUseCase_ListVersions_Call) Run(run func(categoryId string)) *VersionUseCase_ListVersions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *VersionUseCase_ListVersions_Call) Return(_a0 []*entity.Version, _a1 error) *VersionUseCase_ListVersions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VersionUseCase_ListVersions_Call) RunAndReturn(run func(string) ([]*entity.Version, error)) *VersionUseCase_ListVersions_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveVersion provides a mock function with given fields: versionId
func (_m *VersionUseCase) RemoveVersion(versionId string) error {
	ret := _m.Called(versionId)

	if len(ret) == 0 {
		panic("no return value specified for RemoveVersion")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(versionId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VersionUseCase_RemoveVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveVersion'
type VersionUseCase_RemoveVersion_Call struct {
	*mock.Call
}

// RemoveVersion is a helper method to define mock.On call
//   - versionId string
func (_e *VersionUseCase_Expecter) RemoveVersion(versionId interface{}) *VersionUseCase_RemoveVersion_Call {
	return &VersionUseCase_RemoveVersion_Call{Call: _e.mock.On("RemoveVersion", versionId)}
}

func (_c *VersionUseCase_RemoveVersion_Call) Run(run func(versionId string)) *VersionUseCase_RemoveVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *VersionUseCase_RemoveVersion_Call) Return(_a0 error) *VersionUseCase_RemoveVersion_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *VersionUseCase_RemoveVersion_Call) RunAndReturn(run func(string) error) *VersionUseCase_RemoveVersion_Call {
	_c.Call.Return(run)
	return _c
}

// NewVersionUseCase creates a new instance of VersionUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVersionUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *VersionUseCase {
	mock := &VersionUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
