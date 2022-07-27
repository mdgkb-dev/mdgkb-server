package fileinfos

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/pro-assistance/pro-assister/helper"

	"github.com/uptrace/bun"
)

type IService interface {
	Create(info *models.FileInfo) error
	Update(info *models.FileInfo) error
	Upsert(info *models.FileInfo) error
	UpsertMany(infos models.FileInfos) error
}

type IRepository interface {
	db() *bun.DB
	create(info *models.FileInfo) error
	update(info *models.FileInfo) error
	upsert(info *models.FileInfo) error
	upsertMany(infos models.FileInfos) error
	//deleteMany([]string) error
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
