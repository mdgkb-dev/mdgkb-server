package vacancies

import (
	"mdgkb/mdgkb-server/models"
	"mime/multipart"
	"github.com/gin-gonic/gin"
)

func (s *FilesService) Upload(c *gin.Context, item *models.VacancyResponse, files map[string][]*multipart.FileHeader) (err error) {
	for i, file := range files {
		err := s.helper.Uploader.Upload(c, file, item.SetFilePath(&i))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *FilesService) UploadVacancy(c *gin.Context, item *models.Vacancy, files map[string][]*multipart.FileHeader) (err error) {
	for i, file := range files {
		err := s.helper.Uploader.Upload(c, file, item.SetFilePath(&i))
		if err != nil {
			return err
		}
	}
	return nil
}
