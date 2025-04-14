// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	entity "github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"
	mock "github.com/stretchr/testify/mock"
)

// DeviceUseCase is an autogenerated mock type for the DeviceUseCase type
type DeviceUseCase struct {
	mock.Mock
}

type DeviceUseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *DeviceUseCase) EXPECT() *DeviceUseCase_Expecter {
	return &DeviceUseCase_Expecter{mock: &_m.Mock}
}

// CreateDevice provides a mock function with given fields: device
func (_m *DeviceUseCase) CreateDevice(device *entity.Device) (*entity.Device, error) {
	ret := _m.Called(device)

	if len(ret) == 0 {
		panic("no return value specified for CreateDevice")
	}

	var r0 *entity.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Device) (*entity.Device, error)); ok {
		return rf(device)
	}
	if rf, ok := ret.Get(0).(func(*entity.Device) *entity.Device); ok {
		r0 = rf(device)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.Device) error); ok {
		r1 = rf(device)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceUseCase_CreateDevice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateDevice'
type DeviceUseCase_CreateDevice_Call struct {
	*mock.Call
}

// CreateDevice is a helper method to define mock.On call
//   - device *entity.Device
func (_e *DeviceUseCase_Expecter) CreateDevice(device interface{}) *DeviceUseCase_CreateDevice_Call {
	return &DeviceUseCase_CreateDevice_Call{Call: _e.mock.On("CreateDevice", device)}
}

func (_c *DeviceUseCase_CreateDevice_Call) Run(run func(device *entity.Device)) *DeviceUseCase_CreateDevice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Device))
	})
	return _c
}

func (_c *DeviceUseCase_CreateDevice_Call) Return(_a0 *entity.Device, _a1 error) *DeviceUseCase_CreateDevice_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DeviceUseCase_CreateDevice_Call) RunAndReturn(run func(*entity.Device) (*entity.Device, error)) *DeviceUseCase_CreateDevice_Call {
	_c.Call.Return(run)
	return _c
}

// FindDeviceVersion provides a mock function with given fields: deviceId
func (_m *DeviceUseCase) FindDeviceVersion(deviceId string) (*entity.Version, error) {
	ret := _m.Called(deviceId)

	if len(ret) == 0 {
		panic("no return value specified for FindDeviceVersion")
	}

	var r0 *entity.Version
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.Version, error)); ok {
		return rf(deviceId)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.Version); ok {
		r0 = rf(deviceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Version)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(deviceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceUseCase_FindDeviceVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindDeviceVersion'
type DeviceUseCase_FindDeviceVersion_Call struct {
	*mock.Call
}

// FindDeviceVersion is a helper method to define mock.On call
//   - deviceId string
func (_e *DeviceUseCase_Expecter) FindDeviceVersion(deviceId interface{}) *DeviceUseCase_FindDeviceVersion_Call {
	return &DeviceUseCase_FindDeviceVersion_Call{Call: _e.mock.On("FindDeviceVersion", deviceId)}
}

func (_c *DeviceUseCase_FindDeviceVersion_Call) Run(run func(deviceId string)) *DeviceUseCase_FindDeviceVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *DeviceUseCase_FindDeviceVersion_Call) Return(_a0 *entity.Version, _a1 error) *DeviceUseCase_FindDeviceVersion_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DeviceUseCase_FindDeviceVersion_Call) RunAndReturn(run func(string) (*entity.Version, error)) *DeviceUseCase_FindDeviceVersion_Call {
	_c.Call.Return(run)
	return _c
}

// ListDevices provides a mock function with given fields: categoryId
func (_m *DeviceUseCase) ListDevices(categoryId string) (*[]entity.Device, error) {
	ret := _m.Called(categoryId)

	if len(ret) == 0 {
		panic("no return value specified for ListDevices")
	}

	var r0 *[]entity.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*[]entity.Device, error)); ok {
		return rf(categoryId)
	}
	if rf, ok := ret.Get(0).(func(string) *[]entity.Device); ok {
		r0 = rf(categoryId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(categoryId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceUseCase_ListDevices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDevices'
type DeviceUseCase_ListDevices_Call struct {
	*mock.Call
}

// ListDevices is a helper method to define mock.On call
//   - categoryId string
func (_e *DeviceUseCase_Expecter) ListDevices(categoryId interface{}) *DeviceUseCase_ListDevices_Call {
	return &DeviceUseCase_ListDevices_Call{Call: _e.mock.On("ListDevices", categoryId)}
}

func (_c *DeviceUseCase_ListDevices_Call) Run(run func(categoryId string)) *DeviceUseCase_ListDevices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *DeviceUseCase_ListDevices_Call) Return(_a0 *[]entity.Device, _a1 error) *DeviceUseCase_ListDevices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DeviceUseCase_ListDevices_Call) RunAndReturn(run func(string) (*[]entity.Device, error)) *DeviceUseCase_ListDevices_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveDevice provides a mock function with given fields: deviceId
func (_m *DeviceUseCase) RemoveDevice(deviceId string) error {
	ret := _m.Called(deviceId)

	if len(ret) == 0 {
		panic("no return value specified for RemoveDevice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(deviceId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeviceUseCase_RemoveDevice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveDevice'
type DeviceUseCase_RemoveDevice_Call struct {
	*mock.Call
}

// RemoveDevice is a helper method to define mock.On call
//   - deviceId string
func (_e *DeviceUseCase_Expecter) RemoveDevice(deviceId interface{}) *DeviceUseCase_RemoveDevice_Call {
	return &DeviceUseCase_RemoveDevice_Call{Call: _e.mock.On("RemoveDevice", deviceId)}
}

func (_c *DeviceUseCase_RemoveDevice_Call) Run(run func(deviceId string)) *DeviceUseCase_RemoveDevice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *DeviceUseCase_RemoveDevice_Call) Return(_a0 error) *DeviceUseCase_RemoveDevice_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeviceUseCase_RemoveDevice_Call) RunAndReturn(run func(string) error) *DeviceUseCase_RemoveDevice_Call {
	_c.Call.Return(run)
	return _c
}

// SyncDeviceVersion provides a mock function with given fields: deviceId, versionName
func (_m *DeviceUseCase) SyncDeviceVersion(deviceId string, versionName string) (*entity.Device, error) {
	ret := _m.Called(deviceId, versionName)

	if len(ret) == 0 {
		panic("no return value specified for SyncDeviceVersion")
	}

	var r0 *entity.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*entity.Device, error)); ok {
		return rf(deviceId, versionName)
	}
	if rf, ok := ret.Get(0).(func(string, string) *entity.Device); ok {
		r0 = rf(deviceId, versionName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(deviceId, versionName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceUseCase_SyncDeviceVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SyncDeviceVersion'
type DeviceUseCase_SyncDeviceVersion_Call struct {
	*mock.Call
}

// SyncDeviceVersion is a helper method to define mock.On call
//   - deviceId string
//   - versionName string
func (_e *DeviceUseCase_Expecter) SyncDeviceVersion(deviceId interface{}, versionName interface{}) *DeviceUseCase_SyncDeviceVersion_Call {
	return &DeviceUseCase_SyncDeviceVersion_Call{Call: _e.mock.On("SyncDeviceVersion", deviceId, versionName)}
}

func (_c *DeviceUseCase_SyncDeviceVersion_Call) Run(run func(deviceId string, versionName string)) *DeviceUseCase_SyncDeviceVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *DeviceUseCase_SyncDeviceVersion_Call) Return(_a0 *entity.Device, _a1 error) *DeviceUseCase_SyncDeviceVersion_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DeviceUseCase_SyncDeviceVersion_Call) RunAndReturn(run func(string, string) (*entity.Device, error)) *DeviceUseCase_SyncDeviceVersion_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateHardware provides a mock function with given fields: deviceId, minimunHardware
func (_m *DeviceUseCase) UpdateHardware(deviceId string, minimunHardware float64) (*entity.Device, error) {
	ret := _m.Called(deviceId, minimunHardware)

	if len(ret) == 0 {
		panic("no return value specified for UpdateHardware")
	}

	var r0 *entity.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(string, float64) (*entity.Device, error)); ok {
		return rf(deviceId, minimunHardware)
	}
	if rf, ok := ret.Get(0).(func(string, float64) *entity.Device); ok {
		r0 = rf(deviceId, minimunHardware)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(string, float64) error); ok {
		r1 = rf(deviceId, minimunHardware)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceUseCase_UpdateHardware_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateHardware'
type DeviceUseCase_UpdateHardware_Call struct {
	*mock.Call
}

// UpdateHardware is a helper method to define mock.On call
//   - deviceId string
//   - minimunHardware float64
func (_e *DeviceUseCase_Expecter) UpdateHardware(deviceId interface{}, minimunHardware interface{}) *DeviceUseCase_UpdateHardware_Call {
	return &DeviceUseCase_UpdateHardware_Call{Call: _e.mock.On("UpdateHardware", deviceId, minimunHardware)}
}

func (_c *DeviceUseCase_UpdateHardware_Call) Run(run func(deviceId string, minimunHardware float64)) *DeviceUseCase_UpdateHardware_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(float64))
	})
	return _c
}

func (_c *DeviceUseCase_UpdateHardware_Call) Return(_a0 *entity.Device, _a1 error) *DeviceUseCase_UpdateHardware_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DeviceUseCase_UpdateHardware_Call) RunAndReturn(run func(string, float64) (*entity.Device, error)) *DeviceUseCase_UpdateHardware_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateTargetVersion provides a mock function with given fields: versionId, updateDurationHours, devicesId
func (_m *DeviceUseCase) UpdateTargetVersion(versionId string, updateDurationHours float64, devicesId ...string) (*entity.Device, error) {
	_va := make([]interface{}, len(devicesId))
	for _i := range devicesId {
		_va[_i] = devicesId[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, versionId, updateDurationHours)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTargetVersion")
	}

	var r0 *entity.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(string, float64, ...string) (*entity.Device, error)); ok {
		return rf(versionId, updateDurationHours, devicesId...)
	}
	if rf, ok := ret.Get(0).(func(string, float64, ...string) *entity.Device); ok {
		r0 = rf(versionId, updateDurationHours, devicesId...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(string, float64, ...string) error); ok {
		r1 = rf(versionId, updateDurationHours, devicesId...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceUseCase_UpdateTargetVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTargetVersion'
type DeviceUseCase_UpdateTargetVersion_Call struct {
	*mock.Call
}

// UpdateTargetVersion is a helper method to define mock.On call
//   - versionId string
//   - updateDurationHours float64
//   - devicesId ...string
func (_e *DeviceUseCase_Expecter) UpdateTargetVersion(versionId interface{}, updateDurationHours interface{}, devicesId ...interface{}) *DeviceUseCase_UpdateTargetVersion_Call {
	return &DeviceUseCase_UpdateTargetVersion_Call{Call: _e.mock.On("UpdateTargetVersion",
		append([]interface{}{versionId, updateDurationHours}, devicesId...)...)}
}

func (_c *DeviceUseCase_UpdateTargetVersion_Call) Run(run func(versionId string, updateDurationHours float64, devicesId ...string)) *DeviceUseCase_UpdateTargetVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(string), args[1].(float64), variadicArgs...)
	})
	return _c
}

func (_c *DeviceUseCase_UpdateTargetVersion_Call) Return(_a0 *entity.Device, _a1 error) *DeviceUseCase_UpdateTargetVersion_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DeviceUseCase_UpdateTargetVersion_Call) RunAndReturn(run func(string, float64, ...string) (*entity.Device, error)) *DeviceUseCase_UpdateTargetVersion_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateTargetVersionByCategory provides a mock function with given fields: categoryId, versionId, updateDuration
func (_m *DeviceUseCase) UpdateTargetVersionByCategory(categoryId string, versionId string, updateDuration float64) {
	_m.Called(categoryId, versionId, updateDuration)
}

// DeviceUseCase_UpdateTargetVersionByCategory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTargetVersionByCategory'
type DeviceUseCase_UpdateTargetVersionByCategory_Call struct {
	*mock.Call
}

// UpdateTargetVersionByCategory is a helper method to define mock.On call
//   - categoryId string
//   - versionId string
//   - updateDuration float64
func (_e *DeviceUseCase_Expecter) UpdateTargetVersionByCategory(categoryId interface{}, versionId interface{}, updateDuration interface{}) *DeviceUseCase_UpdateTargetVersionByCategory_Call {
	return &DeviceUseCase_UpdateTargetVersionByCategory_Call{Call: _e.mock.On("UpdateTargetVersionByCategory", categoryId, versionId, updateDuration)}
}

func (_c *DeviceUseCase_UpdateTargetVersionByCategory_Call) Run(run func(categoryId string, versionId string, updateDuration float64)) *DeviceUseCase_UpdateTargetVersionByCategory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(float64))
	})
	return _c
}

func (_c *DeviceUseCase_UpdateTargetVersionByCategory_Call) Return() *DeviceUseCase_UpdateTargetVersionByCategory_Call {
	_c.Call.Return()
	return _c
}

func (_c *DeviceUseCase_UpdateTargetVersionByCategory_Call) RunAndReturn(run func(string, string, float64)) *DeviceUseCase_UpdateTargetVersionByCategory_Call {
	_c.Run(run)
	return _c
}

// NewDeviceUseCase creates a new instance of DeviceUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeviceUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeviceUseCase {
	mock := &DeviceUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
