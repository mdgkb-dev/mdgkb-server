package children

import (
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func (s *FilesService) Upload(c *gin.Context, item *models.Child, files map[string][]*multipart.FileHeader) (err error) {
	//for i, file := range files {
	//err = s.helper.Uploader.Upload(c, file, item.SetFilePath(&i))
	//if err != nil {
	//	return err
	//}
	//}
	return nil
}
