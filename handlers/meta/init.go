package meta

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/pro-assistance/pro-assister/helper"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	GetCount(c *gin.Context)
	GetSchema(c *gin.Context)
	GetSocial(c *gin.Context)
	GetOptions(c *gin.Context)
	GetApplicationsCounts(c *gin.Context)
	GetWeb(c *gin.Context)
}

type IService interface {
	GetCount(*string) (*int, error)
	GetOptions(*models.OptionModel) (models.Options, error)
	GetApplicationsCounts() (models.ApplicationsCounts, error)
	SendApplicationsCounts() error
}

type IRepository interface {
	getCount(*string) (*int, error)
	getOptions(*models.OptionModel) (models.Options, error)
	getApplicationsCounts() (models.ApplicationsCounts, error)
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
