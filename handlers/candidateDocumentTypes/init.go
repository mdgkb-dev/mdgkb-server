package candidateDocumentTypes

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
	GetAll() (models.CandidateDocumentTypes, error)
	Get(string) (*models.CandidateDocumentType, error)
	Create(*models.CandidateDocumentType) error
	Update(item *models.CandidateDocumentType) error
	Delete(string) error

	UpsertMany(item CandidateDocumentTypesWithDelete) error
	DeleteMany(uuid []uuid.UUID) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.CandidateDocumentType) error
	getAll() (models.CandidateDocumentTypes, error)
	get(string) (*models.CandidateDocumentType, error)
	update(item *models.CandidateDocumentType) error
	delete(string) error

	upsertMany(item models.CandidateDocumentTypes) error
	deleteMany(uuid []uuid.UUID) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.CandidateDocumentTypes, map[string][]*multipart.FileHeader) error
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
