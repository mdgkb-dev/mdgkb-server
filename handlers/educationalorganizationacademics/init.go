package educationalorganizationacademics

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"
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
	DeleteByDoctorID(id uuid.NullUUID) error
}

type IRepository interface {
	setQueryFilter(*gin.Context) error
	getAll() (models.EducationalOrganizationAcademics, error)
	upsertMany(models.EducationalOrganizationAcademics) error
	upsert(*models.EducationalOrganizationAcademic) error
	deleteMany([]string) error
	deleteByDoctorID(id uuid.NullUUID) error
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
	ctx         context.Context
	helper      *helper.Helper
	queryFilter *sqlHelper.QueryFilter
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

func NewHandler(s IService, helper *helper.Helper) *Handler {
	return &Handler{service: s, helper: helper}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}
