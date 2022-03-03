package donorRules

import (
	"context"
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	UpdateMany(c *gin.Context)
	AddToUser(c *gin.Context)
	DeleteFromUser(c *gin.Context)
}

type IService interface {
	GetAll(*uuid.UUID) (models.DonorRules, error)
	UpsertMany(DonorRulesWithDeleted) error
	AddToUser(*models.DonorRuleUser) error
	DeleteFromUser(*models.DonorRuleUser) error
}

type IRepository interface {
	getDB() *bun.DB
	getAll(*uuid.UUID) (models.DonorRules, error)
	upsertMany(models.DonorRules) error
	deleteMany([]uuid.UUID) error
	addToUser(*models.DonorRuleUser) error
	deleteFromUser(*models.DonorRuleUser) error
}

type IFilesService interface {
	Upload(*gin.Context, models.DonorRules, map[string][]*multipart.FileHeader) error
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
