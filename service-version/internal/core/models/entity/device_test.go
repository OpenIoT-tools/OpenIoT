package entity_test

import (
	"testing"

	"github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"
	"github.com/OpenIoT-tools/OpenIoT/tests/fixture"
	"github.com/stretchr/testify/assert"
)

func TestNewDevice(t *testing.T) {
	t.Run("When a valid device is sent, it should return the device", func(t *testing.T) {
		version := fixture.GetVersions(1)[0]
		category := version.GetCategory()

		device, err := entity.NewDevice(1.5, version, category)
		assert.Nil(t, err)
		assert.NotNil(t, device)
	})
	t.Run("when version category and device category are different, it should return an error", func(t *testing.T) {
		version := fixture.GetVersions(1)[0]
		category := fixture.GetCategories(1)[0]

		device, err := entity.NewDevice(1.5, version, category)

		assert.Nil(t, device)
		assert.NotNil(t, err)
	})
	t.Run("when the version is incompatible with the hardware, it should return an error", func(t *testing.T) {
		version := fixture.GetVersions(1)[0]
		category := version.GetCategory()

		device, err := entity.NewDevice(3.5, version, category)
		assert.NotNil(t, err)
		assert.Nil(t, device)
	})
	t.Run("when invalid hardware is sent, it should return an error", func(t *testing.T) {
		version := fixture.GetVersions(1)[0]
		category := version.GetCategory()

		device, err := entity.NewDevice(-1, version, category)
		assert.Nil(t, device)
		assert.NotNil(t, err)
	})
}

func TestUpdateTargetVersion(t *testing.T) {
	t.Run("when the version category is incompatible, it should return an error", func(t *testing.T) {
		version := fixture.GetVersions(1)[0]
		device := fixture.GetDevices(1)[0]
		updatedDevice, err := device.UpdateTargetVersion(version)

		assert.Nil(t, updatedDevice)
		assert.NotNil(t, err)
	})
	t.Run("when the hardware version is incompatible, it should return an error", func(t *testing.T) {
		device := fixture.GetDevices(1)[0]
		version, _ := entity.NewVersion("", "", 5, 7, device.GetCategory())

		updatedDevice, err := device.UpdateTargetVersion(version)

		assert.Nil(t, updatedDevice)
		assert.NotNil(t, err)
	})
	t.Run("when a valid version is submitted, it should update the target version", func(t *testing.T) {
		device := fixture.GetDevices(1)[0]
		version, _ := entity.NewVersion("", "", 1, 3, device.GetCategory())

		updatedDevice, err := device.UpdateTargetVersion(version)

		assert.Nil(t, err)
		assert.Equal(t, version, updatedDevice.GetTargetVersion())
	})
}

func TestUpdateCurrentVersion(t *testing.T) {
	t.Run("when the sent version is different from the target version, it should return an error", func(t *testing.T) {
		device := fixture.GetDevices(1)[0]
		version, _ := entity.NewVersion("", "", 1, 3, device.GetCategory())

		updatedDevice, err := device.UpdateCurrentVersion(version)

		assert.Nil(t, updatedDevice)
		assert.NotNil(t, version, err)
	})

	t.Run("when a valid version is submitted, shouldUpdate current version", func(t *testing.T) {
		device := fixture.GetDevices(1)[0]
		version, _ := entity.NewVersion("", "", 1, 3, device.GetCategory())

		_, err := device.UpdateTargetVersion(version)
		assert.Nil(t, err)

		updatedDevice, err := device.UpdateCurrentVersion(version)

		assert.Nil(t, err)
		assert.Equal(t, updatedDevice.GetTargetVersion(), updatedDevice.GetCurrentVersion())
	})
}

func TestUpdateHardwareVersion(t *testing.T) {
	t.Run("when the hardware version is incompatible with the target version, it should return an error", func(t *testing.T) {
		device := fixture.GetDevices(1)[0]
		device, err := device.UpdateHardwareVersion(10)

		assert.Nil(t, device)
		assert.NotNil(t, err)
	})
	t.Run("when the hardware version is valid, you must update the hardware version", func(t *testing.T) {
		device := fixture.GetDevices(1)[0]
		device, err := device.UpdateHardwareVersion(1.8)

		assert.Nil(t, err)
		assert.NotNil(t, device)
	})
}
