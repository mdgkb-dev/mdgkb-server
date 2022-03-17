package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/mdgkb-server/models"
)

type IHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	RefreshToken(c *gin.Context)
	RefreshPassword(c *gin.Context)
	RestorePassword(c *gin.Context)
}

type IService interface {
	Register(user *models.User) (*models.TokensWithUser, error)
	Login(user *models.User) (*models.TokensWithUser, error)
	FindUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	DropUUID(*models.User) error
	UpdatePassword(*models.User) error
}

type IRepository interface {
	getDB() *bun.DB
}

type IFilesService interface {
	//Upload(*gin.Context, *models.VacancyResponse, map[string][]*multipart.FileHeader) error
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
