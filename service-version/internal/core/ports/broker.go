package ports

import "github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"

type DeviceBroker interface {
	UpdateWithCanary(device ...*entity.Device) error
	UpdateWithBlueGreen(device ...*entity.Device) error
}
