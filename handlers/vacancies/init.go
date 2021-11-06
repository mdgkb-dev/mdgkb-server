package vacancies

import (
	"context"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
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

	CreateResponse(c *gin.Context)
}

type IService interface {
	GetAll() (models.Vacancies, error)
	GetAllWithResponses() (models.Vacancies, error)
	Get(*string) (*models.Vacancy, error)
	Create(*models.Vacancy) error
	Update(*models.Vacancy) error
	Delete(*string) error

	CreateResponse(*models.VacancyResponse) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.Vacancy) error
	getAll() (models.Vacancies, error)
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
}

type Service struct {
	repository IRepository
}

type Repository struct {
	db  *bun.DB
	ctx context.Context
}

type FilesService struct {
	uploader uploadHelper.Uploader
}

func CreateHandler(db *bun.DB, uploader uploadHelper.Uploader) *Handler {
	repo := NewRepository(db)
	service := NewService(repo)
	filesService := NewFilesService(uploader)
	return NewHandler(service, filesService)
}

// NewHandler constructor
func NewHandler(service IService, filesService IFilesService) *Handler {
	return &Handler{service: service, filesService: filesService}
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db, ctx: context.Background()}
}

func NewFilesService(uploader uploadHelper.Uploader) *FilesService {
	return &FilesService{uploader: uploader}
}
