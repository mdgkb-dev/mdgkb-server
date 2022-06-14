package formValues

import (
	"context"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/uptrace/bun"
)

type IHandler interface {
	Update(c *gin.Context)
	UpdateMany(c *gin.Context)
	Get(c *gin.Context)
	DocumentsToPDF(c *gin.Context)
}

type IService interface {
	Upsert(info *models.FormValue) error
	UpsertMany(models.FormValues) error
	Get(*string) (*models.FormValue, error)
}

type IRepository interface {
	getDB() *bun.DB
	upsert(info *models.FormValue) error
	upsertMany(models.FormValues) error
	get(*string) (*models.FormValue, error)
}

type IFilesService interface {
	Upload(*gin.Context, *models.FormValue, map[string][]*multipart.FileHeader) error
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

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helper.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}

func NewHandler(s IService, filesService IFilesService, helper *helper.Helper) *Handler {
	return &Handler{service: s, filesService: filesService, helper: helper}
}

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}
