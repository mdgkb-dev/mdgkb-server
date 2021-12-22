package partnerTypes

import (
	"context"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/models"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type IService interface {
	GetAll() (models.PartnerTypes, error)
	Get(string) (*models.PartnerType, error)
	Create(*models.PartnerType) error
	Update(*models.PartnerType) error
	Delete(string) error
}

type IRepository interface {
	getDB() *bun.DB
	getAll() (models.PartnerTypes, error)
	get(string) (*models.PartnerType, error)
	create(*models.PartnerType) error
	update(item *models.PartnerType) error
	delete(string) error
}

type Handler struct {
	service      IService
	helper       *helpers.Helper
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

type FilesService struct {
	helper *helpers.Helper
}

func CreateHandler(db *bun.DB, helper *helpers.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	return NewHandler(service, helper)
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
