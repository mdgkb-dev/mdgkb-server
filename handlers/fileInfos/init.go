package fileInfos

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

type IService interface {
	Create(info *models.FileInfo) error
	Update(info *models.FileInfo) error
	Upsert(info *models.FileInfo) error
	UpsertMany(infos models.FileInfos) error
}

type IRepository interface {
	getDB() *bun.DB
	create(info *models.FileInfo) error
	update(info *models.FileInfo) error
	upsert(info *models.FileInfo) error
	upsertMany(infos models.FileInfos) error
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
