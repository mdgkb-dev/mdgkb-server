package vacancyResponse

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
	GetAll() (models.VacancyResponses, error)
	Get(*string) (*models.VacancyResponse, error)
	Create(*models.VacancyResponse) error
	Update(*models.VacancyResponse) error
	Delete(*string) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.VacancyResponse) error
	getAll() (models.VacancyResponses, error)
	get(*string) (*models.VacancyResponse, error)
	update(*models.VacancyResponse) error
	delete(*string) error
}

type Handler struct {
	service IService
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
	return NewHandler(service)
}

// NewHandler constructor
func NewHandler(s IService) *Handler {
	return &Handler{service: s}
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db, ctx: context.Background()}
}