package valuetypes

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/pro-assistance/pro-assister/helper"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	GetAll(c *gin.Context)
}

type IService interface {
	GetAll() (models.ValueTypes, error)
}

type IRepository interface {
	getAll() (models.ValueTypes, error)
}

type Handler struct {
	helper  *helper.Helper
	service IService
}

type Service struct {
	helper     *helper.Helper
	repository IRepository
}

type Repository struct {
	helper *helper.Helper
	ctx    context.Context
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
