package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Category struct {
	id        string
	name      string
	createdAt time.Time
	updatedAt time.Time
}

func NewCategory(name string) (*Category, error) {
	category := &Category{
		id:        uuid.NewString(),
		name:      name,
		createdAt: time.Now(),
	}
	if err := category.validName(); err != nil {
		return nil, err
	}
	return category, nil
}

func (c *Category) UpdateName(name string) (*Category, error) {
	if err := c.validName(); err != nil {
		return nil, err
	}
	c.name = name
	c.updatedAt = time.Now()
	return c, nil
}

func (c *Category) validName() error {
	if len(c.name) > 30 {
		return fmt.Errorf("name must be less than 50 characters")
	}
	return nil
}

func (c *Category) GetId() string {
	return c.id
}
