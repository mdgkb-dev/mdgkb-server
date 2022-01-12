package documentTypes

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)

	GetDocumentsTypesForTablesNames(c *gin.Context)
}

type IService interface {
	GetAll(params models.DocumentsParams) ([]*models.DocumentType, error)
	Get(*string) (*models.DocumentType, error)
	Create(*models.DocumentType) error
	Update(*models.DocumentType) error
	Delete(*string) error
	UpsertMany(models.DocumentTypes) error
	DeleteMany([]uuid.UUID) error

	GetDocumentsTypesForTablesNames() map[string]string
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.DocumentType) error
	getAll(params models.DocumentsParams) (models.DocumentTypes, error)
	get(*string) (*models.DocumentType, error)
	update(*models.DocumentType) error
	delete(*string) error
	upsertMany(models.DocumentTypes) error
	deleteMany([]uuid.UUID) error
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

func CreateService(db *bun.DB) *Service {
	repo := NewRepository(db)
	return NewService(repo)
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
