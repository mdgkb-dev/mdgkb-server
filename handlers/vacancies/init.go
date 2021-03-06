package vacancies

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
	GetBySlug(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)

	CreateResponse(c *gin.Context)
}

type IService interface {
	GetAll() (models.Vacancies, error)
	GetAllWithResponses() (models.Vacancies, error)
	Get(*string) (*models.Vacancy, error)
	GetBySlug(*string) (*models.Vacancy, error)
	Create(*models.Vacancy) error
	Update(*models.Vacancy) error
	Delete(*string) error

	CreateResponse(*models.VacancyResponse) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.Vacancy) error
	getAll() (models.Vacancies, error)
	getBySlug(*string) (*models.Vacancy, error)
	getAllWithResponses() (models.Vacancies, error)
	get(*string) (*models.Vacancy, error)
	update(*models.Vacancy) error
	delete(*string) error

	createResponse(*models.VacancyResponse) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.VacancyResponse, map[string][]*multipart.FileHeader) error
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
