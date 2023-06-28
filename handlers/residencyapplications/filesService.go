package residencyapplications

import (
	"mdgkb/mdgkb-server/models"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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

func (s *FilesService) UploadFormFiles(c *gin.Context, item *models.FormValue, files map[string][]*multipart.FileHeader) (err error) {
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
		"item.FormValue.User.Human.DateBirth":   item.FormValue.User.Human.DateBirth.Format("02.01.2006"),
		"item.FormValue.User.Human.PlaceBirth":  item.FormValue.User.Human.PlaceBirth,
		"item.FormValue.User.Human.Citizenship": item.FormValue.User.Human.Citizenship,
		"item.FormValue.User.Human.Snils":       item.FormValue.User.Human.Snils,
		"item.FormValue.User.Human.PostIndex":   item.FormValue.User.Human.PostIndex,
		"item.FormValue.User.Human.Address":     item.FormValue.User.Human.Address,
		"item.FormValue.User.Email":             item.FormValue.User.Email,
		"item.FormValue.User.Phone":             item.FormValue.User.Phone,
		"CourseName":                            item.GetCourseName(),
		"DiplomaSeries":                         item.Diploma.Series,
		"DiplomaNumber":                         item.Diploma.Number,
		"DiplomaSpeciality":                     item.Diploma.Speciality,
		"DiplomaDate":                           item.Diploma.Date.Format("02.01.2006"),
		"UniversityEndYear":                     item.Diploma.UniversityEndDate.Format("2006"),
		"UniversityName":                        item.Diploma.UniversityName,
	}
	m["FreeApplication"] = point
	m["PaidApplication"] = ""
	if item.Paid {
		m["FreeApplication"] = ""
		m["PaidApplication"] = point
	}
	m["AdditionalApplication"] = point
	m["MainApplication"] = ""
	if item.Main {
		m["MainApplication"] = point
		m["AdditionalApplication"] = ""
	}

	m["PrimaryAccreditation"] = point
	m["PrimaryAccreditationPlace"] = item.PrimaryAccreditationPlace
	m["PrimaryAccreditationPoints"] = item.PrimaryAccreditationPoints
	m["PrimaryAccreditationSpecialisation"] = item.GetCourseName()

	m["PrimaryAccreditationNotPass"] = ""
	m["EntranceExamPlace"] = ""
	m["MdgkbExam"] = ""
	m["EntranceExamSpecialisation"] = ""

	if !item.PrimaryAccreditation {
		m["PrimaryAccreditation"] = ""
		m["PrimaryAccreditationPlace"] = ""
		m["PrimaryAccreditationPoints"] = ""
		m["PrimaryAccreditationSpecialisation"] = ""

		m["PrimaryAccreditationNotPass"] = point
		m["EntranceExamPlace"] = item.PrimaryAccreditationPlace

		if item.MdgkbExam {
			m["MdgkbExam"] = point
			m["EntranceExamPlace"] = ""
			m["EntranceExamSpecialisation"] = item.EntranceExamSpecialisation
		}
	}

	p := []string{}
	for i, point := range item.ResidencyApplicationPointsAchievements {
		p = append(p, strconv.Itoa(i+1)+". "+point.PointsAchievement.Name)
	}
	m["PointsAchievements"] = strings.Join(p, "\n")
	return s.helper.Templater.ReplaceDoc(m, "residencyApplication.docx")
}
