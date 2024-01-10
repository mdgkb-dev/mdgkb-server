package maproutes

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetMapRoute(c *gin.Context)
}

type IService interface {
	setQueryFilter(*gin.Context) error
	GetMapRoute(startID string, endID string) (*models.MapRoute, error)
	CreateMany(items models.MapRoutes) error
	DeleteAll() error
}

type IRepository interface {
	setQueryFilter(*gin.Context) error
	db() *bun.DB
	GetMapRoute(startID string, endID string) (*models.MapRoute, error)
	DeleteAll() error
	CreateMany(items models.MapRoutes) error
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
