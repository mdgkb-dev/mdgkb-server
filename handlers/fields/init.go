package fields

import (
	"context"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IService interface {
	Create(info *models.Field) error
	Update(info *models.Field) error
	Upsert(info *models.Field) error
	UpsertMany(infos models.Fields) error
	DeleteMany(uuid []uuid.UUID) error
}

type IRepository interface {
	db() *bun.DB
	create(info *models.Field) error
	update(info *models.Field) error
	upsert(info *models.Field) error
	upsertMany(infos models.Fields) error
	deleteMany([]uuid.UUID) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.Field, map[string][]*multipart.FileHeader) error
}

type Handler struct {
	service      IService
	filesService IFilesService
	helper       *helper.Helper
}

type Service struct {
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
}

type FilesService struct {
	helper *helper.Helper
}

func CreateService(h *helper.Helper) *Service {
	repo := NewRepository(h)
	return NewService(repo, h)
}

func NewService(repository IRepository, h *helper.Helper) *Service {
	return &Service{repository: repository}
}

func NewRepository(h *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: h}
}

// func (s *FilesService) Upload(c *gin.Context, item *models.Field, files map[string][]*multipart.FileHeader) (err error) {
// 	for i, file := range files {
// 		err = s.helper.Uploader.Upload(c, file, item.SetFilePath(&i))
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
