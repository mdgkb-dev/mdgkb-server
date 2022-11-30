package pagesections

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/pro-assistance/pro-assister/helper"

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
	GetAll(params models.DocumentsParams) ([]*models.PageSection, error)
	Get(*string) (*models.PageSection, error)
	Create(*models.PageSection) error
	Update(*models.PageSection) error
	Delete(*string) error
	Upsert(*models.PageSection) error
	UpsertMany(models.PageSections) error
	DeleteMany([]uuid.UUID) error

	GetDocumentsTypesForTablesNames() map[string]string
}

type IRepository interface {
	db() *bun.DB
	create(*models.PageSection) error
	getAll(params models.DocumentsParams) (models.PageSections, error)
	get(*string) (*models.PageSection, error)
	update(*models.PageSection) error
	delete(*string) error
	upsert(*models.PageSection) error
	upsertMany(models.PageSections) error
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
	ctx    context.Context
	helper *helper.Helper
}

func CreateHandler(helper *helper.Helper) *Handler {
	repo := NewRepository(helper)
	service := NewService(repo, helper)
	return NewHandler(service, helper)
}

func CreateService(helper *helper.Helper) *Service {
	repo := NewRepository(helper)
	return NewService(repo, helper)
}

// NewHandler constructor
func NewHandler(s IService, helper *helper.Helper) *Handler {
	return &Handler{service: s, helper: helper}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}
