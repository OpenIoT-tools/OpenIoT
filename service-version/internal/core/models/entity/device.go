package entity

import "github.com/google/uuid"

type Device struct {
	id              string
	hardwareVersion float64
	category        *Category
	version         *Version
}

func NewDevice(hardwareVersion float64, version *Version, category *Category) (*Device, error) {
	return &Device{
		id:              uuid.NewString(),
		hardwareVersion: hardwareVersion,
		category:        category,
		version:         version,
	}, nil
}
