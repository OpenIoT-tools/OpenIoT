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

// NewCategory should be used to create a category. If an invalid name is passed, an error will be returned.
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
	if len(c.name) > 100 {
		return fmt.Errorf("name must be less than 100 characters")
	}
	return nil
}

func (c *Category) GetId() string {
	return c.id
}

func (c *Category) GetName() string {
	return c.id
}
