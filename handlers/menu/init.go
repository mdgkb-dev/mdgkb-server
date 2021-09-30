package menu

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
}

type IService interface {
	GetAll() (models.Menus, error)
	Get(*string) (*models.Menu, error)
	Create(*models.Menu) error
	Update(*models.Menu) error
	Delete(*string) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.Menu) error
	getAll() (models.Menus, error)
	get(*string) (*models.Menu, error)
	update(*models.Menu) error
	delete(*string) error
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
