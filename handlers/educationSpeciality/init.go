package educationSpeciality

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/mdgkb-server/models"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type IService interface {
	GetAll() (models.EducationSpecialities, error)
	Get(*string) (*models.EducationSpeciality, error)
	Create(*models.EducationSpeciality) error
	Update(*models.EducationSpeciality) error
	Delete(*string) error
}

type IRepository interface {
	create(*models.EducationSpeciality) error
	getAll() (models.EducationSpecialities, error)
	get(*string) (*models.EducationSpeciality, error)
	update(*models.EducationSpeciality) error
	delete(*string) error
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
