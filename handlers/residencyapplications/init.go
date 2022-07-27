package residencyapplications

import (
	"context"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	EmailExists(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	UpdateWithForm(c *gin.Context)
	Delete(c *gin.Context)
	UpsertMany(c *gin.Context)

	FillApplicationTemplate(c *gin.Context)
}

type IService interface {
	setQueryFilter(*gin.Context) error
	GetAll() (models.ResidencyApplicationsWithCount, error)
	Get(*string) (*models.ResidencyApplication, error)
	EmailExists(string, string) (bool, error)
	Create(*models.ResidencyApplication) error
	Update(*models.ResidencyApplication) error
	UpdateWithForm(*models.FormValue) error
	UpsertMany(models.ResidencyApplications) error
	Delete(*string) error
}

type IRepository interface {
	setQueryFilter(*gin.Context) error
	db() *bun.DB
	getAll() (models.ResidencyApplicationsWithCount, error)
	get(*string) (*models.ResidencyApplication, error)
	emailExists(string, string) (bool, error)
	create(*models.ResidencyApplication) error
	update(*models.ResidencyApplication) error
	upsertMany(models.ResidencyApplications) (err error)
	delete(*string) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.ResidencyApplication, map[string][]*multipart.FileHeader) error
	UploadFormFiles(*gin.Context, *models.FormValue, map[string][]*multipart.FileHeader) error
	FillApplicationTemplate(*models.ResidencyApplication) ([]byte, error)
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
