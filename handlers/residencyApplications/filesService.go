package residencyApplications

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"
	"time"
)

func (s *FilesService) Upload(c *gin.Context, item *models.ResidencyApplication, files map[string][]*multipart.FileHeader) (err error) {
	for i, file := range files {
		err = s.helper.Uploader.Upload(c, file, item.SetFilePath(&i))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *FilesService) FillApplicationTemplate(item *models.ResidencyApplication) ([]byte, error) {
	const point = "✔"
	m := map[string]interface{}{
		"item.FormValue.User.Human.Surname":     item.FormValue.User.Human.Surname,
		"item.FormValue.User.Human.Name":        item.FormValue.User.Human.Name,
		"item.FormValue.User.Human.Patronymic":  item.FormValue.User.Human.Patronymic,
		"item.FormValue.User.Human.DateBirth":   item.FormValue.User.Human.DateBirth.Format("01.02.2006"),
		"item.FormValue.User.Human.PlaceBirth":  item.FormValue.User.Human.PlaceBirth,
		"item.FormValue.User.Human.Citizenship": item.FormValue.User.Human.Citizenship,
		"item.FormValue.User.Human.Snils":       item.FormValue.User.Human.Snils,
		"CourseName":                            item.GetCourseName(),
		"DiplomaSeries":                         item.FormValue.GetFieldValueByCode("DiplomaSeries"),
		"DiplomaNumber":                         item.FormValue.GetFieldValueByCode("DiplomaNumber"),
		"DiplomaDate":                           item.FormValue.GetFieldValueByCode("DiplomaDate").(time.Time).Format("01.02.2006"),
	}
	m["FreeApplication"] = point
	m["PaidApplication"] = ""
	if item.Paid {
		m["FreeApplication"] = ""
		m["FreeApplication"] = point
	}
	m["AdditionalApplication"] = point
	m["MainApplication"] = ""
	if item.Main {
		m["MainApplication"] = point
		m["AdditionalApplication"] = ""
	}

	return s.helper.Templater.ReplaceDoc(m, "residencyApplication.docx")
}
