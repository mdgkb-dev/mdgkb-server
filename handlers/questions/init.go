package questions

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	httpHelper2 "github.com/pro-assistance/pro-assister/sqlHelper"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)

	ChangeNewStatus(c *gin.Context)
	ReadAnswers(c *gin.Context)
	Publish(c *gin.Context)
}

type IService interface {
	setQueryFilter(*gin.Context) error

	GetAll() (models.QuestionsWithCount, error)
	Get(string) (*models.Question, error)
	Create(*models.Question) error
	Update(*models.Question) error
	Delete(string) error

	ChangeNewStatus(string, bool) error
	ReadAnswers(string) error
	Publish(string) error
}

type IRepository interface {
	setQueryFilter(*gin.Context) error
	getDB() *bun.DB

	create(*models.Question) error
	getAll() (models.QuestionsWithCount, error)
	get(string) (*models.Question, error)
	update(*models.Question) error
	delete(string) error

	changeNewStatus(string, bool) error
	readAnswers(string) error
	publish(string) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.Question, map[string][]*multipart.FileHeader) error
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
	db          *bun.DB
	ctx         context.Context
	helper      *helper.Helper
	queryFilter *httpHelper2.QueryFilter
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
