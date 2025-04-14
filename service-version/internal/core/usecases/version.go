package usecases

import (
	"github.com/OpenIoT-tools/OpenIoT/internal/core/models/entity"
	"github.com/OpenIoT-tools/OpenIoT/internal/core/ports"
)

type VersionUseCase struct {
	repository ports.VersionRespository
}

func NewVersionUseCase(repository ports.VersionRespository) *VersionUseCase {
	return &VersionUseCase{
		repository: repository,
	}
}

func (v *VersionUseCase) CreateVersion(version *entity.Version) (*entity.Version, error) {
	return v.repository.Create(version)
}

func (v *VersionUseCase) RemoveVersion(versionId string) error {
	return v.repository.Remove(versionId)
}

func (v *VersionUseCase) ListVersions(categoryId string) (*[]entity.Version, error) {
	return v.repository.ListByCategory(categoryId)
}
