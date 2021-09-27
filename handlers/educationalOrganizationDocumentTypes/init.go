package educationalOrganizationDocumentTypes

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

type IService interface {
	GetAll() (models.EducationalOrganizationDocumentTypes, error)
	CreateMany(models.EducationalOrganizationDocumentTypes) error
	UpsertMany(models.EducationalOrganizationDocumentTypes) error
	DeleteMany([]string) error
}

type IRepository interface {
	getDB() *bun.DB
	getAll() (models.EducationalOrganizationDocumentTypes, error)
	createMany(models.EducationalOrganizationDocumentTypes) error
	upsertMany(models.EducationalOrganizationDocumentTypes) error
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
