package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Version struct {
	id                     string
	name                   string
	description            string
	category               *Category
	minimumHardwareVersion float64
	maximumHardwareVersion float64
	createdAt              time.Time
}

func NewVersion(name, description string, minimumHardwareVersion, maximumHardwareVersion float64, category *Category) (*Version, error) {
	version := Version{
		id:                     uuid.NewString(),
		name:                   name,
		description:            description,
		category:               category,
		minimumHardwareVersion: minimumHardwareVersion,
		maximumHardwareVersion: maximumHardwareVersion,
		createdAt:              time.Now(),
	}
	if err := version.validVersion(); err != nil {
		return nil, err
	}
	return &version, nil
}

func (v *Version) GetCategory() *Category {
	return v.category
}

func (v *Version) validVersion() error {
	if err := v.validName(); err != nil {
		return err
	}
	if err := v.validDescription(); err != nil {
		return err
	}
	return nil
}

func (v *Version) validName() error {
	if len(v.name) > 30 {
		return fmt.Errorf("version must be less than 50 characters")
	}
	return nil
}

func (v *Version) validDescription() error {
	if len(v.description) > 200 {
		return fmt.Errorf("description must be less than 200 characters")
	}
	return nil
}
