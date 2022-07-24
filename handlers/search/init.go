package search

import (
	"context"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/pro-assistance/pro-assister/helper"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	Search(c *gin.Context)
	ElasticSearch(c *gin.Context)
	SearchMain(c *gin.Context)
	SearchGroups(c *gin.Context)
}

type IService interface {
	SearchMain(*models.SearchModel) error
	SearchObjects(*models.SearchModel) error
	SearchGroups() (models.SearchGroups, error)
	Search(*models.SearchModel) error
}

type IRepository interface {
	db() *bun.DB
	getGroups(string) (models.SearchGroups, error)
	search(*models.SearchModel) error
	elasticSearch(*models.SearchModel) error
	elasticSuggester(*models.SearchModel) error
}

type IFilesService interface {
	Upload(*gin.Context, interface{}, map[string][]*multipart.FileHeader) error
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
	ctx           context.Context
	helper        *helper.Helper
	elasticsearch *elasticsearch.Client
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
