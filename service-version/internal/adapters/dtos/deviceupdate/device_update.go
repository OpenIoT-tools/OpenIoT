package deviceupdate

import "time"

// DeviceUpdate defines the object that will be sent to the device with update information
type DeviceUpdate struct {
	Token           string    `json:"token"`
	UpdateTime      time.Time `json:"update_time"`
	HardwareVersion float64   `json:"hardware_version"`
}

// NewDeviceUpdate creates an object that will transport unprotected information
func NewDeviceUpdate(updateTime time.Time, hardwareVersion float64) *DeviceUpdate {
	return &DeviceUpdate{
		UpdateTime:      updateTime,
		HardwareVersion: hardwareVersion,
	}
}

// NewSecureDeviceUpdate creates an object that will carry update information inside a token
func NewSecureDeviceUpdate(token string) *DeviceUpdate {
	return &DeviceUpdate{
		Token: token,
	}
}
