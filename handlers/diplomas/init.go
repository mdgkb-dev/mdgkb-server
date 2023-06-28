package diplomas

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

type IService interface {
	Upsert(info *models.Diploma) error
}

type IRepository interface {
	db() *bun.DB
	upsert(info *models.Diploma) error
}

type Service struct {
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
}

func CreateService(helper *helper.Helper) *Service {
	repo := NewRepository(helper)
	return NewService(repo, helper)
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}
