package search

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

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
	ElasticSearch(*models.SearchModel) error
}

type IRepository interface {
	getDB() *bun.DB
	getGroups(string) (models.SearchGroups, error)
	search(*models.SearchGroup, string) error
	elasticSearch(*models.SearchModel) error
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
	db            *bun.DB
	ctx           context.Context
	helper        *helper.Helper
	elasticsearch *elasticsearch.Client
}

type FilesService struct {
	helper *helper.Helper
}

func CreateHandler(db *bun.DB, helper *helper.Helper, elasticSearchClient *elasticsearch.Client) *Handler {
	repo := NewRepository(db, helper, elasticSearchClient)
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

func NewRepository(db *bun.DB, helper *helper.Helper, elasticSearchClient *elasticsearch.Client) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper, elasticsearch: elasticSearchClient}
}

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}
