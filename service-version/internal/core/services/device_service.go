package services

import (
	"github.com/OpenIoT-tools/OpenIoT/consts"
	"github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"
	"github.com/OpenIoT-tools/OpenIoT/internal/core/ports"
)

type DeviceService struct {
	repository        ports.DeviceRespository
	versionRepository ports.VersionRespository
	deploy            ports.DeployService
}

func NewDeviceService(repository ports.DeviceRespository, versionRepository ports.VersionRespository, deploy ports.DeployService) *DeviceService {
	return &DeviceService{
		repository:        repository,
		deploy:            deploy,
		versionRepository: versionRepository,
	}
}

func (d *DeviceService) CreateDevice(device *entity.Device) (*entity.Device, error) {
	createdDevice, err := d.repository.CreateDevice(device)
	if err != nil {
		return nil, err
	}
	if err = d.deploy.UpdateWithBlueGreen(createdDevice); err != nil {
		return nil, err
	}

	return createdDevice, nil
}

func (d *DeviceService) RemoveDevice(deviceId string) error {
	return d.repository.RemoveDevice(deviceId)
}

func (d *DeviceService) ListDevices(categoryId string) ([]*entity.Device, error) {
	return d.repository.ListDevicesByCategory(categoryId)
}

func (d *DeviceService) UpdateTargetVersion(versionId, strategy string, devicesId ...string) ([]*entity.Device, error) {
	devices, err := d.repository.ListDevicesById(devicesId...)
	if err != nil {
		return nil, err
	}
	return d.updateDeviceVersion(versionId, strategy, devices...)
}

func (d *DeviceService) UpdateTargetVersionByCategory(categoryId string, versionId string, strategy string) ([]*entity.Device, error) {
	devices, err := d.repository.ListDevicesByCategory(categoryId)
	if err != nil {
		return nil, err
	}
	return d.updateDeviceVersion(versionId, strategy, devices...)
}

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
		if err := d.deploy.UpdateWithBlueGreen(device); err != nil {
			return nil, err
		}
		return device, nil
	}
	return d.repository.UpdateVersion(device)
}

func (d *DeviceService) updateDeviceVersion(versionId string, strategy string, devices ...*entity.Device) ([]*entity.Device, error) {
	devices, err := d.setVersionOnDevices(versionId, devices...)
	if err != nil {
		return nil, err
	}

	if _, err = d.repository.UpdateTargetVersion(devices...); err != nil {
		return nil, err
	}

	deploy := d.selectStrategy(strategy)
	if err = deploy(devices...); err != nil {
		// TODO: retry
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

func (d *DeviceService) selectStrategy(strategy string) func(device ...*entity.Device) error {
	switch strategy {
	case consts.CANARY_DEPLOY_STRATEGY:
		return d.deploy.UpdateWithCanary
	default:
		return d.deploy.UpdateWithBlueGreen
	}
}
