package contactInfo

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

type IService interface {
	Create(*models.ContactInfo) error
	CreateMany(models.ContactInfos) error
	Update(*models.ContactInfo) error
	Upsert(*models.ContactInfo) error
	UpsertMany(*models.ContactInfo) error
	Delete(*string) error
}

type IRepository interface {
	db() *bun.DB
	create(*models.ContactInfo) error
	createMany(models.ContactInfos) error
	update(*models.ContactInfo) error
	upsert(*models.ContactInfo) error
	upsertMany(models.ContactInfos) error
	delete(*string) error
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
