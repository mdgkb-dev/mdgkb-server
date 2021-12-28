package paidPrograms

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"
)

type IHandler interface {
	Get(c *gin.Context)
	Update(c *gin.Context)
}


type IService interface {
	CreateMany(models.PaidPrograms) error
	UpsertMany(models.PaidPrograms) error
	DeleteMany([]uuid.UUID) error

	Get(string) (*models.PaidProgram, error)
	Update(*models.PaidProgram) error
}

type IRepository interface {
	getDB() *bun.DB
	createMany(models.PaidPrograms) error
	upsertMany(models.PaidPrograms) error
	deleteMany([]uuid.UUID) error

	get(string) (*models.PaidProgram, error)
	update(*models.PaidProgram) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.PaidProgram, map[string][]*multipart.FileHeader) error
}



type Handler struct {
	service      IService
	filesService IFilesService
	helper       *helpers.Helper
}

type Service struct {
	repository IRepository
	helper     *helpers.Helper
}

type Repository struct {
	db     *bun.DB
	ctx    context.Context
	helper *helpers.Helper
}

type FilesService struct {
	uploader uploadHelper.Uploader
	helper   *helpers.Helper
}

func CreateService(db *bun.DB, helper *helpers.Helper) *Service {
	repo := NewRepository(db, helper)
	return NewService(repo, helper)
}

func CreateHandler(db *bun.DB, helper *helpers.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	filesService := NewFilesService(helper)
	return NewHandler(service, filesService, helper)
}

// NewHandler constructor
func NewHandler(s IService, filesService IFilesService, helper *helpers.Helper) *Handler {
	return &Handler{service: s, filesService: filesService, helper: helper}
}

func NewService(repository IRepository, helper *helpers.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helpers.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}

func NewFilesService(helper *helpers.Helper) *FilesService {
	return &FilesService{helper: helper}
}
