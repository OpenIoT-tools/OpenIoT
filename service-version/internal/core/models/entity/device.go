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
	d := &Device{
		id:              uuid.NewString(),
		hardwareVersion: hardwareVersion,
		category:        category,
		targetVersion:   version,
	}
	if err := d.validVersion(version); err != nil {
		return nil, err
	}
	if err := d.validHardwareVersion(d.hardwareVersion, d.targetVersion.minimumHardwareVersion, d.targetVersion.maximumHardwareVersion); err != nil {
		return nil, err
	}

	return d, nil
}

// UpdateTargetVersion should be used to update the target version of the device
func (d *Device) UpdateTargetVersion(newVersion *Version) (*Device, error) {
	if err := d.validVersion(newVersion); err != nil {
		return nil, err
	}
	if err := d.validHardwareVersion(d.hardwareVersion, newVersion.minimumHardwareVersion, newVersion.maximumHardwareVersion); err != nil {
		return nil, err
	}
	d.targetVersion = newVersion

	return d, nil
}

// Update Current Version should be used to update the current system version
// newVersion should be equal to currentTargetVersion
func (d *Device) UpdateCurrentVersion(newVersion *Version) (*Device, error) {
	if newVersion.id != d.targetVersion.id {
		return nil, fmt.Errorf("current device version must be equal to targetVersion")
	}
	d.currentVersion = d.targetVersion
	return d, nil
}

func (d *Device) UpdateHardwareVersion(version float64) (*Device, error) {
	if err := d.validHardwareVersion(version, d.targetVersion.minimumHardwareVersion, d.targetVersion.maximumHardwareVersion); err != nil {
		return nil, err
	}
	d.hardwareVersion = version
	return d, nil
}

func (d *Device) validVersion(targetVersion *Version) error {
	if targetVersion.category.id != d.category.id {
		return fmt.Errorf("device version must be the same category as device")
	}
	return nil
}

func (d *Device) validHardwareVersion(hardwareVersion, minimumVersion, maximumVersion float64) error {
	if hardwareVersion < minimumVersion || hardwareVersion > maximumVersion {
		return fmt.Errorf(
			"the version is only compatible with the hardware version between %.2f and %f, but the hardware version is %.2f",
			minimumVersion, maximumVersion, hardwareVersion)
	}
	return nil
}

func (d *Device) GetCategory() *Category {
	return d.category
}

func (d *Device) GetHardwateVersion() float64 {
	return d.hardwareVersion
}

func (d *Device) GetTargetVersion() *Version {
	return d.targetVersion
}

func (d *Device) GetCurrentVersion() *Version {
	return d.currentVersion
}

func (d *Device) GetId() string {
	return d.id
}
