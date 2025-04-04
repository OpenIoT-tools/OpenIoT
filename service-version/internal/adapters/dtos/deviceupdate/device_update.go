package deviceupdate

import "time"

// DeviceUpdate defines the object that will be sent to the device with update information
type DeviceUpdate struct {
	Token           string    `json:"token"`
	UpdateTime      time.Time `json:"update_time"`
	UrlUpdate       string    `json:"url_update"`
	HardwareVersion float64   `json:"hardware_version"`
}

// NewDeviceUpdate creates an object that will transport unprotected information
func NewDeviceUpdate(updateTime time.Time, urlUpdate string, hardwareVersion float64) *DeviceUpdate {
	return &DeviceUpdate{
		UpdateTime:      updateTime,
		UrlUpdate:       urlUpdate,
		HardwareVersion: hardwareVersion,
	}
}

// NewSecureDeviceUpdate creates an object that will carry update information inside a token
func NewSecureDeviceUpdate(token string) *DeviceUpdate {
	return &DeviceUpdate{
		Token: token,
	}
}
