package educationalOrganizationPages

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

type IService interface {
	GetAll() (models.EducationalOrganizationPages, error)
	CreateMany(models.EducationalOrganizationPages) error
	UpsertMany(models.EducationalOrganizationPages) error
	DeleteMany([]string) error
}

type IRepository interface {
	getDB() *bun.DB
	getAll() (models.EducationalOrganizationPages, error)
	createMany(models.EducationalOrganizationPages) error
	upsertMany(models.EducationalOrganizationPages) error
	deleteMany([]string) error
}

type Handler struct {
	service IService
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
