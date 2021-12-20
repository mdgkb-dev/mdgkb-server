package projects

import (
	"context"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/models"

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
	getDB() *bun.DB
	create(*models.Project) error
	getAll() (models.Projects, error)
	get(*string) (*models.Project, error)
	update(*models.Project) error
	delete(*string) error

	getBySlug(*string) (*models.Project, error)
}

type Handler struct {
	service      IService
	helper       *helpers.Helper
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

// NewHandler constructor
func NewHandler(s IService, helper *helpers.Helper) *Handler {
	return &Handler{service: s,  helper: helper}
}

func NewService(repository IRepository, helper *helpers.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helpers.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}

