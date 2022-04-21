package educationalOrganizationAcademics

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
}

type IService interface {
	setQueryFilter(*gin.Context) error
	GetAll() (models.EducationalOrganizationAcademics, error)
	UpsertMany(models.EducationalOrganizationAcademics) error
	Upsert(*models.EducationalOrganizationAcademic) error
	DeleteMany([]string) error
	DeleteByDoctorID(id uuid.NullUUID) (error)
}

type IRepository interface {
	setQueryFilter(*gin.Context) error
	getAll() (models.EducationalOrganizationAcademics, error)
	upsertMany(models.EducationalOrganizationAcademics) error
	upsert(*models.EducationalOrganizationAcademic) error
	deleteMany([]string) error
	deleteByDoctorID(id uuid.NullUUID) (error)
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
	db  *bun.DB
	ctx context.Context
	helper      *helper.Helper
	queryFilter *sqlHelper.QueryFilter
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

func NewHandler(s IService, helper *helper.Helper) *Handler {
	return &Handler{service: s, helper: helper}
}	

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helper.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}
