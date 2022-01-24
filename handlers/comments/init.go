package comments

import (
	"context"
	"mdgkb/mdgkb-server/helpers"
	httpHelper2 "mdgkb/mdgkb-server/helpers/httpHelperV2"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	UpdateOne(c *gin.Context)
}

type IService interface {
	setQueryFilter(*gin.Context) error
	CreateMany(comments models.Comments) error
	GetAll(*commentsParams) (models.Comments, error)
	UpdateOne(*models.Comment) error
}

type IRepository interface {
	setQueryFilter(*gin.Context) error
	getDB() *bun.DB
	createMany(comments models.Comments) error
	upsertMany(comments models.Comments) error
	deleteMany([]string) error
	getAll(*commentsParams) (models.Comments, error)
	updateOne(*models.Comment) error
}

type Handler struct {
	service IService
	helper  *helpers.Helper
}

type Service struct {
	repository IRepository
	helper     *helpers.Helper
}

type Repository struct {
	db          *bun.DB
	ctx         context.Context
	helper      *helpers.Helper
	queryFilter *httpHelper2.QueryFilter
}

func CreateHandler(db *bun.DB, helper *helpers.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	return NewHandler(service, helper)
}
func CreateService(db *bun.DB, helper *helpers.Helper) *Service {
	repo := NewRepository(db, helper)
	return NewService(repo, helper)
}

// NewHandler constructor
func NewHandler(s IService, helper *helpers.Helper) *Handler {
	return &Handler{service: s, helper: helper}
}

func NewService(repository IRepository, helper *helpers.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helpers.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}
