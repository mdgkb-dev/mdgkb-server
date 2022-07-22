package visitsApplications

import (
	"context"
	"mdgkb/mdgkb-server/handlers/baseHandler"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type IService interface {
	baseHandler.IService
	GetAll() (models.VisitsApplicationsWithCount, error)
	Get(*string) (*models.VisitsApplication, error)
	Create(*models.VisitsApplication) error
	Update(*models.VisitsApplication) error
	Delete(*string) error
}

type IRepository interface {
	baseHandler.IRepository
	getAll() (models.VisitsApplicationsWithCount, error)
	get(*string) (*models.VisitsApplication, error)
	create(*models.VisitsApplication) error
	update(*models.VisitsApplication) error
	delete(*string) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.VisitsApplication, map[string][]*multipart.FileHeader) error
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
	helper *helper.Helper
}

func CreateHandler(helper *helper.Helper) *Handler {
	repo := NewRepository(helper)
	service := NewService(repo, helper)
	filesService := NewFilesService(helper)
	return NewHandler(service, filesService, helper)
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
