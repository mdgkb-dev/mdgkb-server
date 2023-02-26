package doctors

import (
	"mdgkb/mdgkb-server/handlers/comments"
	"mdgkb/mdgkb-server/handlers/doctorpaidservices"
	"mdgkb/mdgkb-server/handlers/doctorsdivisions"
	"mdgkb/mdgkb-server/handlers/human"
	"mdgkb/mdgkb-server/handlers/timetables"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Service) Create(item *models.Doctor) error {
	if item == nil {
		return nil
	}
	err := timetables.CreateService(s.helper).Create(item.Timetable)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	if err != nil {
		return err
	}
	err = doctorpaidservices.CreateService(s.helper).CreateMany(item.DoctorPaidServices)
	if err != nil {
		return err
	}
	err = doctorsdivisions.CreateService(s.helper).CreateMany(item.DoctorsDivisions)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Doctor) error {
	if item == nil {
		return nil
	}
	err := timetables.CreateService(s.helper).Upsert(item.Timetable)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.upsert(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	doctorPaidServicesService := doctorpaidservices.CreateService(s.helper)
	err = doctorPaidServicesService.UpsertMany(item.DoctorPaidServices)
	if err != nil {
		return err
	}
	err = doctorPaidServicesService.DeleteMany(item.DoctorPaidServicesForDelete)
	if err != nil {
		return err
	}
	doctorsDivisionsService := doctorsdivisions.CreateService(s.helper)
	err = doctorsDivisionsService.UpsertMany(item.DoctorsDivisions)
	if err != nil {
		return err
	}
	err = doctorsDivisionsService.DeleteMany(item.DoctorsDivisionsForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.Doctors, error) {
	return s.repository.getAll()
}

func (s *Service) GetAllTimetables() (models.Doctors, error) {
	return s.repository.getAllTimetables()
}

func (s *Service) GetAllAdmin() (models.DoctorsWithCount, error) {
	return s.repository.getAllAdmin()
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
	commentsService := comments.CreateService(s.helper)
	err := commentsService.UpsertOne(item.Comment)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	return s.repository.createComment(item)
}

func (s *Service) UpdateComment(item *models.DoctorComment) error {
	commentsService := comments.CreateService(s.helper)
	err := commentsService.UpdateOne(item.Comment)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
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
	_, err := s.repository.getAll()
	if err != nil {
		return err
	}
	humans := make(models.Humans, 0)
	//for i := range items {
	//	items[i].Human.Slug = s.helper.Util.MakeSlug(items[i].Human.GetFullName())
	//	humans = append(humans, items[i].Human)
	//}
	err = human.CreateService(s.helper).UpsertMany(humans)
	return err
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}

func (s *Service) Search(query string) (models.Doctors, error) {
	queryRu := s.helper.Util.TranslitToRu(query)
	return s.repository.search(queryRu)
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
