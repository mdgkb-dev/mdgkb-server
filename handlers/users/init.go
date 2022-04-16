package users

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	GetByEmail(c *gin.Context)
	Update(c *gin.Context)
	AddToUser(c *gin.Context)
	RemoveFromUser(c *gin.Context)
}

type IService interface {
	GetAll() (models.Users, error)
	Get(string) (*models.User, error)
	GetByEmail(string) (*models.User, error)
	EmailExists(string) (bool, error)
	AddToUser(map[string]interface{}, string) error
	Update(*models.User) error
	Upsert(*models.User) error
	RemoveFromUser(map[string]interface{}, string) error
	UpsertEmail(*models.User) error
	DropUUID(*models.User) error
	UpdatePassword(*models.User) error
	SetAccessLink(item *models.User) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.User) error
	getAll() (models.Users, error)
	get(string) (*models.User, error)
	getByEmail(string) (*models.User, error)
	emailExists(string) (bool, error)
	update(*models.User) error
	upsert(*models.User) error
	upsertEmail(*models.User) error

	addToUser(map[string]interface{}, string) error
	removeFromUser(map[string]interface{}, string) error
	dropUUID(*models.User) error
	updatePassword(*models.User) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.User, map[string][]*multipart.FileHeader) error
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
	helper *helper.Helper
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
