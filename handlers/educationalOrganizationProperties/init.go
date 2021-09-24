package educationalOrganizationProperties

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

type IService interface {
	GetAll() (models.EducationalOrganizationProperties, error)
	UpsertMany(models.EducationalOrganizationProperties) error
	DeleteMany([]string) error
}

type IRepository interface {
	getAll() (models.EducationalOrganizationProperties, error)
	upsertMany(models.EducationalOrganizationProperties) error
	deleteMany([]string) error
}

type Service struct {
	repository IRepository
}

type Repository struct {
	db  *bun.DB
	ctx context.Context
}

func CreateService(db *bun.DB) *Service {
	repo := NewRepository(db)
	return NewService(repo)
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db, ctx: context.Background()}
}
