package heads

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"mdgkb/mdgkb-server/handlers/baseHandler"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

type IService interface {
	baseHandler.IService
	Create(*models.Head) error
	GetAll() (models.Heads, error)
	Get(string) (*models.Head, error)
	Delete(string) error
	Update(*models.Head) error
}

type IRepository interface {
	baseHandler.IRepository
	create(*models.Head) error
	getAll() (models.Heads, error)
	get(string) (*models.Head, error)
	delete(string) error
	update(*models.Head) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.Head, map[string][]*multipart.FileHeader) error
}

type Handler struct {
	service      IService
	filesService IFilesService
	helper       *helper.Helper
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

type FilesService struct {
	uploader uploadHelper.Uploader
	helper   *helper.Helper
}

func CreateHandler(helper *helper.Helper) *Handler {
	repo := NewRepository(helper)
	service := NewService(repo, helper)
	filesService := NewFilesService(helper)
	return NewHandler(service, filesService, helper)
}

func CreateService(helper *helper.Helper) *Service {
	repo := NewRepository(helper)
	return NewService(repo, helper)
}

// NewHandler constructor
func NewHandler(s IService, filesService IFilesService, helper *helper.Helper) *Handler {
	return &Handler{service: s, filesService: filesService, helper: helper}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}
