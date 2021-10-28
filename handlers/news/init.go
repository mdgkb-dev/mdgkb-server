package news

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
	Create(*models.News) error
	Update(*models.News) error
	CreateLike(*models.NewsLike) error
	AddTag(*models.NewsToTag) error
	RemoveTag(*models.NewsToTag) error
	CreateComment(*models.NewsComment) error
	UpdateComment(*models.NewsComment) error
	RemoveComment(string) error
	GetAll(*newsParams) ([]models.News, error)
	Delete(string) error
	DeleteLike(string) error
	GetBySlug(string) (*models.News, error)
	GetByMonth(*monthParams) ([]models.News, error)
	CreateViewOfNews(*models.NewsView) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.News) error
	update(*models.News) error
	createLike(*models.NewsLike) error
	addTag(*models.NewsToTag) error
	removeTag(*models.NewsToTag) error
	createComment(*models.NewsComment) error
	updateComment(*models.NewsComment) error
	removeComment(string) error
	getAll(*newsParams) ([]models.News, error)
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
