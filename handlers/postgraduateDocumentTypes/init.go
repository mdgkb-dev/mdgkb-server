package postgraduateDocumentTypes

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
	GetAll() (models.PostgraduateDocumentTypes, error)
	Get(string) (*models.PostgraduateDocumentType, error)
	Create(*models.PostgraduateDocumentType) error
	Update(item *models.PostgraduateDocumentType) error
	Delete(string) error

	UpsertMany(item PostgraduateDocumentTypesWithDelete) error
	DeleteMany(uuid []uuid.UUID) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.PostgraduateDocumentType) error
	getAll() (models.PostgraduateDocumentTypes, error)
	get(string) (*models.PostgraduateDocumentType, error)
	update(item *models.PostgraduateDocumentType) error
	delete(string) error

	upsertMany(item models.PostgraduateDocumentTypes) error
	deleteMany(uuid []uuid.UUID) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.PostgraduateDocumentTypes, map[string][]*multipart.FileHeader) error
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
	db     *bun.DB
	ctx    context.Context
	helper *helper.Helper
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

func CreateService(db *bun.DB, helper *helper.Helper) *Service {
	repo := NewRepository(db, helper)
	return NewService(repo, helper)
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
