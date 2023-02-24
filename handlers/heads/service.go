package heads

import (
	"mdgkb/mdgkb-server/handlers/contactinfo"
	"mdgkb/mdgkb-server/handlers/departments"
	"mdgkb/mdgkb-server/handlers/timetables"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Service) Create(item *models.Head) error {
	if item == nil {
		return nil
	}
	err := contactinfo.CreateService(s.helper).Create(item.ContactInfo)
	if err != nil {
		return err
	}
	err = timetables.CreateService(s.helper).Create(item.Timetable)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.upsert(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	err = departments.CreateService(s.helper).CreateMany(item.Departments)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Head) error {
	if item == nil {
		return nil
	}
	err := contactinfo.CreateService(s.helper).Upsert(item.ContactInfo)
	if err != nil {
		return err
	}
	err = timetables.CreateService(s.helper).Upsert(item.Timetable)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.upsert(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	departmentsService := departments.CreateService(s.helper)
	err = departmentsService.UpsertMany(item.Departments)
	if err != nil {
		return err
	}
	err = departmentsService.DeleteMany(item.DepartmentsForDelete)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAll() (models.Heads, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.Head, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) SetQueryFilter(c *gin.Context) error {
	return s.repository.SetQueryFilter(c)
}

func (s *Service) UpdateAll(items models.Heads) error {
	return s.repository.updateAll(items)
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
