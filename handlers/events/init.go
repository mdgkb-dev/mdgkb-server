package events

import (
	"context"
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/uptrace/bun"
)

type IHandler interface {
	CreateEventApplication(c *gin.Context)
	EventApplicationsPDF(c *gin.Context)
}


type IService interface {
	Create(info *models.Event) error
	Get(string) (*models.Event, error)
	Update(info *models.Event) error
	Upsert(info *models.Event) error
	UpsertMany(infos models.Events) error
	CreateEventApplication(*models.EventApplication) error
}

type IRepository interface {
	getDB() *bun.DB
	create(info *models.Event) error
	get(string) (*models.Event, error)
	update(info *models.Event) error
	upsert(info *models.Event) error
	upsertMany(infos models.Events) error
	createEventApplication(*models.EventApplication) error
	//deleteMany([]string) error
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
	helper *helpers.Helper
}

type IFilesService interface {
	Upload(*gin.Context, *models.Event, map[string][]*multipart.FileHeader) error
}

func CreateHandler(db *bun.DB, helper *helpers.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	filesService := NewFilesService(helper)
	return NewHandler(service, filesService, helper)
}


func CreateService(db *bun.DB, helper *helpers.Helper) *Service {
	repo := NewRepository(db, helper )
	return NewService(repo, helper )
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
