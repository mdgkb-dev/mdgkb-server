package residencyApplications

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"
	"strconv"
	"strings"
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
	const point = `✓`
	m := map[string]interface{}{
		"item.FormValue.User.Human.Surname":     item.FormValue.User.Human.Surname,
		"item.FormValue.User.Human.Name":        item.FormValue.User.Human.Name,
		"item.FormValue.User.Human.Patronymic":  item.FormValue.User.Human.Patronymic,
		"item.FormValue.User.Human.DateBirth":   item.FormValue.User.Human.DateBirth.Format("01.02.2006"),
		"item.FormValue.User.Human.PlaceBirth":  item.FormValue.User.Human.PlaceBirth,
		"item.FormValue.User.Human.Citizenship": item.FormValue.User.Human.Citizenship,
		"item.FormValue.User.Human.Snils":       item.FormValue.User.Human.Snils,
		"item.FormValue.User.Human.PostIndex":   item.FormValue.User.Human.PostIndex,
		"item.FormValue.User.Human.Address":     item.FormValue.User.Human.Address,
		"item.FormValue.User.Email":             item.FormValue.User.Email,
		"item.FormValue.User.Phone":             item.FormValue.User.Phone,
		"CourseName":                            item.GetCourseName(),
		"DiplomaSeries":                         item.FormValue.GetFieldValueByCode("DiplomaSeries"),
		"DiplomaNumber":                         item.FormValue.GetFieldValueByCode("DiplomaNumber"),
		"DiplomaDate":                           item.FormValue.GetFieldValueByCode("DiplomaDate").(*time.Time).Format("01.02.2006"),
		"UniversityEndYear":                     item.FormValue.GetFieldValueByCode("UniversityEndYear").(*time.Time).Format("2006"),
		"UniversityName":                        item.FormValue.GetFieldValueByCode("UniversityName").(string),
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

	m["PrimaryAccreditation"] = ""
	m["PrimaryAccreditationNotPass"] = point
	m["PrimaryAccreditationPlace"] = ""
	m["PrimaryAccreditationPoints"] = ""
	if item.PrimaryAccreditation {
		m["PrimaryAccreditation"] = point
		m["PrimaryAccreditationNotPass"] = ""
		m["PrimaryAccreditationPoints"] = item.PrimaryAccreditationPoints
		m["PrimaryAccreditationPlace"] = item.PrimaryAccreditationPlace
	}
	p := []string{}
	for i, point := range item.ResidencyApplicationPointsAchievements {
		p = append(p, strconv.Itoa(i+1)+". "+point.PointsAchievement.Name)
	}
	m["PointsAchievements"] = strings.Join(p, "\n")
	fmt.Println(m["PointsAchievements"])
	return s.helper.Templater.ReplaceDoc(m, "residencyApplication.docx")
}
