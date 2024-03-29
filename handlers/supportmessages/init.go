package supportmessages

import (
	"context"
	"mime/multipart"

	"mdgkb/mdgkb-server/models"

	"github.com/pro-assistance/pro-assister/helper"

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
}

type IService interface {
	setQueryFilter(*gin.Context) error

	GetAll() (models.SupportMessagesWithCount, error)
	Get(string) (*models.SupportMessage, error)
	Create(*models.SupportMessage) error
	Update(*models.SupportMessage) error
	Delete(string) error

	ChangeNewStatus(string, bool) error
}

type IRepository interface {
	setQueryFilter(*gin.Context) error
	db() *bun.DB

	create(*models.SupportMessage) error
	getAll() (models.SupportMessagesWithCount, error)
	get(string) (*models.SupportMessage, error)
	update(*models.SupportMessage) error
	delete(string) error

	changeNewStatus(string, bool) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.SupportMessage, map[string][]*multipart.FileHeader) error
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
