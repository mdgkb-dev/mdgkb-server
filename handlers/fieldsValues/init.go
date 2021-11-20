package fieldsValues

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

type IService interface {
	Create( *models.FieldValue) error
	Update( *models.FieldValue) error
	Upsert( *models.FieldValue) error
	UpsertMany( models.Fields) error
}

type IRepository interface {
	getDB() *bun.DB
	create( *models.FieldValue) error
	update( *models.FieldValue) error
	upsert( *models.FieldValue) error
	upsertMany( models.FieldValues) error
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
