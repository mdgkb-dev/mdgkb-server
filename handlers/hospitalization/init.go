package hospitalization

import (
	"context"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/models"

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
	getDB() *bun.DB
	getAll() (models.Hospitalizations, error)
	get(string) (*models.Hospitalization, error)
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

func CreateHandler(db *bun.DB, helper *helpers.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	return NewHandler(service, helper)
}

func NewHandler(service IService, helper *helpers.Helper) *Handler {
	return &Handler{service: service, helper: helper}
}

func NewService(repository IRepository, helper *helpers.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helpers.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}
