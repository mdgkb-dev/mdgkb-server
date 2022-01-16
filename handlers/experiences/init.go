package experiences

import (
	"context"
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

type IService interface {
	CreateMany(models.Experiences) error
}

type IRepository interface {
	getDB() *bun.DB
	createMany(models.Experiences) error
	upsertMany(models.Experiences) error
	deleteMany([]uuid.UUID) error
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
