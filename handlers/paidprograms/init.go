package paidprograms

import (
	"context"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
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
	db() *bun.DB
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

func CreateService(helper *helper.Helper) *Service {
	repo := NewRepository(helper)
	return NewService(repo, helper)
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
