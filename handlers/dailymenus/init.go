package dailymenus

import (
	"context"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	UpdateAll(c *gin.Context)
	PDF(c *gin.Context)
	GetWeb(c *gin.Context)
	GetTodayMenu(c *gin.Context)
}

type IService interface {
	setQueryFilter(*gin.Context) error
	Create(*models.DailyMenu) error
	GetAll() (models.DailyMenus, error)
	Get(string) (*models.DailyMenu, error)
	Delete(string) error
	Update(*models.DailyMenu) error
	UpdateAll(models.DailyMenus) error
	GetTodayActive() (*models.DailyMenu, error)
}

type IRepository interface {
	setQueryFilter(*gin.Context) error
	db() *bun.DB
	create(*models.DailyMenu) error
	getAll() (models.DailyMenus, error)
	get(string) (*models.DailyMenu, error)
	delete(string) error
	update(*models.DailyMenu) error
	updateAll(models.DailyMenus) error
	getTodayActive() (*models.DailyMenu, error)
}

type IFilesService interface {
	Upload(*gin.Context, *models.DailyMenu, map[string][]*multipart.FileHeader) error
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
