package formStatusToFormStatuses

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

type IService interface {
	UpsertMany(models.FormStatusToFormStatuses) error
	DeleteMany([]string) error
}

type IRepository interface {
	db() *bun.DB
	upsertMany(models.FormStatusToFormStatuses) error
	deleteMany([]string) error
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
