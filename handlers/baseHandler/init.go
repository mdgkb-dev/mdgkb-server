package baseHandler

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	httpHelper2 "github.com/pro-assistance/pro-assister/sqlHelper"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IService interface {
	SetQueryFilter(*gin.Context) error
}

type IRepository interface {
	SetQueryFilter(*gin.Context) error
	DB() *bun.DB
}

type IFilesService interface {
	Upload(*gin.Context, *models.DpoCourse, map[string][]*multipart.FileHeader) error
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
	ctx         context.Context
	helper      *helper.Helper
	queryFilter *httpHelper2.QueryFilter
}

type FilesService struct {
	helper *helper.Helper
}
