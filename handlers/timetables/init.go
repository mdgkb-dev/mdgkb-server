package timetables

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)


type IHandler interface {
	GetValueTypes(c *gin.Context) error
}

type IService interface {
	Create(timetable *models.Timetable) error
	Upsert(timetable *models.Timetable) error
	GetAllWeekdays() (models.Weekdays, error)
}

type IRepository interface {
	getDB() *bun.DB
	create(timetable *models.Timetable) error
	upsert(timetable *models.Timetable) error
	getAllWeekdays() (models.Weekdays, error)
}

type Handler struct {
	service IService
}

type Service struct {
	repository IRepository
}

type Repository struct {
	db *bun.DB
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

// NewHandler constructor
func NewHandler(s IService) *Handler {
	return &Handler{service: s }
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db, ctx: context.Background()}
}


