package ports

import "github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"

type VersionUseCase interface {
	CreateVersion(version *entity.Version) (*entity.Version, error)
	RemoveVersion(versionId string) error
	ListVersions(categoryId string) (*[]entity.Version, error)
}

type DeviceUseCase interface {
	CreateDevice(device *entity.Device) (*entity.Device, error)
	RemoveDevice(deviceId string) error
	ListDevices(categoryId string) (*[]entity.Device, error)
	UpdateTargetVersion(versionId, updateDuration uint32, devicesId ...string) (*entity.Device, error)
	UpdateTargetVersionByCategory(categoryId string, versionId string, updateDuration uint32)
	UpdateHardware(deviceId string, minimunHardware float64) (*entity.Device, error)
}

type CategoryUseCase interface {
	CreateCategory(category *entity.Category) (*entity.Category, error)
	ListCategories() (*[]entity.Category, error)
	RemoveCategory(categoryId string) error
}

type DeployUseCase interface {
	SendUpdate(hoursLong uint32, device ...*entity.Device) error
}
