package comments

import (
	"context"
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"
	httpHelper2 "github.com/pro-assistance/pro-assister/sqlHelper"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	GetAllMain(c *gin.Context)
	UpdateOne(c *gin.Context)
	UpsertOne(c *gin.Context)
}

type IService interface {
	setQueryFilter(*gin.Context) error
	CreateMany(comments models.Comments) error
	GetAll() (models.CommentsWithCount, error)
	GetAllMain() (models.Comments, error)
	UpdateOne(*models.Comment) error
	UpsertOne(*models.Comment) error
}

type IRepository interface {
	setQueryFilter(*gin.Context) error
	getDB() *bun.DB
	get(uuid.UUID) (models.Comment, error)
	createMany(comments models.Comments) error
	upsertMany(comments models.Comments) error
	deleteMany([]string) error
	getAll() (models.CommentsWithCount, error)
	getAllMain() (models.Comments, error)
	updateOne(*models.Comment) error
	upsertOne(*models.Comment) error
}

type Handler struct {
	service IService
	helper  *helper.Helper
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

func CreateHandler(db *bun.DB, helper *helper.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	return NewHandler(service, helper)
}
func CreateService(db *bun.DB, helper *helper.Helper) *Service {
	repo := NewRepository(db, helper)
	return NewService(repo, helper)
}

// NewHandler constructor
func NewHandler(s IService, helper *helper.Helper) *Handler {
	return &Handler{service: s, helper: helper}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helper.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}
