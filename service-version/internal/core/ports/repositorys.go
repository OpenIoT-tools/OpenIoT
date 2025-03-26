package ports

import (
	"github.com/OpenIoT-tools/OpenIoT/internal/adapters/dtos"
	"github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"
)

type VersionRespository interface {
	CreateVersion(version *entity.Version) (*entity.Version, error)
	RemoveVersion(versionId string) error
	ListVersionsByCategory(categoryId string) (*[]entity.Version, error)
	findVersionById(versionId string) (*entity.Version, error)
}

type DeviceRespository interface {
	CreateDevice(device *entity.Device) (*entity.Device, error)
	RemoveDevice(deviceId string) error
	ListDevicesByCategory(categoryId string) (dtos.PaginationResponse[entity.Device], error)
	findDeviceById(deviceId string) (*entity.Device, error)
	UpdateDeviceVersion(version entity.Version, device entity.Device) (*entity.Device, error)
	UpdateDeviceMinimunHardware(deviceId string, minimunHardware float64) (*entity.Device, error)
}

type CategoryRespository interface {
	CreateCategory(category *entity.Category) (*entity.Category, error)
	ListCategory() (*[]entity.Category, error)
	RemoveCategory(categoryId string) error
}
