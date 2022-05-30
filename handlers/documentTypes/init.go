package documentTypes

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
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
	Upsert(*models.DocumentType) error
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
	upsert(*models.DocumentType) error
	upsertMany(models.DocumentTypes) error
	deleteMany([]uuid.UUID) error
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
	db     *bun.DB
	ctx    context.Context
	helper *helper.Helper
}

func CreateHandler(db *bun.DB, helper *helper.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	return NewHandler(service, helper)
}

func CreateService(db *bun.DB, helper *helper.Helper) *Service {
	repo := NewRepository(db, helper)
	return NewService(repo, helper)
}

// NewHandler constructor
func NewHandler(s IService, helper *helper.Helper) *Handler {
	return &Handler{service: s, helper: helper}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helper.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}
