package search

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"mdgkb/mdgkb-server/helpers"
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
	helper       *helpers.Helper
}

type Service struct {
	repository IRepository
	helper     *helpers.Helper
}

type Repository struct {
	db            *bun.DB
	ctx           context.Context
	helper        *helpers.Helper
	elasticsearch *elasticsearch.Client
}

type FilesService struct {
	helper *helpers.Helper
}

func CreateHandler(db *bun.DB, helper *helpers.Helper, elasticSearchClient *elasticsearch.Client) *Handler {
	repo := NewRepository(db, helper, elasticSearchClient)
	service := NewService(repo, helper)
	filesService := NewFilesService(helper)
	return NewHandler(service, filesService, helper)
}

// NewHandler constructor
func NewHandler(s IService, filesService IFilesService, helper *helpers.Helper) *Handler {
	return &Handler{service: s, filesService: filesService, helper: helper}
}

func NewService(repository IRepository, helper *helpers.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helpers.Helper, elasticSearchClient *elasticsearch.Client) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper, elasticsearch: elasticSearchClient}
}

func NewFilesService(helper *helpers.Helper) *FilesService {
	return &FilesService{helper: helper}
}
