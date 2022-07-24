package donorrules

import (
	"context"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"

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
	UpsertMany(RulesWithDeleted) error
	AddToUser(*models.DonorRuleUser) error
	DeleteFromUser(*models.DonorRuleUser) error
}

type IRepository interface {
	db() *bun.DB
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
	helper       *helper.Helper
}

type Service struct {
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
}

type FilesService struct {
	//uploader uploadHelper.Uploader
	helper *helper.Helper
}

func CreateHandler(helper *helper.Helper) *Handler {
	repo := NewRepository(helper)
	service := NewService(repo, helper)
	filesService := NewFilesService(helper)
	return NewHandler(service, filesService, helper)
}

func CreateService(helper *helper.Helper) *Service {
	repo := NewRepository(helper)
	return NewService(repo, helper)
}

// NewHandler constructor
func NewHandler(s IService, filesService IFilesService, helper *helper.Helper) *Handler {
	return &Handler{service: s, filesService: filesService, helper: helper}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}
