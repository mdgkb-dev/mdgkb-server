package timetables

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAllWeekdays(c *gin.Context)
}

type IService interface {
	Create(timetable *models.Timetable) error
	Upsert(timetable *models.Timetable) error
	GetAllWeekdays() (models.Weekdays, error)
}

type IRepository interface {
	db() *bun.DB
	create(timetable *models.Timetable) error
	upsert(timetable *models.Timetable) error
	getAllWeekdays() (models.Weekdays, error)
}

type Handler struct {
	service IService
	helper  *helper.Helper
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
	return NewHandler(service, h)
}

func CreateService(h *helper.Helper) *Service {
	repo := NewRepository(h)
	return NewService(repo, h)
}

// NewHandler constructor
func NewHandler(s IService, h *helper.Helper) *Handler {
	return &Handler{service: s, helper: h}
}

func NewService(repository IRepository, h *helper.Helper) *Service {
	return &Service{repository: repository}
}

func NewRepository(h *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: h}
}
