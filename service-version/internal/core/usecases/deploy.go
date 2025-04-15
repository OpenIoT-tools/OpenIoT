package usecases

import (
	"fmt"
	"math"
	"time"

	"github.com/OpenIoT-tools/OpenIoT/internal/adapters/dtos/deviceupdate"
	"github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"
	"github.com/OpenIoT-tools/OpenIoT/internal/core/ports"
	"github.com/OpenIoT-tools/OpenIoT/internal/security"
)

type Deploy struct {
	broker        ports.Broker
	securityToken security.SecurityToken
}

func NewDeploy(broker ports.Broker, securityToken security.SecurityToken) *Deploy {
	return &Deploy{
		broker:        broker,
		securityToken: securityToken,
	}
}

// UpdateWithBlueGreen is responsible for updating all devices sent
// This update will be made over the defined time
func (d *Deploy) SendUpdate(hoursLong float64, devices ...*entity.Device) (devicesByGroup int, updateInterval int, err error) {
	if len(devices) < 1 {
		return 0, 0, fmt.Errorf("no devices sent for update")
	}
	devicesByGroup, updateInterval = d.getNumberOfDevicesPerMinute(hoursLong, len(devices))

	date := time.Now().Add(time.Minute * 10)
	for i, device := range devices {
		if i > 0 && i%int(devicesByGroup) == 0 {
			date = date.Add(time.Duration(updateInterval) * time.Minute)
		}
		if err = d.sendUpdate(device, date); err != nil {
			return 0, 0, err
		}
	}
	return devicesByGroup, updateInterval, nil
}

func (d *Deploy) sendUpdate(device *entity.Device, updateTime time.Time) error {
	updateData := map[string]any{
		"update_time":     updateTime.Unix(),
		"hardwareVersion": device.GetHardwateVersion(),
	}
	token, err := d.securityToken.GenerateToken(updateData, 10, "DEVICE_PRIVATE_KEY")
	if err != nil {
		return err
	}

	deviceDTO := deviceupdate.NewSecureDeviceUpdate(token)
	if err := d.broker.SendUpdateToDevice(deviceDTO, device); err != nil {
		return err
	}
	return nil
}

func (d *Deploy) getNumberOfDevicesPerMinute(hoursLong float64, numberOfDevices int) (devices int, minutes int) {
	totalMinutes := math.Ceil(hoursLong * 60)
	if totalMinutes == 0 {
		return numberOfDevices, 1
	} else if numberOfDevices > int(totalMinutes) {
		return int(math.Ceil(float64(numberOfDevices) / totalMinutes)), 1
	} else {
		return 1, int(math.Floor(totalMinutes / float64(numberOfDevices)))
	}
}
