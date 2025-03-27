package ports

import (
	"github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"
)

type VersionRespository interface {
	CreateVersion(version *entity.Version) (*entity.Version, error)
	RemoveVersion(versionId string) error
	ListVersionsByCategory(categoryId string) (*[]entity.Version, error)
	FindVersionById(versionId string) (*entity.Version, error)
}

type DeviceRespository interface {
	CreateDevice(device *entity.Device) (*entity.Device, error)
	RemoveDevice(deviceId string) error
	ListDevicesByCategory(categoryId string) ([]*entity.Device, error)
	ListDevicesById(devicesId ...string) ([]*entity.Device, error)
	FindDeviceById(deviceId string) (*entity.Device, error)
	UpdateVersion(device entity.Device) (*entity.Device, error)
	UpdateTargetVersion(device ...*entity.Device) ([]*entity.Device, error)
	UpdateHardware(device *entity.Device) (*entity.Device, error)
}

type CategoryRespository interface {
	CreateCategory(category *entity.Category) (*entity.Category, error)
	ListCategory() (*[]entity.Category, error)
	RemoveCategory(categoryId string) error
}
