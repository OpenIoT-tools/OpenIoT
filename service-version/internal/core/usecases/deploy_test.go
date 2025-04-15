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

func TestDeploy_SendUpdate(t *testing.T) {
	broker := mocks.NewBroker(t)
	security := mocks.NewSecurityToken(t)
	broker.On("SendUpdateToDevice", mock.Anything, mock.Anything).Return(nil)
	security.On("GenerateToken", mock.Anything, mock.Anything, mock.Anything).Return("token", nil)

	type testData struct {
		name                             string
		updateTime                       float64
		devices                          []*entity.Device
		expectedDevices, expectedMinutes int
		expectedErr                      error
		broker                           *mocks.Broker
		security                         *mocks.SecurityToken
	}

	testCases := []testData{
		{
			name:            "when 2 hours and 10 devices are sent, it should update 1 device every 12 minutes",
			updateTime:      2,
			devices:         fixture.GetDevices(10),
			expectedDevices: 1,
			expectedMinutes: 12,
			expectedErr:     nil,
			broker:          broker,
			security:        security,
		},
		{
			name:            "when 0.5 hours and 10 devices are sent, it should update 1 device every 3 minutes",
			updateTime:      0.5,
			devices:         fixture.GetDevices(10),
			expectedDevices: 1,
			expectedMinutes: 3,
			expectedErr:     nil,
			broker:          broker,
			security:        security,
		},
		{
			name:            "when 0.5 hours and 100 devices are sent, it should update 4 devices per minute",
			updateTime:      0.5,
			devices:         fixture.GetDevices(100),
			expectedDevices: 4,
			expectedMinutes: 1,
			expectedErr:     nil,
			broker:          broker,
			security:        security,
		},
		{
			name:            "when the hours are 0, should update all devices at the same time",
			updateTime:      2,
			devices:         fixture.GetDevices(10),
			expectedDevices: 1,
			expectedMinutes: 12,
			expectedErr:     nil,
			broker:          broker,
			security:        security,
		},
		{
			name:            "when devices are not sent, it should return an error",
			updateTime:      2,
			devices:         fixture.GetDevices(0),
			expectedDevices: 0,
			expectedMinutes: 0,
			expectedErr:     fmt.Errorf("no devices sent for update"),
			broker:          broker,
			security:        security,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			deploy := usecases.NewDeploy(broker, security)
			devices, updateTime, err := deploy.SendUpdate(test.updateTime, test.devices...)

			assert.Equal(t, test.expectedDevices, devices)
			assert.Equal(t, test.expectedMinutes, updateTime)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}
