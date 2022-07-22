package maskTokens

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"

	"github.com/uptrace/bun"
)

type IService interface {
	UpsertMany(infos models.MaskTokens) error
	DeleteMany(uuid []uuid.UUID) error
}

type IRepository interface {
	getDB() *bun.DB
	upsertMany(infos models.MaskTokens) error
	deleteMany([]uuid.UUID) error
}

type Handler struct {
	service IService
	helper       *helper.Helper
}

type Service struct {
	repository IRepository
}

type Repository struct {
	db  *bun.DB
	ctx context.Context
}

type FilesService struct {
	helper *helper.Helper
}

func CreateService(db *bun.DB) *Service {
	repo := NewRepository(db)
	return NewService(repo)
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db, ctx: context.Background()}
}
