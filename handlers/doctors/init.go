package doctors

import (
	"context"
	"mdgkb/mdgkb-server/helpers"
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
	service IService
	filesService IFilesService
	helper *helpers.Helper
}

type Service struct {
	repository IRepository
	helper *helpers.Helper
}

type Repository struct {
	db  *bun.DB
	ctx context.Context
	helper *helpers.Helper
}

type FilesService struct {
	helper *helpers.Helper
}

func CreateHandler(db *bun.DB, helper *helpers.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	filesService := NewFilesService(helper)
	return NewHandler(service,filesService, helper)
}

// NewHandler constructor
func NewHandler(s IService,filesService IFilesService, helper *helpers.Helper) *Handler {
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
