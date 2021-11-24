package newsSlides

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
}

type IService interface {
	GetAll() (models.NewsSlides, error)
	Get(string) (*models.NewsSlide, error)
	Create(*models.NewsSlide) error
	Update(*models.NewsSlide) error
	Delete(string) error
}

type IRepository interface {
	getDB() *bun.DB
	getAll() (models.NewsSlides, error)
	get(string) (*models.NewsSlide, error)
	create(*models.NewsSlide) error
	update(*models.NewsSlide) error
	delete(string) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.NewsSlide, map[string][]*multipart.FileHeader) error
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
	return NewHandler(service, filesService, helper)
}

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
