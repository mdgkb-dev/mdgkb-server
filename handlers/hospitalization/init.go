package hospitalization

import (
	"context"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context) error
}
type IService interface {
	GetAll() (models.Hospitalizations, error)
}
type IRepository interface {
	getDB() *bun.DB
	getAll() (models.Hospitalizations, error)
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

func CreateHandler(db *bun.DB, uploader *uploadHelper.Uploader) *Handler {
	repo := NewRepository(db)
	service := NewService(repo)
	return NewHandler(service)
}

func NewHandler(service IService) *Handler {
	return &Handler{service: service}
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db, ctx: context.Background()}
}
