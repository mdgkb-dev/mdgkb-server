package vacancyRequirements

import (
	"context"
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/mdgkb-server/models"
)

type IService interface {
	CreateMany(models.VacancyRequirements) error
	UpsertMany(models.VacancyRequirements) error
	DeleteMany([]uuid.UUID) error
}

type IRepository interface {
	createMany(models.VacancyRequirements) error
	upsertMany(models.VacancyRequirements) error
	deleteMany([]uuid.UUID) error
}

type Service struct {
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	helper *helper.Helper
	ctx    context.Context
}

func CreateService(h *helper.Helper) *Service {
	repo := NewRepository(h)
	return NewService(repo, h)
}

func NewService(repository IRepository, h *helper.Helper) *Service {
	return &Service{repository: repository}
}

func NewRepository(h *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: h}
}
