package maproutenodes

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"

	"github.com/uptrace/bun"
)

type IService interface {
	CreateMany(items models.MapRouteNodes) error
	DeleteAll() error
}

type IRepository interface {
	db() *bun.DB
	DeleteAll() error
	CreateMany(items models.MapRouteNodes) error
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
	ctx         context.Context
	helper      *helper.Helper
	queryFilter *sqlHelper.QueryFilter
}

func CreateHandler(helper *helper.Helper) *Handler {
	repo := NewRepository(helper)
	service := NewService(repo, helper)
	return NewHandler(service, helper)
}

func CreateService(helper *helper.Helper) *Service {
	repo := NewRepository(helper)
	return NewService(repo, helper)
}

// NewHandler constructor
func NewHandler(s IService, helper *helper.Helper) *Handler {
	return &Handler{service: s, helper: helper}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}
