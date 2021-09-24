package educationalOrganizationTeachers

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

type IService interface {
	GetAll() (models.EducationalOrganizationTeachers, error)
	UpsertMany(models.EducationalOrganizationTeachers) error
	DeleteMany([]string) error
}

type IRepository interface {
	getAll() (models.EducationalOrganizationTeachers, error)
	upsertMany(models.EducationalOrganizationTeachers) error
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
