package usecases

import (
	"time"

	"github.com/OpenIoT-tools/OpenIoT/internal/adapters/dtos/deviceupdate"
	"github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"
	"github.com/OpenIoT-tools/OpenIoT/internal/core/ports"
	"github.com/OpenIoT-tools/OpenIoT/internal/security"
)

type Deploy struct {
	broker ports.Broker
}

func NewDeploy(broker ports.Broker) *Deploy {
	return &Deploy{
		broker: broker,
	}
}

// UpdateWithBlueGreen is responsible for updating all devices sent
// This update will be made over the defined time
func (d *Deploy) SendUpdate(hoursLong uint32, devices ...*entity.Device) error {
	devicesByGroup, updateInterval := d.getNumberOfDevicesPerMinute(hoursLong, uint32(len(devices)))

	date := time.Now().Add(time.Minute * 10)
	for i, device := range devices {
		if i > 0 && i%int(devicesByGroup) == 0 {
			date = date.Add(time.Duration(updateInterval) * time.Minute)
		}
		if err := d.sendUpdate(device, date); err != nil {
			return err
		}
	}
	return nil
}

func (d *Deploy) sendUpdate(device *entity.Device, updateTime time.Time) error {
	updateData := map[string]any{
		"update_time":     updateTime.Unix(),
		"hardwareVersion": device.GetHardwateVersion(),
	}
	token, err := security.GenerateToken(updateData, 10, "DEVICE_PRIVATE_KEY")
	if err != nil {
		return err
	}

	deviceDTO := deviceupdate.NewSecureDeviceUpdate(token)
	if err := d.broker.SendUpdateToDevice(deviceDTO, device); err != nil {
		return err
	}
	return nil
}

func (d *Deploy) getNumberOfDevicesPerMinute(hoursLong uint32, numberOfDevices uint32) (uint32, uint32) {
	totalMinutes := hoursLong * 60
	if totalMinutes == 0 {
		return numberOfDevices, 1
	} else if numberOfDevices > totalMinutes {
		return numberOfDevices / totalMinutes, 1
	} else {
		return 1, totalMinutes / numberOfDevices
	}
}
