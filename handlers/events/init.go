package events

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

type IService interface {
	Create(info *models.Event) error
	Update(info *models.Event) error
	Upsert(info *models.Event) error
	UpsertMany(infos models.Events) error
}

type IRepository interface {
	getDB() *bun.DB
	create(info *models.Event) error
	update(info *models.Event) error
	upsert(info *models.Event) error
	upsertMany(infos models.Events) error
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
