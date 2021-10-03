package divisions

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"
)

type IHandler interface {
	GetAll(c *gin.Context) error
	Get(c *gin.Context) error
	Create(c *gin.Context) error
	Delete(c *gin.Context) error
	Update(c *gin.Context) error
	CreateComment(c *gin.Context) error
	UpdateComment(c *gin.Context) error
	RemoveComment(c *gin.Context) error
}


type IService interface {
	Create(*models.Division) error
	GetAll() (models.Divisions, error)
	Get(*string) (*models.Division, error)
	Delete(*string) error
	Update(*models.Division) error
	CreateComment(*models.DivisionComment) error
	UpdateComment(*models.DivisionComment) error
	RemoveComment(*string) error
}

type IRepository interface {
	getDB() *bun.DB
	create( *models.Division) error
	getAll() (models.Divisions, error)
	get(*string) (*models.Division, error)
	delete(*string) error
	update( *models.Division) error
	createComment( *models.DivisionComment) error
	updateComment( *models.DivisionComment) error
	removeComment(*string) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.Division, map[string][]*multipart.FileHeader) error
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

func CreateHandler(db *bun.DB, uploader *uploadHelper.Uploader) *Handler {
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

func NewFilesService(uploader *uploadHelper.Uploader) *FilesService {
	return &FilesService{uploader: *uploader}
}
