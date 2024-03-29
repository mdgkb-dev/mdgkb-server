package employees

import (
	"context"

	"mdgkb/mdgkb-server/handlers/accreditations"
	"mdgkb/mdgkb-server/handlers/certificates"
	"mdgkb/mdgkb-server/handlers/certifications"
	"mdgkb/mdgkb-server/handlers/doctors"
	"mdgkb/mdgkb-server/handlers/educationalacademics"
	"mdgkb/mdgkb-server/handlers/educations"
	"mdgkb/mdgkb-server/handlers/experiences"
	"mdgkb/mdgkb-server/handlers/heads"
	"mdgkb/mdgkb-server/handlers/human"
	"mdgkb/mdgkb-server/handlers/regalias"
	"mdgkb/mdgkb-server/handlers/teachingactivities"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.Employee) error {
	err := human.CreateService(s.helper).Create(item.Human)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.Create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	err = heads.CreateService(s.helper).Create(item.Head)
	if err != nil {
		return err
	}
	err = doctors.S.Create(context.TODO(), item.Doctor)
	if err != nil {
		return err
	}
	err = educationalacademics.CreateService(s.helper).Create(item.EducationalAcademic)
	if err != nil {
		return err
	}

	err = regalias.CreateService(s.helper).CreateMany(item.Regalias)
	if err != nil {
		return err
	}
	err = educations.CreateService(s.helper).CreateMany(item.Educations)
	if err != nil {
		return err
	}
	err = experiences.CreateService(s.helper).CreateMany(item.Experiences)
	if err != nil {
		return err
	}
	err = certificates.CreateService(s.helper).CreateMany(item.Certificates)
	if err != nil {
		return err
	}
	err = teachingactivities.CreateService(s.helper).CreateMany(item.TeachingActivities)
	if err != nil {
		return err
	}
	err = certifications.CreateService(s.helper).CreateMany(item.Certifications)
	if err != nil {
		return err
	}
	err = accreditations.CreateService(s.helper).CreateMany(item.Accreditations)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Employee) error {
	err := human.CreateService(s.helper).Update(item.Human)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.Update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	headsService := heads.CreateService(s.helper)
	err = headsService.Update(item.Head)
	if err != nil {
		return err
	}
	err = headsService.DeleteMany(item.HeadsForDelete)
	if err != nil {
		return err
	}

	err = doctors.S.Update(context.TODO(), item.Doctor)
	if err != nil {
		return err
	}
	err = doctors.S.DeleteMany(context.TODO(), item.DoctorsForDelete)
	if err != nil {
		return err
	}

	educationalAcademicsService := educationalacademics.CreateService(s.helper)
	err = educationalAcademicsService.Update(item.EducationalAcademic)
	if err != nil {
		return err
	}
	err = educationalAcademicsService.DeleteMany(item.EducationalAcademicsForDelete)
	if err != nil {
		return err
	}

	regaliasService := regalias.CreateService(s.helper)
	err = regaliasService.UpsertMany(item.Regalias)
	if err != nil {
		return err
	}
	err = regaliasService.DeleteMany(item.RegaliasForDelete)
	if err != nil {
		return err
	}

	educationsService := educations.CreateService(s.helper)
	err = educationsService.UpsertMany(item.Educations)
	if err != nil {
		return err
	}
	err = educationsService.DeleteMany(item.EducationsForDelete)
	if err != nil {
		return err
	}
	experiencesService := experiences.CreateService(s.helper)
	err = experiencesService.UpsertMany(item.Experiences)
	if err != nil {
		return err
	}
	err = experiencesService.DeleteMany(item.ExperiencesForDelete)
	if err != nil {
		return err
	}
	certificatesService := certificates.CreateService(s.helper)
	err = certificatesService.UpsertMany(item.Certificates)
	if err != nil {
		return err
	}
	err = certificatesService.DeleteMany(item.CertificatesForDelete)
	if err != nil {
		return err
	}

	teachingactivitiesService := teachingactivities.CreateService(s.helper)
	err = teachingactivitiesService.UpsertMany(item.TeachingActivities)
	if err != nil {
		return err
	}
	err = teachingactivitiesService.DeleteMany(item.TeachingActivitiesForDelete)
	if err != nil {
		return err
	}

	certificationsService := certifications.CreateService(s.helper)
	err = certificationsService.UpsertMany(item.Certifications)
	if err != nil {
		return err
	}
	err = certificationsService.DeleteMany(item.CertificationsForDelete)
	if err != nil {
		return err
	}

	accreditationsService := accreditations.CreateService(s.helper)
	err = accreditationsService.UpsertMany(item.Accreditations)
	if err != nil {
		return err
	}
	err = accreditationsService.DeleteMany(item.AccreditationsForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.EmployeesWithCount, error) {
	return s.repository.GetAll()
}

func (s *Service) Get(slug string) (*models.Employee, error) {
	item, err := s.repository.Get(slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *Service) SetQueryFilter(c *gin.Context) (err error) {
	err = s.repository.SetQueryFilter(c)
	return err
}
