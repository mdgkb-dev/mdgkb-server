package human

import (
	"context"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

type IService interface {
	Create(*models.Human) error
	Update(*models.Human) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.Human) error
	update(*models.Human) error
}

type Handler struct {
	service IService
	helper *helpers.Helper
}

type Service struct {
	repository IRepository
	helper *helpers.Helper
}

type Repository struct {
	db  *bun.DB
	ctx context.Context
	helper *helpers.Helper
}

func CreateService(db *bun.DB, helper *helpers.Helper) *Service {
	repo := NewRepository(db, helper)
	return NewService(repo, helper)
}

func NewService(repository IRepository, helper *helpers.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helpers.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}
