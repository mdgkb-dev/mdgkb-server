package human

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

type IService interface {
	Create(*models.Human) error
	CreateMany(models.Humans) error
	Update(*models.Human) error
	UpsertMany(models.Humans) error
	Upsert(*models.Human) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.Human) error
	createMany(models.Humans) error
	update(*models.Human) error
	upsertMany(models.Humans) error
	upsert(*models.Human) error
	getAllBySlug(string) (models.Humans, error)
}

type Handler struct {
	service IService
	helper  *helper.Helper
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

func CreateService(db *bun.DB, helper *helper.Helper) *Service {
	repo := NewRepository(db, helper)
	return NewService(repo, helper)
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helper.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}
