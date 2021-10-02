package pages

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context) error
	Get(c *gin.Context) error
	Create(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error

	GetBySlug(c *gin.Context) error
}

type IService interface {
	GetAll() (models.Pages, error)
	Get(*string) (*models.Page, error)
	Create(*models.Page) error
	Update(*models.Page) error
	Delete(*string) error

	GetBySlug(*string) (*models.Page, error)
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.Page) error
	getAll() (models.Pages, error)
	get(*string) (*models.Page, error)
	update(*models.Page) error
	delete(*string) error

	getBySlug(*string) (*models.Page, error)
}


type Handler struct {
	service      IService
}

type Service struct {
	repository IRepository
}

type Repository struct {
	db  *bun.DB
	ctx context.Context
}

func CreateHandler(db *bun.DB) *Handler {
	repo := NewRepository(db)
	service := NewService(repo)
	return NewHandler(service )
}

// NewHandler constructor
func NewHandler(service IService) *Handler {
	return &Handler{service: service}
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db, ctx: context.Background()}
}
