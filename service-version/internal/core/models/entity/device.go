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
	if err := device.validHardwareVersion(device.hardwareVersion); err != nil {
		return nil, err
	}

	return device, nil
}

func (d *Device) UpdateTargetVersion(newVersion *Version) (*Device, error) {
	if err := d.validVersion(); err != nil {
		return nil, err
	}
	if err := d.validHardwareVersion(d.hardwareVersion); err != nil {
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

func (d *Device) UpdateHardwareVersion(version float64) (*Device, error) {
	if err := d.validHardwareVersion(version); err != nil {
		return nil, err
	}
	d.hardwareVersion = version
	return d, nil
}

func (d *Device) validVersion() error {
	if d.targetVersion.category.id != d.category.id {
		return fmt.Errorf("device version must be the same category as device")
	}
	return nil
}

func (d *Device) validHardwareVersion(hardwareVersion float64) error {
	if hardwareVersion < d.targetVersion.minimumHardwareVersion || hardwareVersion > d.targetVersion.maximumHardwareVersion {
		return fmt.Errorf(
			"the version is only compatible with the hardware version between %f and %f, but the hardware version is %f",
			d.targetVersion.minimumHardwareVersion, d.targetVersion.maximumHardwareVersion, hardwareVersion)
	}
	return nil
}
