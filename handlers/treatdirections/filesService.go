package treatdirections

import (
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func (s *FilesService) Upload(c *gin.Context, item *models.TreatDirection, files map[string][]*multipart.FileHeader) (err error) {
	return nil
}
