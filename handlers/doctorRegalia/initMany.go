package doctorRegalia

import (
	"context"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

type IService interface {
	CreateMany(models.DoctorRegalias) error
	UpsertMany(models.DoctorRegalias) error
	DeleteMany([]string) error
}

type IRepository interface {
	createMany(models.DoctorRegalias) error
	upsertMany(models.DoctorRegalias) error
	deleteMany([]string) error
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
