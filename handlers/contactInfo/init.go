package contactInfo

import (
	"context"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

type IService interface {
	Create(*models.ContactInfo) error
	Update(*models.ContactInfo) error
	Upsert(*models.ContactInfo) error
	Delete(*string) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.ContactInfo) error
	update(*models.ContactInfo) error
	upsert(*models.ContactInfo) error
	delete(*string) error
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
