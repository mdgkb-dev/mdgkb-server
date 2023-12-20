package news

import (
	"context"
	"mdgkb/mdgkb-server/models"
	"mdgkb/mdgkb-server/models/exportmodels"
	"mime/multipart"

	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	GetMain(c *gin.Context)
	GetSubMain(c *gin.Context)
	GetAll(c *gin.Context)
	GetBySLug(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	CreateLike(c *gin.Context)
	AddTag(c *gin.Context)
	RemoveTag(c *gin.Context)
	Delete(c *gin.Context)
	DeleteLike(c *gin.Context)
	CreateComment(c *gin.Context)
	UpdateComment(c *gin.Context)
	RemoveComment(c *gin.Context)
	GetSuggestionNews(c *gin.Context)
	GetNewsComments(c *gin.Context)
}

type IService interface {
	SetQueryFilter(*gin.Context) error
	Create(*models.News) error
	Update(*models.News) error
	CreateLike(*models.NewsLike) error
	AddTag(*models.NewsToTag) error
	RemoveTag(*models.NewsToTag) error
	CreateComment(*models.NewsComment) error
	UpdateComment(*models.NewsComment) error
	RemoveComment(string) error
	GetAll() (models.NewsWithCount, error)
	GetMain() (models.NewsWithCount, error)
	GetSubMain() (models.NewsWithCount, error)
	Delete(string) error
	DeleteLike(string) error
	GetBySlug(*gin.Context, string) (*models.News, error)
	GetNewsComments(*gin.Context, string) (*models.NewsComments, error)
	CreateViewOfNews(*models.NewsView) error
	GetSuggestionNews(id string) ([]*models.News, error)
	GetAggregateViews(*exportmodels.NewsView) (models.ChartDataSets, error)
}

type IRepository interface {
	create(*models.News) error
	update(*models.News) error
	createLike(*models.NewsLike) error
	addTag(*models.NewsToTag) error
	removeTag(*models.NewsToTag) error
	createComment(*models.NewsComment) error
	updateComment(*models.NewsComment) error
	removeComment(string) error
	getMain() (models.NewsWithCount, error)
	getSubMain() (models.NewsWithCount, error)
	getAll() (models.NewsWithCount, error)
	delete(string) error
	deleteLike(string) error
	getBySlug(string) (*models.News, error)
	getNewsComments(string) (*models.NewsComments, error)
	createViewOfNews(*models.NewsView) error
	SetQueryFilter(*gin.Context) error
	GetSuggestionNews(id string) ([]*models.News, error)
	GetAggregateViews(*exportmodels.NewsView) (models.ChartDataSets, error)
}

type IFilesService interface {
	Upload(*gin.Context, *models.News, map[string][]*multipart.FileHeader) error
}

type Handler struct {
	service      IService
	filesService IFilesService
	helper       *helper.Helper
}

type Service struct {
	// basehandler.Service
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	// baseHandler.Repository

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
