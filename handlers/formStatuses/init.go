package formStatuses

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	GetAllByGroupId(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	UpdateMany(c *gin.Context)
	Delete(c *gin.Context)
}

type IService interface {
	setQueryFilter(*gin.Context) error
	GetAll() (models.FormStatuses, error)
	GetAllByGroupId(*string) (models.FormStatuses, error)
	Get(*string) (*models.FormStatus, error)
	Upsert(*models.FormStatus) error
	UpsertMany(models.FormStatuses) error
	Delete(*string) error
}

type IRepository interface {
	setQueryFilter(*gin.Context) error
	getDB() *bun.DB
	getAll() (models.FormStatuses, error)
	GetAllByGroupId(*string) (models.FormStatuses, error)
	get(*string) (*models.FormStatus, error)
	upsert(*models.FormStatus) error
	upsertMany(models.FormStatuses) error
	delete(*string) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.FormStatus, map[string][]*multipart.FileHeader) error
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
	db          *bun.DB
	ctx         context.Context
	helper      *helper.Helper
	queryFilter *sqlHelper.QueryFilter
}

type FilesService struct {
	helper *helper.Helper
}

func CreateHandler(db *bun.DB, helper *helper.Helper) *Handler {
	repo := NewRepository(db, helper)
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

func NewRepository(db *bun.DB, helper *helper.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}
