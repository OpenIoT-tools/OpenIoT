package ports

type VersionService interface {
	CreateVersion()
	RemoveVersion()
	ListVersions()
}

type DeviceService interface {
	CreateDevice()
	RemoveDevice()
	ListDevices()
	UpdateDeviceVersion()
	UpdateDeviceMinimunHardware()
}

type CategoryService interface {
	CreateCategory()
	ListCategory()
	RemoveCategory()
}
