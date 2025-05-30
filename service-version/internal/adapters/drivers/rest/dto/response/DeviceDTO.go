package response

import (
	"github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"
)

type deviceDTO struct {
	id              string
	hardwareVersion float64
	category        string
	currentVersion  string
}

func NewDevice(device *entity.Device) *deviceDTO {
	return &deviceDTO{
		id:              device.GetId(),
		hardwareVersion: device.GetHardwateVersion(),
		category:        device.GetCategory().GetName(),
		currentVersion:  device.GetCurrentVersion().GetName(),
	}
}
