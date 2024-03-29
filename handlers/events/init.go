package events

import (
	"context"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/helper"

	"github.com/uptrace/bun"
)

type IHandler interface {
	CreateEventApplication(c *gin.Context)
	EventApplicationsPDF(c *gin.Context)
	GetAllForMain(c *gin.Context)
}

type IService interface {
	Create(info *models.Event) error
	Get(string) (*models.Event, error)
	GetAllForMain() (models.Events, error)
	Update(info *models.Event) error
	Upsert(info *models.Event) error
	UpsertMany(infos models.Events) error
	CreateEventApplication(*models.EventApplication) error
}

type IRepository interface {
	db() *bun.DB
	create(info *models.Event) error
	get(string) (*models.Event, error)
	getAllForMain() (models.Events, error)
	update(info *models.Event) error
	upsert(info *models.Event) error
	upsertMany(infos models.Events) error
	createEventApplication(*models.EventApplication) error
	//deleteMany([]string) error
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
	helper *helper.Helper
}

type IFilesService interface {
	Upload(*gin.Context, *models.Event, map[string][]*multipart.FileHeader) error
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
