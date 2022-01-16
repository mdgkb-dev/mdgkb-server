package doctors

import (
	"github.com/gin-gonic/gin"
	certificates "mdgkb/mdgkb-server/handlers/certifiactes"
	"mdgkb/mdgkb-server/handlers/doctorPaidServices"
	"mdgkb/mdgkb-server/handlers/educations"
	"mdgkb/mdgkb-server/handlers/experiences"
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/handlers/human"
	"mdgkb/mdgkb-server/handlers/regalias"
	"mdgkb/mdgkb-server/handlers/timetables"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Doctor) error {
	err := fileInfos.CreateService(s.repository.getDB()).Create(item.FileInfo)
	if err != nil {
		return err
	}
	err = human.CreateService(s.repository.getDB(), s.helper).Create(item.Human)
	if err != nil {
		return err
	}
	err = timetables.CreateService(s.repository.getDB()).Create(item.Timetable)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	err = regalias.CreateService(s.repository.getDB()).CreateMany(item.Regalias)
	if err != nil {
		return err
	}
	err = educations.CreateService(s.repository.getDB()).CreateMany(item.Educations)
	if err != nil {
		return err
	}
	err = experiences.CreateService(s.repository.getDB()).CreateMany(item.Experiences)
	if err != nil {
		return err
	}
	err = certificates.CreateService(s.repository.getDB()).CreateMany(item.Certificates)
	if err != nil {
		return err
	}
	err = doctorPaidServices.CreateService(s.repository.getDB()).CreateMany(item.DoctorPaidServices)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Doctor) error {
	err := fileInfos.CreateService(s.repository.getDB()).Upsert(item.FileInfo)
	if err != nil {
		return err
	}
	err = human.CreateService(s.repository.getDB(), s.helper).Update(item.Human)
	if err != nil {
		return err
	}
	err = timetables.CreateService(s.repository.getDB()).Upsert(item.Timetable)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	doctorRegaliaService := regalias.CreateService(s.repository.getDB())
	err = doctorRegaliaService.UpsertMany(item.Regalias)
	if err != nil {
		return err
	}
	err = doctorRegaliaService.DeleteMany(item.RegaliasForDelete)
	if err != nil {
		return err
	}
	educationsService := educations.CreateService(s.repository.getDB())
	err = educationsService.UpsertMany(item.Educations)
	if err != nil {
		return err
	}
	err = educationsService.DeleteMany(item.EducationsForDelete)
	if err != nil {
		return err
	}
	experiencesService := experiences.CreateService(s.repository.getDB())
	err = experiencesService.UpsertMany(item.Experiences)
	if err != nil {
		return err
	}
	err = experiencesService.DeleteMany(item.ExperiencesForDelete)
	if err != nil {
		return err
	}
	certificatesService := certificates.CreateService(s.repository.getDB())
	err = certificatesService.UpsertMany(item.Certificates)
	if err != nil {
		return err
	}
	err = certificatesService.DeleteMany(item.CertificatesForDelete)
	if err != nil {
		return err
	}

	doctorPaidServicesService := doctorPaidServices.CreateService(s.repository.getDB())
	err = doctorPaidServicesService.UpsertMany(item.DoctorPaidServices)
	if err != nil {
		return err
	}
	err = doctorPaidServicesService.DeleteMany(item.DoctorPaidServicesForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(params *doctorsParams) (models.DoctorsWithCount, error) {
	return s.repository.getAll(params)
}

func (s *Service) Get(slug string) (*models.Doctor, error) {
	item, err := s.repository.get(slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) GetByDivisionID(divisionID string) (models.Doctors, error) {
	return s.repository.getByDivisionID(divisionID)
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) CreateComment(item *models.DoctorComment) error {
	return s.repository.createComment(item)
}

func (s *Service) UpdateComment(item *models.DoctorComment) error {
	return s.repository.updateComment(item)
}

func (s *Service) RemoveComment(id string) error {
	return s.repository.removeComment(id)
}

func (s *Service) UpsertMany(items models.Doctors) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.upsertMany(items)
}

func (s *Service) CreateSlugs() error {
	_, err := s.repository.getAll(nil)
	if err != nil {
		return err
	}
	humans := make(models.Humans, 0)
	//for i := range items {
	//	items[i].Human.Slug = s.helper.MakeSlug(items[i].Human.GetFullName())
	//	humans = append(humans, items[i].Human)
	//}
	err = human.CreateService(s.repository.getDB(), s.helper).UpsertMany(humans)
	return err
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}

func (s *Service) Search(query string) (models.Doctors, error) {
	queryRu := s.helper.TranslitToRu(query)
	return s.repository.search(queryRu)
}
