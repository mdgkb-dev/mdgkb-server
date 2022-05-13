package hospitalizationsToDocumentTypes

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

type IService interface {
	CreateMany(items models.HospitalizationsToDocumentTypes) error
}

type IRepository interface {
	getDB() *bun.DB
	createMany(items models.HospitalizationsToDocumentTypes) error
	upsertMany(items models.HospitalizationsToDocumentTypes) error
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
