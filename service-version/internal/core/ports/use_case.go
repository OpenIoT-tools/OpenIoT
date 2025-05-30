package ports

import "github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"

type VersionUseCase interface {
	CreateVersion(version *entity.Version) (*entity.Version, error)
	RemoveVersion(versionId string) error
	ListVersions(categoryId string) ([]*entity.Version, error)
}

type DeviceUseCase interface {
	FindDeviceVersion(id string) (*entity.Version, error)
	FindDeviceById(id string) (*entity.Device, error)
	CreateDevice(device *entity.Device) (*entity.Device, error)
	RemoveDevice(id string) error
	ListDevices(categoryId string) ([]*entity.Device, error)
	UpdateTargetVersion(versionId string, updateDurationHours float64, devicesId ...string) ([]*entity.Device, error)
	UpdateTargetVersionByCategory(categoryId string, versionId string, updateDuration float64) ([]*entity.Device, error)
	UpdateHardware(deviceId string, minimunHardware float64) (*entity.Device, error)
	SyncDeviceVersion(deviceId, versionName string) (*entity.Device, error)
}

type CategoryUseCase interface {
	CreateCategory(category *entity.Category) (*entity.Category, error)
	ListCategories() ([]*entity.Category, error)
	RemoveCategory(categoryId string) error
}

type DeployUseCase interface {
	SendUpdate(hoursLong float64, devices ...*entity.Device) (devicesByGroup int, updateInterval int, err error)
}
