package services

import (
	"github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"
	"github.com/OpenIoT-tools/OpenIoT/internal/core/ports"
)

type CategoryService struct {
	repository ports.CategoryRespository
}

func NewCategory(repository *ports.CategoryRespository) *CategoryService {
	return &CategoryService{
		repository: *repository,
	}
}

func (c *CategoryService) CreateCategory(category *entity.Category) (*entity.Category, error) {
	return c.repository.CreateCategory(category)
}

func (c *CategoryService) ListCateogories() (*[]entity.Category, error) {
	return c.repository.ListCategory()
}

func (c *CategoryService) RemoveCategory(categoryId string) error {
	return c.repository.RemoveCategory(categoryId)
}
