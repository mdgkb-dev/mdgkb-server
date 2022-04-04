package formValues

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/pro-assistance/pro-assister/helper"
	"github.com/uptrace/bun"
)

type IService interface {
	Upsert(info *models.FormValue) error
}

type IRepository interface {
	getDB() *bun.DB
	upsert(info *models.FormValue) error
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
