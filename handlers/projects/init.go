package projects

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/pro-assistance/pro-assister/helper"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)

	GetBySlug(c *gin.Context)
}

type IService interface {
	GetAll() (models.Projects, error)
	Get(*string) (*models.Project, error)
	Create(*models.Project) error
	Update(*models.Project) error
	Delete(*string) error

	GetBySlug(*string) (*models.Project, error)
}

type IRepository interface {
	db() *bun.DB
	create(*models.Project) error
	getAll() (models.Projects, error)
	get(*string) (*models.Project, error)
	update(*models.Project) error
	delete(*string) error

	getBySlug(*string) (*models.Project, error)
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
