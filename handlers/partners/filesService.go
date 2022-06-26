package partners

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"
)

func (s *FilesService) Upload(c *gin.Context, item *models.Partner, files map[string][]*multipart.FileHeader) (err error) {
	for i, file := range files {
		err = s.helper.Uploader.Upload(c, file, item.SetFilePath(&i))
		if err != nil {
			return err
		}
	}
	return nil
}
