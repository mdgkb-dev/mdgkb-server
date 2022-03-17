package documentTypes

import (
	"context"
	"mdgkb/mdgkb-server/helpers"
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
	helper  *helpers.Helper
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

func CreateService(db *bun.DB, helper *helpers.Helper) *Service {
	repo := NewRepository(db, helper)
	return NewService(repo, helper)
}

// NewHandler constructor
func NewHandler(s IService, helper *helpers.Helper) *Handler {
	return &Handler{service: s, helper: helper}
}

func NewService(repository IRepository, helper *helpers.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helpers.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}
