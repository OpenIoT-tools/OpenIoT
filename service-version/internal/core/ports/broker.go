package ports

import "github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"

type DeviceBroker interface {
	SendUpdateCommand(device *entity.Device) error
}
