package newsSlides

import (
	"context"
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
	UpdateAll(c *gin.Context)
}

type IService interface {
	GetAll() (models.NewsSlides, error)
	Get(string) (*models.NewsSlide, error)
	Create(*models.NewsSlide) error
	Update(*models.NewsSlide) error
	Delete(string) error
	UpdateAll(models.NewsSlides) error
}

type IRepository interface {
	db() *bun.DB
	getAll() (models.NewsSlides, error)
	get(string) (*models.NewsSlide, error)
	create(*models.NewsSlide) error
	update(*models.NewsSlide) error
	delete(string) error
	updateAll(models.NewsSlides) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.NewsSlide, map[string][]*multipart.FileHeader) error
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
