package comments

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	UpdateOne(c *gin.Context)
}

type IService interface {
	CreateMany(comments models.Comments) error
	GetAll(*commentsParams) (models.Comments, error)
	UpdateOne(*models.Comment) error
}

type IRepository interface {
	getDB() *bun.DB
	createMany(comments models.Comments) error
	upsertMany(comments models.Comments) error
	deleteMany([]string) error
	getAll(*commentsParams) (models.Comments, error)
	updateOne(*models.Comment) error
}

type Handler struct {
	service IService
}

type Service struct {
	repository IRepository
}

type Repository struct {
	db  *bun.DB
	ctx context.Context
}

func CreateHandler(db *bun.DB) *Handler {
	repo := NewRepository(db)
	service := NewService(repo)
	return NewHandler(service)
}

func CreateService(db *bun.DB) *Service {
	repo := NewRepository(db)
	return NewService(repo)
}

func NewHandler(service IService) *Handler {
	return &Handler{service: service}
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db, ctx: context.Background()}
}
