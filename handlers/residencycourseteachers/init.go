package residencycourseteachers

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/pro-assistance/pro-assister/helper"

	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

type IService interface {
	CreateMany(teachers models.ResidencyCoursesTeachers) error
}

type IRepository interface {
	db() *bun.DB
	createMany(models.ResidencyCoursesTeachers) error
	upsertMany(models.ResidencyCoursesTeachers) error
	deleteMany([]uuid.UUID) error
}

type Handler struct {
	service IService
}

type Service struct {
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
}

func CreateService(h *helper.Helper) *Service {
	repo := NewRepository(h)
	return NewService(repo, h)
}

func NewService(repository IRepository, h *helper.Helper) *Service {
	return &Service{repository: repository, helper: h}
}

func NewRepository(h *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: h}
}
