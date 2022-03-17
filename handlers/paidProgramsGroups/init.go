package paidProgramsGroups

import (
	"context"
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/uploadHelper"
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
	UpdateMany(c *gin.Context)
}

type IService interface {
	Create(*models.PaidProgramsGroup) error
	GetAll() (models.PaidProgramsGroups, error)
	Get(string) (*models.PaidProgramsGroup, error)
	Delete(string) error
	Update(*models.PaidProgramsGroup) error
	UpsertMany(models.PaidProgramsGroupsStruct) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.PaidProgramsGroup) error
	getAll() (models.PaidProgramsGroups, error)
	get(string) (*models.PaidProgramsGroup, error)
	delete(string) error
	update(*models.PaidProgramsGroup) error
	upsertMany(models.PaidProgramsGroups) error
	deleteMany([]uuid.UUID) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.PaidProgramsGroup, map[string][]*multipart.FileHeader) error
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
	db     *bun.DB
	ctx    context.Context
	helper *helper.Helper
}

type FilesService struct {
	uploader uploadHelper.Uploader
	helper   *helper.Helper
}

func CreateHandler(db *bun.DB, helper *helper.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	filesService := NewFilesService(helper)
	return NewHandler(service, filesService, helper)
}

func CreateService(db *bun.DB, helper *helper.Helper) *Service {
	repo := NewRepository(db, helper)
	return NewService(repo, helper)
}

// NewHandler constructor
func NewHandler(s IService, filesService IFilesService, helper *helper.Helper) *Handler {
	return &Handler{service: s, filesService: filesService, helper: helper}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helper.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}
