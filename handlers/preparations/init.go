package preparations

import (
	"context"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/uploadHelper"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	GetTags(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	UpdateMany(c *gin.Context)
}

type IService interface {
	Create(*models.Preparation) error
	GetAll() (models.Preparations, error)
	Get(string) (*models.Preparation, error)
	Delete(string) error
	Update(*models.Preparation) error
	UpsertMany(WithDeleted) error
	GetTags() (models.PreparationsTags, error)
}

type IRepository interface {
	db() *bun.DB
	create(*models.Preparation) error
	getAll() (models.Preparations, error)
	get(string) (*models.Preparation, error)
	delete(string) error
	update(*models.Preparation) error
	upsertMany(models.Preparations) error
	deleteMany([]uuid.UUID) error
	getTags() (models.PreparationsTags, error)
}

type IFilesService interface {
	Upload(*gin.Context, *models.Preparation, map[string][]*multipart.FileHeader) error
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
	ctx    context.Context
	helper *helper.Helper
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
