package educationalOrganization

import (
	"context"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	Get(c *gin.Context) error
	Update(c *gin.Context) error
}

type IService interface {
	Get() (*models.EducationalOrganization, error)
	Update(*models.EducationalOrganization) error
}

type IRepository interface {
	getDB() *bun.DB
}

type Handler struct {
	service IService
	uploader   helpers.Uploader
}

type Service struct {
	repository IRepository
}

type Repository struct {
	db  *bun.DB
	ctx context.Context
}

func CreateHandler(db *bun.DB, uploader helpers.Uploader) *Handler {
	repo := NewRepository(db)
	service := NewService(repo)
	return NewHandler(service, uploader )
}

// NewHandler constructor
func NewHandler(s IService,  uploader helpers.Uploader) *Handler {
	return &Handler{service: s, uploader:   uploader}
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db, ctx: context.Background()}
}
