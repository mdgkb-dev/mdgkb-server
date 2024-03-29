package teachers

import (
	"context"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"

	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	CreateSlugs(c *gin.Context)
}

type IService interface {
	GetAll() (models.Teachers, error)
	Get(string) (*models.Teacher, error)
	Create(*models.Teacher) error
	Update(*models.Teacher) error
	Delete(string) error
	UpsertMany(models.Teachers) error
	DeleteMany([]string) error
	setQueryFilter(*gin.Context) error
	CreateSlugs() error
}

type IRepository interface {
	setQueryFilter(*gin.Context) error
	db() *bun.DB
	getAll() (models.Teachers, error)
	get(string) (*models.Teacher, error)
	create(*models.Teacher) error
	update(*models.Teacher) error
	delete(string) error

	upsertMany(models.Teachers) error
	deleteMany([]string) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.Teacher, map[string][]*multipart.FileHeader) error
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
