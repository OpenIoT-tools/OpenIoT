package usecases

import (
	"github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"
	"github.com/OpenIoT-tools/OpenIoT/internal/core/ports"
)

type DeviceService struct {
	repository        ports.DeviceRespository
	versionRepository ports.VersionRespository
	deploy            ports.DeployUseCase
}

func NewDeviceService(repository ports.DeviceRespository, versionRepository ports.VersionRespository, deploy ports.DeployUseCase) *DeviceService {
	return &DeviceService{
		repository:        repository,
		deploy:            deploy,
		versionRepository: versionRepository,
	}
}

// Find Device Version should be used to get the device version
func (d *DeviceService) FindDeviceVersion(deviceId string) (*entity.Version, error) {
	device, err := d.repository.FindDeviceById(deviceId)
	if err != nil {
		return nil, err
	}
	return device.GetTargetVersion(), nil
}

// CreateDevice shoud be user for create an device
func (d *DeviceService) CreateDevice(device *entity.Device) (*entity.Device, error) {
	createdDevice, err := d.repository.CreateDevice(device)
	if err != nil {
		return nil, err
	}
	if err = d.deploy.SendUpdate(0, createdDevice); err != nil {
		return nil, err
	}

	return createdDevice, nil
}

// RemoveDevice shoud be user for remove an device
func (d *DeviceService) RemoveDevice(deviceId string) error {
	return d.repository.RemoveDevice(deviceId)
}

// ListDevices should be used to list all devices in the category
func (d *DeviceService) ListDevices(categoryId string) ([]*entity.Device, error) {
	return d.repository.ListDevicesByCategory(categoryId)
}

// UpdateTargetVersion should be used for both update and device group
func (d *DeviceService) UpdateTargetVersion(versionId string, updateDurationHours uint32, devicesId ...string) ([]*entity.Device, error) {
	devices, err := d.repository.ListDevicesById(devicesId...)
	if err != nil {
		return nil, err
	}
	return d.updateDeviceVersion(versionId, updateDurationHours, devices...)
}

// UpdateTargetVersionByCategory should be used to update all devices in the category
func (d *DeviceService) UpdateTargetVersionByCategory(categoryId string, versionId string, updateDurationHours uint32) ([]*entity.Device, error) {
	devices, err := d.repository.ListDevicesByCategory(categoryId)
	if err != nil {
		return nil, err
	}
	return d.updateDeviceVersion(versionId, updateDurationHours, devices...)
}

// UpdateHardware should be user to update device hardware
func (d *DeviceService) UpdateHardware(deviceId string, hardware float64) (*entity.Device, error) {
	device, err := d.repository.FindDeviceById(deviceId)
	if err != nil {
		return nil, err
	}

	if _, err = device.UpdateHardwareVersion(hardware); err != nil {
		return nil, err
	}
	return d.repository.UpdateHardware(device)
}

// SyncDeviceVersion should be used to ensure the device runs the same version as defined in the system
func (d *DeviceService) SyncDeviceVersion(deviceId, versionName string) (*entity.Device, error) {
	device, err := d.repository.FindDeviceById(deviceId)
	if err != nil {
		return nil, err
	}
	version, err := d.versionRepository.FindByNameAndCategory(versionName, device.GetCategory())
	if err != nil {
		return nil, err
	}

	device, err = device.UpdateCurrentVersion(version)
	if err != nil {
		if err := d.deploy.SendUpdate(0, device); err != nil {
			return nil, err
		}
		return device, nil
	}
	return d.repository.UpdateVersion(device)
}

func (d *DeviceService) updateDeviceVersion(versionId string, updateDurationHours uint32, devices ...*entity.Device) ([]*entity.Device, error) {
	devices, err := d.setVersionOnDevices(versionId, devices...)
	if err != nil {
		return nil, err
	}
	if _, err = d.repository.UpdateTargetVersion(devices...); err != nil {
		return nil, err
	}

	if err = d.deploy.SendUpdate(updateDurationHours, devices...); err != nil {
		return nil, err
	}
	return devices, nil
}

func (d *DeviceService) setVersionOnDevices(versionId string, devices ...*entity.Device) ([]*entity.Device, error) {
	version, err := d.versionRepository.FindById(versionId)
	if err != nil {
		return nil, err
	}
	for _, device := range devices {
		if _, err := device.UpdateTargetVersion(version); err != nil {
			return nil, err
		}
	}
	return devices, err
}
