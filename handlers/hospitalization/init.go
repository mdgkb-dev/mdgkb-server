package hospitalization

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/pro-assistance/pro-assister/helper"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context) error
	PDF(c *gin.Context)
}
type IService interface {
	GetAll() (models.Hospitalizations, error)
	Get(string) (*models.Hospitalization, error)
}
type IRepository interface {
	db() *bun.DB
	getAll() (models.Hospitalizations, error)
	get(string) (*models.Hospitalization, error)
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
	ctx    context.Context
	helper *helper.Helper
}

func CreateHandler(helper *helper.Helper) *Handler {
	repo := NewRepository(helper)
	service := NewService(repo, helper)
	return NewHandler(service, helper)
}

func NewHandler(service IService, helper *helper.Helper) *Handler {
	return &Handler{service: service, helper: helper}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}
