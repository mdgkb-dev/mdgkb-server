package preparationRulesGroups

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

type IService interface {
	CreateMany(services models.PreparationRulesGroups) error
	UpsertMany(models.PreparationRulesGroups) error
	DeleteMany([]uuid.UUID) error
}

type IRepository interface {
	getDB() *bun.DB
	createMany(models.PreparationRulesGroups) error
	upsertMany(models.PreparationRulesGroups) error
	deleteMany([]uuid.UUID) error
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
