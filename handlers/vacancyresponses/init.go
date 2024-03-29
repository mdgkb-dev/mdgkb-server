package vacancyresponses

import (
	"context"
	"mime/multipart"

	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	EmailExists(c *gin.Context)

	PDF(c *gin.Context)
}

type IService interface {
	setQueryFilter(c *gin.Context) error
	GetAll() (models.VacancyResponsesWithCount, error)
	Get(string) (*models.VacancyResponse, error)
	Create(*models.VacancyResponse) error
	Update(*models.VacancyResponse) error
	Delete(string) error
	EmailExists(string, string) (bool, error)
	DeleteMany([]uuid.UUID) error
}

type IRepository interface {
	// basehandler.IRepository
	setQueryFilter(c *gin.Context) error
	create(*models.VacancyResponse) error
	getAll() (models.VacancyResponsesWithCount, error)
	get(string) (*models.VacancyResponse, error)
	update(*models.VacancyResponse) error
	delete(string) error
	emailExists(string, string) (bool, error)
	deleteMany([]uuid.UUID) error
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

func CreateService(helper *helper.Helper) *Service {
	repo := NewRepository(helper)
	return NewService(repo, helper)
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
