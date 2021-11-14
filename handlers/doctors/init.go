package doctors

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
	GetByDivisionID(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	CreateComment(c *gin.Context)
	UpdateComment(c *gin.Context)
	RemoveComment(c *gin.Context)
}

type IService interface {
	Create(*models.Doctor) error
	GetAll(*doctorsParams) (models.Doctors, error)
	Get(string) (*models.Doctor, error)
	Delete(string) error
	Update(*models.Doctor) error

	GetByDivisionID(string) (models.Doctors, error)
	CreateComment(*models.DoctorComment) error
	UpdateComment(*models.DoctorComment) error
	RemoveComment(string) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.Doctor) error
	getAll(*doctorsParams) (models.Doctors, error)
	get(string) (*models.Doctor, error)
	getByDivisionID(string) (models.Doctors, error)
	delete(string) error
	update(*models.Doctor) error
	createComment(*models.DoctorComment) error
	updateComment(*models.DoctorComment) error
	removeComment(string) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.Doctor, map[string][]*multipart.FileHeader) error
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
