package ports

import "github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"

type VersionService interface {
	CreateVersion(version *entity.Version) (*entity.Version, error)
	RemoveVersion(versionId string) error
	ListVersions(categoryId string) (*[]entity.Version, error)
}

type DeviceService interface {
	CreateDevice(device *entity.Device) (*entity.Device, error)
	RemoveDevice(deviceId string) error
	ListDevices(categoryId string) (*[]entity.Device, error)
	UpdateVersion(versionId, deviceId string) (*entity.Device, error)
	UpdateVersionByCategory(categoryId string, versionId string, strategy string)
	UpdateHardware(deviceId string, minimunHardware float64) (*entity.Device, error)
}

type CategoryService interface {
	CreateCategory(category *entity.Category) (*entity.Category, error)
	ListCategories() (*[]entity.Category, error)
	RemoveCategory(categoryId string) error
}
