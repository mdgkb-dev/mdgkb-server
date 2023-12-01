package fileinfos

import (
	"context"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/pro-assistance/pro-assister/helper"

	"github.com/uptrace/bun"
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Create(c *gin.Context)
}
type IService interface {
	Create(info *models.FileInfo) error
	Update(info *models.FileInfo) error
	Upsert(info *models.FileInfo) error
	UpsertMany(infos models.FileInfos) error
}

type IRepository interface {
	db() *bun.DB
	create(info *models.FileInfo) error
	update(info *models.FileInfo) error
	upsert(info *models.FileInfo) error
	upsertMany(infos models.FileInfos) error
	//deleteMany([]string) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.FileInfo, map[string][]*multipart.FileHeader) error
}

type Handler struct {
	service IService
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

func CreateHandler(helper *helper.Helper) *Handler {
	repo := NewRepository(helper)
	service := NewService(repo, helper)
	filesService := NewFilesService(helper)
	return NewHandler(service, filesService, helper)
}

func CreateService(h *helper.Helper) *Service {
	repo := NewRepository(h)
	return NewService(repo, h)
}

func NewHandler(s IService, filesService IFilesService, helper *helper.Helper) *Handler {
	return &Handler{service: s, filesService: filesService, helper: helper}
}

func NewService(repository IRepository, h *helper.Helper) *Service {
	return &Service{repository: repository, helper: h}
}

func NewRepository(h *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: h}
}

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}
