package sideOrganizations

import (
	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/mdgkb-server/models"
)

type IHandler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	UpdateStatus(c *gin.Context)
	Delete(c *gin.Context)
}

type Handler struct {
	repository IRepository
	helper     *helper.Helper
}

func CreateHandler(helper *helper.Helper) *Handler {
	repo := NewRepository(helper)
	return NewHandler(repo, helper)
}

// NewHandler constructor
func NewHandler(repository IRepository, h *helper.Helper) *Handler {
	return &Handler{
		repository: repository,
		helper:     h,
	}
}

type IRepository interface {
	create(*gin.Context, *models.SideOrganization) error
	getAll(*gin.Context) ([]models.SideOrganization, error)
	get(*gin.Context, string) (models.SideOrganization, error)
	update(*gin.Context, *models.SideOrganization) error
	updateStatus(*gin.Context, *models.SideOrganization) error
	delete(*gin.Context, string) error
}

type Repository struct {
	helper *helper.Helper
}

func NewRepository(h *helper.Helper) *Repository {
	return &Repository{h}
}
