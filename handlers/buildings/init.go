package buildings

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/helper"
)

type IRepository interface {
	create(*gin.Context, *models.Building) error
	getAll(*gin.Context) ([]models.Building, error)
	getByFloorID(*gin.Context, string) (models.Building, error)
	getByID(*gin.Context, string) (models.Building, error)
	delete(*gin.Context, string) error
	update(*gin.Context, *models.Building) error
}
type IHandler interface {
	GetAll(c *gin.Context)
	GetByFloorID(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
}

type Handler struct {
	repository IRepository
	helper     *helper.Helper
	//uploader   helper.Uploader
}

func CreateHandler(helper *helper.Helper) *Handler {
	repo := NewRepository(helper)
	return NewHandler(repo, helper)
}

// NewHandler constructor
func NewRepository(h *helper.Helper) *Repository {
	return &Repository{
		helper: h,
		ctx:    context.Background(),
	}
}

// NewHandler constructor
func NewHandler(repository IRepository, h *helper.Helper) *Handler {
	return &Handler{
		repository: repository,
		helper:     h,
	}
}
