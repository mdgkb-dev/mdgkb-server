package users

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
	GetByEmail(c *gin.Context)
}

type IService interface {
	GetAll() (models.Users, error)
	Get(string) (*models.User, error)
	GetByEmail(string) (*models.User, error)
	EmailExists(string) (bool, error)
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.User) error
	getAll() (models.Users, error)
	get(string) (*models.User, error)
	getByEmail(string) (*models.User, error)
	emailExists(string) (bool, error)
}

type IFilesService interface {
	Upload(*gin.Context, *models.User, map[string][]*multipart.FileHeader) error
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


func CreateService(db *bun.DB) *Service {
	repo := NewRepository(db)
	return NewService(repo)
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
