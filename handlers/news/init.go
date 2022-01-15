package news

import (
	"context"
	"mdgkb/mdgkb-server/helpers"
	httpHelper2 "mdgkb/mdgkb-server/helpers/httpHelperV2"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	GetAllAdmin(c *gin.Context)
	GetAllRelationsNews(c *gin.Context)
	GetBySLug(c *gin.Context)
	GetByMonth(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	CreateLike(c *gin.Context)
	AddTag(c *gin.Context)
	RemoveTag(c *gin.Context)
	Delete(c *gin.Context)
	DeleteLike(c *gin.Context)
	CreateComment(c *gin.Context)
	UpdateComment(c *gin.Context)
	RemoveComment(c *gin.Context)
}

type IService interface {
	setQueryFilter(*gin.Context) error

	Create(*models.News) error
	Update(*models.News) error
	CreateLike(*models.NewsLike) error
	AddTag(*models.NewsToTag) error
	RemoveTag(*models.NewsToTag) error
	CreateComment(*models.NewsComment) error
	UpdateComment(*models.NewsComment) error
	RemoveComment(string) error
	GetAll(*newsParams) ([]*models.News, error)
	GetAllAdmin() (models.NewsWithCount, error)
	GetAllRelationsNews(*newsParams) ([]models.News, error)
	Delete(string) error
	DeleteLike(string) error
	GetBySlug(string) (*models.News, error)
	GetByMonth(*monthParams) ([]models.News, error)
	CreateViewOfNews(*models.NewsView) error
}

type IRepository interface {
	setQueryFilter(*gin.Context) error

	getDB() *bun.DB
	create(*models.News) error
	update(*models.News) error
	createLike(*models.NewsLike) error
	addTag(*models.NewsToTag) error
	removeTag(*models.NewsToTag) error
	createComment(*models.NewsComment) error
	updateComment(*models.NewsComment) error
	removeComment(string) error
	getAll(*newsParams) ([]*models.News, error)
	getAllAdmin() (models.NewsWithCount, error)
	getAllRelationsNews(*newsParams) ([]models.News, error)
	delete(string) error
	deleteLike(string) error
	getBySlug(string) (*models.News, error)
	getByMonth(*monthParams) ([]models.News, error)
	createViewOfNews(*models.NewsView) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.News, map[string][]*multipart.FileHeader) error
}

type Handler struct {
	service      IService
	filesService IFilesService
	helper       *helpers.Helper
}

type Service struct {
	repository IRepository
	helper     *helpers.Helper
}

type Repository struct {
	db          *bun.DB
	ctx         context.Context
	helper      *helpers.Helper
	queryFilter *httpHelper2.QueryFilter
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

// NewHandler constructor
func NewHandler(s IService, filesService IFilesService, helper *helpers.Helper) *Handler {
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
