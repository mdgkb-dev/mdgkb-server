package fields

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

type IService interface {
	Create(info *models.Field) error
	Update(info *models.Field) error
	Upsert(info *models.Field) error
	UpsertMany(infos models.Fields) error
}

type IRepository interface {
	getDB() *bun.DB
	create(info *models.Field) error
	update(info *models.Field) error
	upsert(info *models.Field) error
	upsertMany(infos models.Fields) error
	//deleteMany([]string) error
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
