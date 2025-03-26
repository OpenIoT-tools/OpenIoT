package services

import (
	"github.com/OpenIoT-tools/OpenIoT/internal/adapters/dtos"
	"github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"
	"github.com/OpenIoT-tools/OpenIoT/internal/core/ports"
)

type DeviceService struct {
	repository ports.DeviceRespository
	broker     ports.DeviceBroker
}

func NewDeviceService(repository *ports.DeviceRespository, broker *ports.DeviceBroker) *DeviceService {
	return &DeviceService{
		repository: *repository,
		broker:     *broker,
	}
}

func (d *DeviceService) CreateDevice(device *entity.Device) (*entity.Device, error) {
	createdDevice, err := d.repository.CreateDevice(device)
	if err != nil {
		return nil, err
	}
	if err = d.broker.SendUpdateCommand(createdDevice); err != nil {
		return nil, err
	}

	return createdDevice, nil
}

func (d *DeviceService) RemoveDevice(deviceId string) error {
	return d.repository.RemoveDevice(deviceId)
}

func (d *DeviceService) ListDevices(categoryId string) (dtos.PaginationResponse[entity.Device], error) {
	return d.repository.ListDevicesByCategory(categoryId)
}

func (d *DeviceService) UpdateDeviceVersion(versionId, deviceId string) (*entity.Device, error) {
	createdDevice, err := d.repository.UpdateDeviceVersion(device)
	if err != nil {
		return nil, err
	}
	if err = d.broker.SendUpdateCommand(createdDevice); err != nil {
		return nil, err
	}

	return createdDevice, nil
}

func (d *DeviceService) updateDevices(categoryId string, versionId string, strategy string) {

}

func (d *DeviceService) UpdateDeviceMinimunHardware(deviceId string, minimunHardware float64) (*entity.Device, error) {

}
