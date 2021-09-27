package educationalOrganization

import (
	"fmt"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func (s *FilesService) Upload(c *gin.Context, item *models.EducationalOrganization, files map[string][]*multipart.FileHeader) (err error) {
	fmt.Println(1)
	for i, file := range files {
		err = s.uploader.Upload(c, file, item.SetFilePath(&i))
		if err != nil {
			return err
		}
	}
	return nil
}
