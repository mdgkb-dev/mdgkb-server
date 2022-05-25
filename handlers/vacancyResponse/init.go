package vacancyResponse

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
	EmailExists(c *gin.Context)

	PDF(c *gin.Context)
}

type IService interface {
	GetAll() (models.VacancyResponses, error)
	Get(string) (*models.VacancyResponse, error)
	Create(*models.VacancyResponse) error
	Update(*models.VacancyResponse) error
	Delete(string) error
	EmailExists(string, string) (bool, error)
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.VacancyResponse) error
	getAll() (models.VacancyResponses, error)
	get(string) (*models.VacancyResponse, error)
	update(*models.VacancyResponse) error
	delete(string) error
	emailExists(string, string) (bool, error)
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
