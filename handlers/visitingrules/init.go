package visitingrules

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	UpdateMany(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type IService interface {
	GetAll() (models.VisitingRules, error)
	Get(string) (*models.VisitingRule, error)
	Create(*models.VisitingRule) error
	UpsertMany(models.VisitingRules) error
	DeleteMany([]uuid.UUID) error
	UpsertAndDeleteMany(models.VisitingRulesWithDeleted) error
	Update(item *models.VisitingRule) error
	Delete(string) error
}

type IRepository interface {
	db() *bun.DB
	create(*models.VisitingRule) error
	getAll() (models.VisitingRules, error)
	get(string) (*models.VisitingRule, error)
	upsertMany(models.VisitingRules) error
	deleteMany([]uuid.UUID) error
	update(item *models.VisitingRule) error
	delete(string) error
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
