package pages

import (
	"context"
	"mdgkb/mdgkb-server/helpers"
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

	GetBySlug(c *gin.Context)
}

type IService interface {
	GetAll() (models.Pages, error)
	Get(*string) (*models.Page, error)
	Create(*models.Page) error
	Update(*models.Page) error
	Delete(*string) error

	GetBySlug(*string) (*models.Page, error)
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.Page) error
	getAll() (models.Pages, error)
	get(*string) (*models.Page, error)
	update(*models.Page) error
	delete(*string) error

	getBySlug(*string) (*models.Page, error)
}

type IFilesService interface {
	Upload(*gin.Context, *models.Page, map[string][]*multipart.FileHeader) error
}

type Handler struct {
	service IService
	filesService IFilesService
	helper *helpers.Helper
}

type Service struct {
	repository IRepository
	helper *helpers.Helper
}

type Repository struct {
	db  *bun.DB
	ctx context.Context
	helper *helpers.Helper
}

type FilesService struct {
	helper *helpers.Helper
}

func CreateHandler(db *bun.DB, helper *helpers.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	filesService := NewFilesService(helper)
	return NewHandler(service,filesService, helper)
}

// NewHandler constructor
func NewHandler(s IService,filesService IFilesService, helper *helpers.Helper) *Handler {
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
