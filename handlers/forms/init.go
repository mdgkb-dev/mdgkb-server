package forms

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

type IService interface {
	Create(info *models.Form) error
	Update(info *models.Form) error
	Upsert(info *models.Form) error
	UpsertMany(infos models.Forms) error
}

type IRepository interface {
	getDB() *bun.DB
	create(info *models.Form) error
	update(info *models.Form) error
	upsert(info *models.Form) error
	upsertMany(infos models.Forms) error
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
