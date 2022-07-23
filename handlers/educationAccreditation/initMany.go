package educationAccreditation

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/mdgkb-server/models"
)

type IService interface {
	CreateMany(certifications models.EducationAccreditations) error
	UpsertMany(models.EducationAccreditations) error
	DeleteMany([]string) error
}

type IRepository interface {
	createMany(models.EducationAccreditations) error
	upsertMany(models.EducationAccreditations) error
	deleteMany([]string) error
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
