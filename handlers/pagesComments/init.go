package pagesComments

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

type IService interface {
	CreateMany(pageDocuments models.PageComments) error
	UpsertMany(pageDocuments models.PageComments) error
	DeleteMany([]string) error
}

type IRepository interface {
	getDB() *bun.DB
	createMany(models.PageComments) error
	upsertMany(models.PageComments) error
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
	db     *bun.DB
	ctx    context.Context
	helper *helper.Helper
}

func CreateService(db *bun.DB, h *helper.Helper) *Service {
	repo := NewRepository(db, h)
	return NewService(repo, h)
}

func NewService(repository IRepository, h *helper.Helper) *Service {
	return &Service{repository: repository, helper: h}
}

func NewRepository(db *bun.DB, h *helper.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: h}
}
