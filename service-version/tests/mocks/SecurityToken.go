// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SecurityToken is an autogenerated mock type for the SecurityToken type
type SecurityToken struct {
	mock.Mock
}

type SecurityToken_Expecter struct {
	mock *mock.Mock
}

func (_m *SecurityToken) EXPECT() *SecurityToken_Expecter {
	return &SecurityToken_Expecter{mock: &_m.Mock}
}

// GenerateAsymmetricToken provides a mock function with given fields: tokenData, minutesLong
func (_m *SecurityToken) GenerateAsymmetricToken(tokenData map[string]interface{}, minutesLong int) (string, error) {
	ret := _m.Called(tokenData, minutesLong)

	if len(ret) == 0 {
		panic("no return value specified for GenerateAsymmetricToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(map[string]interface{}, int) (string, error)); ok {
		return rf(tokenData, minutesLong)
	}
	if rf, ok := ret.Get(0).(func(map[string]interface{}, int) string); ok {
		r0 = rf(tokenData, minutesLong)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(map[string]interface{}, int) error); ok {
		r1 = rf(tokenData, minutesLong)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SecurityToken_GenerateAsymmetricToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateAsymmetricToken'
type SecurityToken_GenerateAsymmetricToken_Call struct {
	*mock.Call
}

// GenerateAsymmetricToken is a helper method to define mock.On call
//   - tokenData map[string]interface{}
//   - minutesLong int
func (_e *SecurityToken_Expecter) GenerateAsymmetricToken(tokenData interface{}, minutesLong interface{}) *SecurityToken_GenerateAsymmetricToken_Call {
	return &SecurityToken_GenerateAsymmetricToken_Call{Call: _e.mock.On("GenerateAsymmetricToken", tokenData, minutesLong)}
}

func (_c *SecurityToken_GenerateAsymmetricToken_Call) Run(run func(tokenData map[string]interface{}, minutesLong int)) *SecurityToken_GenerateAsymmetricToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]interface{}), args[1].(int))
	})
	return _c
}

func (_c *SecurityToken_GenerateAsymmetricToken_Call) Return(_a0 string, _a1 error) *SecurityToken_GenerateAsymmetricToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SecurityToken_GenerateAsymmetricToken_Call) RunAndReturn(run func(map[string]interface{}, int) (string, error)) *SecurityToken_GenerateAsymmetricToken_Call {
	_c.Call.Return(run)
	return _c
}

// ValidateSymmetricalToken provides a mock function with given fields: tokenStr
func (_m *SecurityToken) ValidateSymmetricalToken(tokenStr string) (map[string]interface{}, error) {
	ret := _m.Called(tokenStr)

	if len(ret) == 0 {
		panic("no return value specified for ValidateSymmetricalToken")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (map[string]interface{}, error)); ok {
		return rf(tokenStr)
	}
	if rf, ok := ret.Get(0).(func(string) map[string]interface{}); ok {
		r0 = rf(tokenStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SecurityToken_ValidateSymmetricalToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateSymmetricalToken'
type SecurityToken_ValidateSymmetricalToken_Call struct {
	*mock.Call
}

// ValidateSymmetricalToken is a helper method to define mock.On call
//   - tokenStr string
func (_e *SecurityToken_Expecter) ValidateSymmetricalToken(tokenStr interface{}) *SecurityToken_ValidateSymmetricalToken_Call {
	return &SecurityToken_ValidateSymmetricalToken_Call{Call: _e.mock.On("ValidateSymmetricalToken", tokenStr)}
}

func (_c *SecurityToken_ValidateSymmetricalToken_Call) Run(run func(tokenStr string)) *SecurityToken_ValidateSymmetricalToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *SecurityToken_ValidateSymmetricalToken_Call) Return(_a0 map[string]interface{}, _a1 error) *SecurityToken_ValidateSymmetricalToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SecurityToken_ValidateSymmetricalToken_Call) RunAndReturn(run func(string) (map[string]interface{}, error)) *SecurityToken_ValidateSymmetricalToken_Call {
	_c.Call.Return(run)
	return _c
}

// NewSecurityToken creates a new instance of SecurityToken. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSecurityToken(t interface {
	mock.TestingT
	Cleanup(func())
}) *SecurityToken {
	mock := &SecurityToken{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
