package usecases_test

import (
	"fmt"
	"testing"

	"github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"
	"github.com/OpenIoT-tools/OpenIoT/internal/core/usecases"
	"github.com/OpenIoT-tools/OpenIoT/tests/fixture"
	"github.com/OpenIoT-tools/OpenIoT/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDevice_FindDeviceVersion(t *testing.T) {
	repository := mocks.NewDeviceRespository(t)
	deploy := mocks.NewDeployUseCase(t)
	versionRepository := mocks.NewVersionRespository(t)
	device := fixture.GetDevices(1)[0]
	type testData struct {
		name          string
		deviceIdParam string
		expectVersion *entity.Version
		expectErr     error
		setRepository func()
	}
	cases := []testData{
		{
			name:          "when device is not found, should return an error",
			deviceIdParam: "notFound",
			expectVersion: nil,
			expectErr:     fmt.Errorf("device not found"),
			setRepository: func() { repository.On("FindDeviceById", "notFound").Return(nil, fmt.Errorf("device not found")) },
		},
		{
			name:          "when device is found, should return an error",
			deviceIdParam: "12345789",
			expectVersion: device.GetTargetVersion(),
			expectErr:     nil,
			setRepository: func() { repository.On("FindDeviceById", "12345789").Return(device, nil) },
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setRepository()
			deviceService := usecases.NewDeviceService(repository, versionRepository, deploy)
			version, err := deviceService.FindDeviceVersion(testCase.deviceIdParam)

			assert.Equal(t, testCase.expectVersion, version)
			assert.Equal(t, testCase.expectErr, err)
		})
	}
}

func TestDevice_CreateDevice(t *testing.T) {
	repository := mocks.NewDeviceRespository(t)
	deploy := mocks.NewDeployUseCase(t)
	versionRepository := mocks.NewVersionRespository(t)
	device := fixture.GetDevices(1)[0]

	type testCase struct {
		name          string
		deviceParam   *entity.Device
		expectDevice  *entity.Device
		expectErr     error
		setRepository func() *mock.Call
		setDeploy     func() *mock.Call
	}

	cases := []testCase{
		{
			name:         "when it is not possible to create the device, it should return an error",
			deviceParam:  device,
			expectDevice: nil,
			expectErr:    fmt.Errorf("cannot create device"),
			setRepository: func() *mock.Call {
				return repository.On("CreateDevice", device).Return(nil, fmt.Errorf("cannot create device"))
			},
			setDeploy: func() *mock.Call { return nil },
		},
		{
			name:          "when it is not possible to send an update, should return an error",
			deviceParam:   device,
			expectDevice:  nil,
			expectErr:     fmt.Errorf("cannot sent update to device"),
			setRepository: func() *mock.Call { return repository.On("CreateDevice", device).Return(device, nil) },
			setDeploy: func() *mock.Call {
				return deploy.On("SendUpdate", 0.0, device).Return(0, 0, fmt.Errorf("cannot sent update to device"))
			},
		},
		{
			name:          "when a valid device is sent, it should return the created device",
			deviceParam:   device,
			expectDevice:  device,
			expectErr:     nil,
			setRepository: func() *mock.Call { return repository.On("CreateDevice", device).Return(device, nil) },
			setDeploy:     func() *mock.Call { return deploy.On("SendUpdate", 0.0, device).Return(10, 1, nil) },
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			repositoryCall := tCase.setRepository()
			deployCall := tCase.setDeploy()
			deviceService := usecases.NewDeviceService(repository, versionRepository, deploy)
			createdDevice, err := deviceService.CreateDevice(device)

			assert.Equal(t, tCase.expectDevice, createdDevice)
			assert.Equal(t, tCase.expectErr, err)

			unsetMocks(repositoryCall, deployCall)
		})
	}
}

func TestDevice_UpdateTargetVersion(t *testing.T) {
	repository := mocks.NewDeviceRespository(t)
	deploy := mocks.NewDeployUseCase(t)
	versionRepository := mocks.NewVersionRespository(t)
	devices := fixture.GetDevices(2)
	version, _ := entity.NewVersion("test", "test", 1.0, 3.0, devices[0].GetCategory())

	type testCase struct {
		name                     string
		versionParam             string
		updateDurationHoursParam float64
		devicesIdParams          []string
		expectDevices            []*entity.Device
		expectErr                error
		setRepository            func() []*mock.Call
		setDeploy                func() []*mock.Call
		setVersion               func() []*mock.Call
	}

	cases := []testCase{
		{
			name:                     "when valid devices are sent, should return devices",
			versionParam:             version.GetId(),
			updateDurationHoursParam: 1,
			devicesIdParams:          []string{devices[0].GetId(), devices[1].GetId()},
			expectDevices:            devices,
			expectErr:                nil,
			setRepository: func() []*mock.Call {
				return []*mock.Call{
					repository.On("ListDevicesById", devices[0].GetId(), devices[1].GetId()).Return(devices, nil),
					repository.On("UpdateTargetVersion", devices[0], devices[1]).Return(devices, nil),
				}
			},
			setDeploy: func() []*mock.Call {
				return []*mock.Call{deploy.On("SendUpdate", 1.0, devices[0], devices[1]).Return(1, 1, nil)}
			},
			setVersion: func() []*mock.Call {
				return []*mock.Call{versionRepository.On("FindById", version.GetId()).Return(version, nil)}
			},
		},
		{
			name:                     "when unable to list devices, should return an error",
			versionParam:             version.GetId(),
			updateDurationHoursParam: 1,
			devicesIdParams:          []string{devices[0].GetId(), devices[1].GetId()},
			expectDevices:            nil,
			expectErr:                fmt.Errorf("error in list"),
			setRepository: func() []*mock.Call {
				return []*mock.Call{
					repository.On("ListDevicesById", devices[0].GetId(), devices[1].GetId()).Return(nil, fmt.Errorf("error in list")),
				}
			},
			setDeploy:  func() []*mock.Call { return nil },
			setVersion: func() []*mock.Call { return nil },
		},
		{
			name:                     "when the version is not found, it should return an error",
			versionParam:             version.GetId(),
			updateDurationHoursParam: 1,
			devicesIdParams:          []string{devices[0].GetId(), devices[1].GetId()},
			expectDevices:            nil,
			expectErr:                fmt.Errorf("version not found"),
			setRepository: func() []*mock.Call {
				return []*mock.Call{
					repository.On("ListDevicesById", devices[0].GetId(), devices[1].GetId()).Return(devices, nil),
				}
			},
			setDeploy: func() []*mock.Call { return nil },
			setVersion: func() []*mock.Call {
				return []*mock.Call{versionRepository.On("FindById", version.GetId()).Return(version, fmt.Errorf("version not found"))}
			},
		},
		{
			name:                     "when unable to update target version, should return an error",
			versionParam:             version.GetId(),
			updateDurationHoursParam: 1,
			devicesIdParams:          []string{devices[0].GetId(), devices[1].GetId()},
			expectDevices:            nil,
			expectErr:                fmt.Errorf("cannot update version"),
			setRepository: func() []*mock.Call {
				return []*mock.Call{
					repository.On("ListDevicesById", devices[0].GetId(), devices[1].GetId()).Return(devices, nil),
					repository.On("UpdateTargetVersion", devices[0], devices[1]).Return(nil, fmt.Errorf("cannot update version")),
				}
			},
			setDeploy: func() []*mock.Call { return nil },
			setVersion: func() []*mock.Call {
				return []*mock.Call{versionRepository.On("FindById", version.GetId()).Return(version, nil)}
			},
		},
		{
			name:                     "when it is not possible to send the update, it should return an error",
			versionParam:             version.GetId(),
			updateDurationHoursParam: 1,
			devicesIdParams:          []string{devices[0].GetId(), devices[1].GetId()},
			expectDevices:            nil,
			expectErr:                fmt.Errorf("update error"),
			setRepository: func() []*mock.Call {
				return []*mock.Call{
					repository.On("ListDevicesById", devices[0].GetId(), devices[1].GetId()).Return(devices, nil),
					repository.On("UpdateTargetVersion", devices[0], devices[1]).Return(devices, nil),
				}
			},
			setDeploy: func() []*mock.Call {
				return []*mock.Call{deploy.On("SendUpdate", 1.0, devices[0], devices[1]).Return(0, 0, fmt.Errorf("update error"))}
			},
			setVersion: func() []*mock.Call {
				return []*mock.Call{versionRepository.On("FindById", version.GetId()).Return(version, nil)}
			},
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			calls := tCase.setRepository()
			calls = append(calls, tCase.setDeploy()...)
			calls = append(calls, tCase.setVersion()...)

			deviceService := usecases.NewDeviceService(repository, versionRepository, deploy)
			createdDevice, err := deviceService.UpdateTargetVersion(tCase.versionParam, tCase.updateDurationHoursParam, tCase.devicesIdParams...)

			assert.Equal(t, tCase.expectDevices, createdDevice)
			assert.Equal(t, tCase.expectErr, err)

			unsetMocks(calls...)
		})
	}
}

func TestDevice_UpdateTargetVersionByCategory(t *testing.T) {
	repository := mocks.NewDeviceRespository(t)
	deploy := mocks.NewDeployUseCase(t)
	versionRepository := mocks.NewVersionRespository(t)
	devices := fixture.GetDevices(2)
	version, _ := entity.NewVersion("test", "test", 1.0, 3.0, devices[0].GetCategory())

	type testCase struct {
		name                        string
		versionParam, categoryParam string
		updateDurationHoursParam    float64
		expectDevices               []*entity.Device
		expectErr                   error
		setRepository               func() []*mock.Call
		setDeploy                   func() []*mock.Call
		setVersion                  func() []*mock.Call
	}

	cases := []testCase{
		{
			name:                     "when valid category are sent, should return devices",
			versionParam:             version.GetId(),
			updateDurationHoursParam: 1,
			categoryParam:            devices[0].GetCategory().GetId(),
			expectDevices:            devices,
			expectErr:                nil,
			setRepository: func() []*mock.Call {
				return []*mock.Call{
					repository.On("ListDevicesByCategory", devices[0].GetCategory().GetId()).Return(devices, nil),
					repository.On("UpdateTargetVersion", devices[0], devices[1]).Return(devices, nil),
				}
			},
			setDeploy: func() []*mock.Call {
				return []*mock.Call{deploy.On("SendUpdate", 1.0, devices[0], devices[1]).Return(1, 1, nil)}
			},
			setVersion: func() []*mock.Call {
				return []*mock.Call{versionRepository.On("FindById", version.GetId()).Return(version, nil)}
			},
		},
		{
			name:                     "when unable to list devices, should return an error",
			versionParam:             version.GetId(),
			updateDurationHoursParam: 1,
			categoryParam:            devices[0].GetCategory().GetId(),
			expectDevices:            nil,
			expectErr:                fmt.Errorf("error in list"),
			setRepository: func() []*mock.Call {
				return []*mock.Call{
					repository.On("ListDevicesByCategory", devices[1].GetCategory().GetId()).Return(nil, fmt.Errorf("error in list")),
				}
			},
			setDeploy:  func() []*mock.Call { return nil },
			setVersion: func() []*mock.Call { return nil },
		},
		{
			name:                     "when the version is not found, it should return an error",
			versionParam:             version.GetId(),
			updateDurationHoursParam: 1,
			categoryParam:            devices[0].GetCategory().GetId(),
			expectDevices:            nil,
			expectErr:                fmt.Errorf("version not found"),
			setRepository: func() []*mock.Call {
				return []*mock.Call{
					repository.On("ListDevicesByCategory", devices[0].GetCategory().GetId()).Return(devices, nil),
				}
			},
			setDeploy: func() []*mock.Call { return nil },
			setVersion: func() []*mock.Call {
				return []*mock.Call{versionRepository.On("FindById", version.GetId()).Return(version, fmt.Errorf("version not found"))}
			},
		},
		{
			name:                     "when unable to update target version, should return an error",
			versionParam:             version.GetId(),
			updateDurationHoursParam: 1,
			categoryParam:            devices[0].GetCategory().GetId(),
			expectDevices:            nil,
			expectErr:                fmt.Errorf("cannot update version"),
			setRepository: func() []*mock.Call {
				return []*mock.Call{
					repository.On("ListDevicesByCategory", devices[0].GetCategory().GetId()).Return(devices, nil),
					repository.On("UpdateTargetVersion", devices[0], devices[1]).Return(nil, fmt.Errorf("cannot update version")),
				}
			},
			setDeploy: func() []*mock.Call { return nil },
			setVersion: func() []*mock.Call {
				return []*mock.Call{versionRepository.On("FindById", version.GetId()).Return(version, nil)}
			},
		},
		{
			name:                     "when it is not possible to send the update, it should return an error",
			versionParam:             version.GetId(),
			updateDurationHoursParam: 1,
			categoryParam:            devices[0].GetCategory().GetId(),
			expectDevices:            nil,
			expectErr:                fmt.Errorf("update error"),
			setRepository: func() []*mock.Call {
				return []*mock.Call{
					repository.On("ListDevicesByCategory", devices[0].GetCategory().GetId()).Return(devices, nil),
					repository.On("UpdateTargetVersion", devices[0], devices[1]).Return(devices, nil),
				}
			},
			setDeploy: func() []*mock.Call {
				return []*mock.Call{deploy.On("SendUpdate", 1.0, devices[0], devices[1]).Return(0, 0, fmt.Errorf("update error"))}
			},
			setVersion: func() []*mock.Call {
				return []*mock.Call{versionRepository.On("FindById", version.GetId()).Return(version, nil)}
			},
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			calls := tCase.setRepository()
			calls = append(calls, tCase.setDeploy()...)
			calls = append(calls, tCase.setVersion()...)

			deviceService := usecases.NewDeviceService(repository, versionRepository, deploy)
			createdDevice, err := deviceService.UpdateTargetVersionByCategory(tCase.categoryParam, tCase.versionParam, tCase.updateDurationHoursParam)

			assert.Equal(t, tCase.expectDevices, createdDevice)
			assert.Equal(t, tCase.expectErr, err)

			unsetMocks(calls...)
		})
	}
}

func TestDevice_SyncDeviceVersion(t *testing.T) {
	repository := mocks.NewDeviceRespository(t)
	deploy := mocks.NewDeployUseCase(t)
	versionRepository := mocks.NewVersionRespository(t)

	t.Run("when a valid version is submitted, it should return the updated device", func(t *testing.T) {
		device := fixture.GetDevices(1)[0]

		calls := make([]*mock.Call, 0, 3)
		calls = append(calls, repository.On("FindDeviceById", device.GetId()).Return(device, nil))
		calls = append(calls, versionRepository.On("FindByNameAndCategory", device.GetTargetVersion().GetName(),
			device.GetTargetVersion().GetCategory()).Return(device.GetTargetVersion(), nil))
		calls = append(calls, repository.On("UpdateVersion", device).Return(device, nil))

		deviceService := usecases.NewDeviceService(repository, versionRepository, deploy)
		createdDevice, err := deviceService.SyncDeviceVersion(device.GetId(), device.GetTargetVersion().GetName())

		assert.Nil(t, err)
		assert.Equal(t, device.GetTargetVersion().GetId(), createdDevice.GetTargetVersion().GetId())

		unsetMocks(calls...)
	})

	t.Run("when the device cannot be found, it should return an error", func(t *testing.T) {
		device := fixture.GetDevices(1)[0]
		calls := make([]*mock.Call, 0, 3)
		calls = append(calls, repository.On("FindDeviceById", device.GetId()).Return(nil, fmt.Errorf("not found")))

		deviceService := usecases.NewDeviceService(repository, versionRepository, deploy)
		createdDevice, err := deviceService.SyncDeviceVersion(device.GetId(), device.GetTargetVersion().GetName())

		assert.Equal(t, fmt.Errorf("not found"), err)
		assert.Nil(t, createdDevice)

		unsetMocks(calls...)
	})
	t.Run("when the version cannot be found, should return an error", func(t *testing.T) {
		device := fixture.GetDevices(1)[0]

		calls := make([]*mock.Call, 0, 3)
		calls = append(calls, repository.On("FindDeviceById", device.GetId()).Return(device, nil))
		calls = append(calls, versionRepository.On("FindByNameAndCategory", device.GetTargetVersion().GetName(),
			device.GetTargetVersion().GetCategory()).Return(nil, fmt.Errorf("not found")))

		deviceService := usecases.NewDeviceService(repository, versionRepository, deploy)
		createdDevice, err := deviceService.SyncDeviceVersion(device.GetId(), device.GetTargetVersion().GetName())

		assert.Equal(t, fmt.Errorf("not found"), err)
		assert.Nil(t, createdDevice)

		unsetMocks(calls...)
	})
	t.Run("when the current version is different from the target version, should send the target version", func(t *testing.T) {
		device := fixture.GetDevices(1)[0]
		version := fixture.GetVersions(1)[0]

		calls := make([]*mock.Call, 0, 3)
		calls = append(calls, repository.On("FindDeviceById", device.GetId()).Return(device, nil))
		calls = append(calls, versionRepository.On("FindByNameAndCategory", version.GetName(),
			device.GetTargetVersion().GetCategory()).Return(version, nil))
		calls = append(calls, deploy.On("SendUpdate", 0.0, device).Return(1, 1, nil))

		deviceService := usecases.NewDeviceService(repository, versionRepository, deploy)
		createdDevice, err := deviceService.SyncDeviceVersion(device.GetId(), version.GetName())

		assert.Nil(t, err)
		assert.NotNil(t, createdDevice)

		unsetMocks(calls...)
	})
	t.Run("when unable to send target version, should return an error", func(t *testing.T) {
		device := fixture.GetDevices(1)[0]
		version := fixture.GetVersions(1)[0]

		calls := make([]*mock.Call, 0, 3)
		calls = append(calls, repository.On("FindDeviceById", device.GetId()).Return(device, nil))
		calls = append(calls, versionRepository.On("FindByNameAndCategory", version.GetName(),
			device.GetTargetVersion().GetCategory()).Return(version, nil))
		calls = append(calls, deploy.On("SendUpdate", 0.0, device).Return(0, 0, fmt.Errorf("update error")))

		deviceService := usecases.NewDeviceService(repository, versionRepository, deploy)
		createdDevice, err := deviceService.SyncDeviceVersion(device.GetId(), version.GetName())

		assert.Equal(t, fmt.Errorf("update error"), err)
		assert.Nil(t, createdDevice)

		unsetMocks(calls...)
	})
	t.Run("when it is not possible to update the version, it should return an error", func(t *testing.T) {
		device := fixture.GetDevices(1)[0]

		calls := make([]*mock.Call, 0, 3)
		calls = append(calls, repository.On("FindDeviceById", device.GetId()).Return(device, nil))
		calls = append(calls, versionRepository.On("FindByNameAndCategory", device.GetTargetVersion().GetName(),
			device.GetTargetVersion().GetCategory()).Return(device.GetTargetVersion(), nil))
		calls = append(calls, deploy.On("SendUpdate", 0.0, device).Return(1, 1, nil))
		calls = append(calls, repository.On("UpdateVersion", device).Return(nil, fmt.Errorf("update error")))

		deviceService := usecases.NewDeviceService(repository, versionRepository, deploy)
		createdDevice, err := deviceService.SyncDeviceVersion(device.GetId(), device.GetTargetVersion().GetName())

		assert.Equal(t, fmt.Errorf("update error"), err)
		assert.Nil(t, createdDevice)

		unsetMocks(calls...)
	})
}

func unsetMocks(calls ...*mock.Call) {
	for _, call := range calls {
		if call != nil {
			call.Unset()
		}
	}
}
