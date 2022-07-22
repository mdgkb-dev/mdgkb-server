package residencyDocumentTypes

import (
	"context"
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type IService interface {
	GetAll() (models.ResidencyDocumentTypes, error)
	Get(string) (*models.ResidencyDocumentType, error)
	Create(*models.ResidencyDocumentType) error
	Update(item *models.ResidencyDocumentType) error
	Delete(string) error

	UpsertMany(item ResidencyDocumentTypesWithDelete) error
	DeleteMany(uuid []uuid.UUID) error
}

type IRepository interface {
	db() *bun.DB
	create(*models.ResidencyDocumentType) error
	getAll() (models.ResidencyDocumentTypes, error)
	get(string) (*models.ResidencyDocumentType, error)
	update(item *models.ResidencyDocumentType) error
	delete(string) error

	upsertMany(item models.ResidencyDocumentTypes) error
	deleteMany(uuid []uuid.UUID) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.ResidencyDocumentTypes, map[string][]*multipart.FileHeader) error
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
