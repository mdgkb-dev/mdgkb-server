package auth

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"
	"github.com/uptrace/bun"
)

type IHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	LoginAs(c *gin.Context)
	Logout(c *gin.Context)
	RefreshToken(c *gin.Context)
	RefreshPassword(c *gin.Context)
	RestorePassword(c *gin.Context)
	CheckUUID(c *gin.Context)
	SavePathPermissions(c *gin.Context)
	GetAllPathPermissions(c *gin.Context)
	GetAllPathPermissionsAdmin(c *gin.Context)
	GetPathPermissionsByRoleID(c *gin.Context)
	CheckPathPermissions(c *gin.Context)
}

type IService interface {
	setQueryFilter(*gin.Context) error
	Register(user *models.User) (*models.TokensWithUser, error)
	Login(user *models.User, skipPassword bool) (*models.TokensWithUser, error)
	FindUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	DropUUID(*models.User) error
	UpdatePassword(*models.User) error
	UpsertManyPathPermissions(models.PathPermissions) error
	GetAllPathPermissions() (models.PathPermissions, error)
	GetAllPathPermissionsAdmin() (models.PathPermissionsWithCount, error)
	GetPathPermissionsByRoleID(id string) (models.PathPermissions, error)
	CheckPathPermissions(path string, roleID string) error
}

type IRepository interface {
	setQueryFilter(*gin.Context) error
	db() *bun.DB
	getAllPathPermissions() (models.PathPermissions, error)
	getAllPathPermissionsAdmin() (models.PathPermissionsWithCount, error)
	getPathPermissionsByRoleID(id string) (models.PathPermissions, error)
	upsertManyPathPermissions(items models.PathPermissions) (err error)
	deleteManyPathPermissions(idPool []uuid.UUID) (err error)
	upsertManyPathPermissionsRoles(items models.PathPermissionsRoles) (err error)
	deleteManyPathPermissionsRoles(idPool []uuid.UUID) (err error)
	checkPathPermissions(path string, roleID string) error
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
	ctx         context.Context
	helper      *helper.Helper
	queryFilter *sqlHelper.QueryFilter
}

type FilesService struct {
	helper *helper.Helper
}

func CreateHandler(helper *helper.Helper) *Handler {
	repo := NewRepository(helper)
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

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}
