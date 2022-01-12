package publicDocumentTypes

import (
	"fmt"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func (s *FilesService) Upload(c *gin.Context, item *models.PublicDocumentType, files map[string][]*multipart.FileHeader) (err error) {
	fmt.Println("files ===>", files)
	for i, file := range files {
		err = s.helper.Uploader.Upload(c, file, item.SetFilePath(&i))
		if err != nil {
			return err
		}
	}
	return nil
}
