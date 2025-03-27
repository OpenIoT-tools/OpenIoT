package services

import (
	"github.com/OpenIoT-tools/OpenIoT/consts"
	"github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"
	"github.com/OpenIoT-tools/OpenIoT/internal/core/ports"
)

type DeviceService struct {
	repository        ports.DeviceRespository
	versionRepository ports.VersionRespository
	broker            ports.DeviceBroker
}

func NewDeviceService(repository ports.DeviceRespository, versionRepository ports.VersionRespository, broker ports.DeviceBroker) *DeviceService {
	return &DeviceService{
		repository:        repository,
		broker:            broker,
		versionRepository: versionRepository,
	}
}

func (d *DeviceService) CreateDevice(device *entity.Device) (*entity.Device, error) {
	createdDevice, err := d.repository.CreateDevice(device)
	if err != nil {
		return nil, err
	}
	if err = d.broker.UpdateWithBlueGreen(createdDevice); err != nil {
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

func (d *DeviceService) UpdateVersion(versionId, strategy string, devicesId ...string) ([]*entity.Device, error) {
	devices, err := d.repository.ListDevicesById(devicesId...)
	if err != nil {
		return nil, err
	}
	return d.updateDeviceVersion(versionId, strategy, devices...)
}

func (d *DeviceService) UpdateVersionByCategory(categoryId string, versionId string, strategy string) ([]*entity.Device, error) {
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

func (d *DeviceService) updateDeviceVersion(versionId string, strategy string, devices ...*entity.Device) ([]*entity.Device, error) {
	version, err := d.versionRepository.FindVersionById(versionId)
	if err != nil {
		return nil, err
	}
	for _, device := range devices {
		if _, err := device.UpdateTargetVersion(version); err != nil {
			return nil, err
		}
	}
	_, err = d.repository.UpdateTargetVersion(devices...)
	if err != nil {
		return nil, err
	}

	deploy := d.selectStrategy(strategy)
	err = deploy(devices...)
	if err != nil {
		// TODO: retry
		return nil, err
	}

	return devices, nil
}

func (d *DeviceService) selectStrategy(strategy string) func(device ...*entity.Device) error {
	switch strategy {
	case consts.CANARY_DEPLOY_STRATEGY:
		return d.broker.UpdateWithCanary
	default:
		return d.broker.UpdateWithBlueGreen
	}
}
