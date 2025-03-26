package entity

import (
	"fmt"

	"github.com/google/uuid"
)

type Device struct {
	id              string
	hardwareVersion float64
	category        *Category
	currentVersion  *Version
	targetVersion   *Version
}

func NewDevice(hardwareVersion float64, version *Version, category *Category) (*Device, error) {
	device := &Device{
		id:              uuid.NewString(),
		hardwareVersion: hardwareVersion,
		category:        category,
		targetVersion:   version,
	}
	if err := device.validVersion(); err != nil {
		return nil, err
	}
	if err := device.validHardwareVersion(); err != nil {
		return nil, err
	}

	return device, nil
}

func (d *Device) UpdateTargetVersion(newVersion *Version) (*Device, error) {
	if err := d.validVersion(); err != nil {
		return nil, err
	}
	if err := d.validHardwareVersion(); err != nil {
		return nil, err
	}
	d.targetVersion = newVersion

	return d, nil
}

func (d *Device) UpdateCurrentVersion(newVersion *Version) (*Device, error) {
	if newVersion.id != d.targetVersion.id {
		return nil, fmt.Errorf("current device version must be equal to targetVersion")
	}
	d.currentVersion = d.targetVersion
	return d, nil
}

func (d *Device) validVersion() error {
	if d.targetVersion.category.id != d.category.id {
		return fmt.Errorf("device version must be the same category as device")
	}
	return nil
}

func (d *Device) validHardwareVersion() error {
	if d.targetVersion.minimumHardwareVersion > d.hardwareVersion {
		return fmt.Errorf("this version is only compatible with devices with versions later than 5.0")
	}
	return nil
}
