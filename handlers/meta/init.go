package meta

import (
	"context"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/models"
	"mdgkb/mdgkb-server/models/schema"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetCount(c *gin.Context)
	GetSchema(c *gin.Context)
	GetSocial(c *gin.Context)
	GetOptions(c *gin.Context)
}

type IService interface {
	GetCount(*string) (*int, error)
	GetSchema() schema.Schema
	GetOptions(*models.OptionModel) (models.Options, error)
}

type IRepository interface {
	getCount(*string) (*int, error)
	getOptions(*models.OptionModel) (models.Options, error)
}

type Handler struct {
	service IService
	helper  *helpers.Helper
}

type Service struct {
	repository IRepository
	helper     *helpers.Helper
}

type Repository struct {
	db     *bun.DB
	ctx    context.Context
	helper *helpers.Helper
}

func CreateHandler(db *bun.DB, helper *helpers.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	return NewHandler(service, helper)
}

func CreateService(db *bun.DB, helper *helpers.Helper) *Service {
	repo := NewRepository(db, helper)
	return NewService(repo, helper)
}

// NewHandler constructor
func NewHandler(s IService, helper *helpers.Helper) *Handler {
	return &Handler{service: s, helper: helper}
}

func NewService(repository IRepository, helper *helpers.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helpers.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}
