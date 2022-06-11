package vacancies

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"
	"mdgkb/mdgkb-server/handlers/baseHandler"
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
	UpdateMany(c *gin.Context)

	CreateResponse(c *gin.Context)
}

type IService interface {
	baseHandler.IService
	GetAll() (models.VacanciesWithCount, error)
	Get(*string) (*models.Vacancy, error)
	GetBySlug(*string) (*models.Vacancy, error)
	Create(*models.Vacancy) error
	Update(*models.Vacancy) error
	Delete(*string) error
	UpdateMany(models.Vacancies) error

	CreateResponse(*models.VacancyResponse) error
}

type IRepository interface {
	baseHandler.IRepository
	create(*models.Vacancy) error
	getAll() (models.VacanciesWithCount, error)
	getBySlug(*string) (*models.Vacancy, error)
	get(*string) (*models.Vacancy, error)
	update(*models.Vacancy) error
	delete(*string) error
	upsertMany(models.Vacancies) error

	createResponse(*models.VacancyResponse) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.VacancyResponse, map[string][]*multipart.FileHeader) error
	UploadVacancy(*gin.Context, *models.Vacancy, map[string][]*multipart.FileHeader) error
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
	db          *bun.DB
	ctx         context.Context
	helper      *helper.Helper
	queryFilter *sqlHelper.QueryFilter
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
