package tags

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/helper"
)

type IRepository interface {
	create(*gin.Context, *models.Tag) error
	getAll(*gin.Context) ([]models.Tag, error)
	get(*gin.Context, string) (models.Tag, error)
	updateStatus(*gin.Context, *models.Tag) error
	delete(*gin.Context, string) error
}
type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	UpdateStatus(c *gin.Context)
}

func CreateHandler(helper *helper.Helper) *Handler {
	repo := NewRepository(helper)
	return NewHandler(repo, helper)
}

type Handler struct {
	repository IRepository
	helper     *helper.Helper
}

// NewHandler constructor
func NewHandler(repository IRepository, h *helper.Helper) *Handler {
	return &Handler{
		helper:     h,
		repository: repository,
	}
}

type Repository struct {
	helper *helper.Helper
}

func NewRepository(h *helper.Helper) *Repository {
	return &Repository{helper: h}
}
