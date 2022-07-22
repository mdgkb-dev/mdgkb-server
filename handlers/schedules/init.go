package schedules

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetValueTypes(c *gin.Context) error
}

type IService interface {
	Create(timetable *models.Schedule) error
	Upsert(timetable *models.Schedule) error
}

type IRepository interface {
	db() *bun.DB
	create(timetable *models.Schedule) error
	upsert(timetable *models.Schedule) error
}

type Handler struct {
	service IService
}

type Service struct {
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
}

func CreateHandler(h *helper.Helper) *Handler {
	repo := NewRepository(h)
	service := NewService(repo, h)
	return NewHandler(service)
}

func CreateService(h *helper.Helper) *Service {
	repo := NewRepository(h)
	return NewService(repo, h)
}

// NewHandler constructor
func NewHandler(s IService) *Handler {
	return &Handler{service: s}
}

func NewService(repository IRepository, h *helper.Helper) *Service {
	return &Service{repository: repository}
}

func NewRepository(h *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: h}
}
