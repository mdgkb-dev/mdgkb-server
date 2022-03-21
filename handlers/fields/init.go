package fields

import (
	"context"
	"mdgkb/mdgkb-server/models"
	"mdgkb/mdgkb-server/helpers"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IService interface {
	Create(info *models.Field) error
	Update(info *models.Field) error
	Upsert(info *models.Field) error
	UpsertMany(infos models.Fields) error
}

type IRepository interface {
	getDB() *bun.DB
	create(info *models.Field) error
	update(info *models.Field) error
	upsert(info *models.Field) error
	upsertMany(infos models.Fields) error
	//deleteMany([]string) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.Field, map[string][]*multipart.FileHeader) error
}

type Handler struct {
	service IService
	filesService IFilesService
	helper       *helpers.Helper
}

type Service struct {
	repository IRepository
}

type Repository struct {
	db  *bun.DB
	ctx context.Context
}

type FilesService struct {
	helper *helpers.Helper
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

// func (s *FilesService) Upload(c *gin.Context, item *models.Field, files map[string][]*multipart.FileHeader) (err error) {
// 	for i, file := range files {
// 		err = s.helper.Uploader.Upload(c, file, item.SetFilePath(&i))
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
