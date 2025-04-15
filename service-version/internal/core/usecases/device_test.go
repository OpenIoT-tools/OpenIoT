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
	// when cannot list devices, should return an error
	// when version is not found, should return an error
	// when is not possible update target version, should return an error
	// when cannot storage updated device, should return an error
	// when cannot send updating, should return an error
	// when valid device is sent, should return an error
}

func unsetMocks(calls ...*mock.Call) {
	for _, call := range calls {
		if call != nil {
			call.Unset()
		}
	}
}
