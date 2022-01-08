package divisions

import (
	"context"
	"mdgkb/mdgkb-server/helpers"
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
	Delete(c *gin.Context)
	Update(c *gin.Context)
	CreateComment(c *gin.Context)
	UpdateComment(c *gin.Context)
	RemoveComment(c *gin.Context)
}

type IService interface {
	Create(*models.Division) error
	GetAll(bool) (models.Divisions, error)
	Get(string, bool) (*models.Division, error)
	Delete(string) error
	Update(*models.Division) error
	CreateComment(*models.DivisionComment) error
	UpdateComment(*models.DivisionComment) error
	RemoveComment(string) error
	GetBySearch(string) (models.Divisions, error)
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.Division) error
	getAll(bool) (models.Divisions, error)
	get(string, bool) (*models.Division, error)
	delete(string) error
	update(*models.Division) error
	createComment(*models.DivisionComment) error
	updateComment(*models.DivisionComment) error
	removeComment(string) error
	getBySearch(string) (models.Divisions, error)
}

type IFilesService interface {
	Upload(*gin.Context, *models.Division, map[string][]*multipart.FileHeader) error
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
	db     *bun.DB
	ctx    context.Context
	helper *helpers.Helper
}

type FilesService struct {
	uploader uploadHelper.Uploader
	helper   *helpers.Helper
}

func CreateHandler(db *bun.DB, helper *helpers.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	filesService := NewFilesService(helper)
	return NewHandler(service, filesService, helper)
}

func CreateService(db *bun.DB, helper *helpers.Helper) *Service {
	repo := NewRepository(db, helper)
	return NewService(repo, helper)
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
