package fixture

import (
	"fmt"

	"github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"
)

func GetDevices(cont int) []*entity.Device {
	version := GetVersions(1)[0]
	devices := make([]*entity.Device, 0, cont)

	for i := 0; i < cont; i++ {
		device, _ := entity.NewDevice(1.0, version, version.GetCategory())
		devices = append(devices, device)
	}
	return devices
}

func GetVersions(cont int) []*entity.Version {
	versions := make([]*entity.Version, 0, cont)

	for i := 0; i < cont; i++ {
		version, _ := entity.NewVersion(fmt.Sprintf("test version %d", i), "description", 1.0, 2.0, GetCategories(1)[0])
		versions = append(versions, version)
	}
	return versions
}

func GetCategories(cont int) []*entity.Category {
	categories := make([]*entity.Category, 0, cont)

	for i := 0; i < cont; i++ {
		category, _ := entity.NewCategory(fmt.Sprintf("test category %d", i))
		categories = append(categories, category)
	}
	return categories
}
