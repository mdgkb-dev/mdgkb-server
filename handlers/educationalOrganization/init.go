package educationalOrganization

import (
	"context"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	Get(c *gin.Context) error
	Update(c *gin.Context) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.EducationalOrganization, map[string][]*multipart.FileHeader) error
}

type IService interface {
	Get() (*models.EducationalOrganization, error)
	Update(*models.EducationalOrganization) error
}

type FilesService struct {
	uploader uploadHelper.Uploader
}

type IRepository interface {
	getDB() *bun.DB
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
