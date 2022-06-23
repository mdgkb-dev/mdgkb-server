package residencyApplicationsPointsAchievements

import (
	"context"
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

type IService interface {
	CreateMany(item models.ResidencyApplicationPointsAchievements) error
}

type IRepository interface {
	getDB() *bun.DB
	createMany(models.ResidencyApplicationPointsAchievements) error
	upsertMany(models.ResidencyApplicationPointsAchievements) error
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
