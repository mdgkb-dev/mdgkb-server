package partners

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func (s *FilesService) Upload(c *gin.Context, item *models.Partner, files map[string][]*multipart.FileHeader) (err error) {
	for i, file := range files {
		path := item.SetFilePath(&i)
		if path == nil {
			return errors.New("file does not relate to anything")
		}
		uploadPath := "/home/lakkinzimusic/prog/mdgkb/mdgkb-server/static"
		pathDirs := strings.Split(*path, string(os.PathSeparator))
		pathToFile := filepath.Join(uploadPath, filepath.Join(pathDirs[:len(pathDirs)-1]...))
		err = os.MkdirAll(pathToFile, os.ModePerm)
		if err != nil {
			return err
		}

		fullPath := filepath.Join(uploadPath, *path)
		err = c.SaveUploadedFile(file[0], fullPath)
		if err != nil {
			return err
		}
	}
	return nil
}
