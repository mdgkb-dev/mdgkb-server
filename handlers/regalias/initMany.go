package regalias

import (
	"context"
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/mdgkb-server/models"
)

type IService interface {
	CreateMany(models.Regalias) error
	UpsertMany(models.Regalias) error
	DeleteMany([]uuid.UUID) error
}

type IRepository interface {
	createMany(models.Regalias) error
	upsertMany(models.Regalias) error
	deleteMany([]uuid.UUID) error
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
	return &Service{repository: repository}
}

func NewRepository(h *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: h}
}
