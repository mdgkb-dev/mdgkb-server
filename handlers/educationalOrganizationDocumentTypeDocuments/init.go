package educationalOrganizationDocumentTypeDocuments

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

type IService interface {
	CreateMany(models.EducationalOrganizationDocumentTypeDocuments) error
	UpsertMany(models.EducationalOrganizationDocumentTypeDocuments) error
	DeleteMany([]string) error
}

type IRepository interface {
	getDB() *bun.DB
	createMany(models.EducationalOrganizationDocumentTypeDocuments) error
	upsertMany(models.EducationalOrganizationDocumentTypeDocuments) error
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
