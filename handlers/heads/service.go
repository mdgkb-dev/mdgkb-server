package heads

import (
	"mdgkb/mdgkb-server/handlers/contactinfo"
	"mdgkb/mdgkb-server/handlers/departments"
	"mdgkb/mdgkb-server/handlers/fileinfos"
	"mdgkb/mdgkb-server/handlers/human"
	"mdgkb/mdgkb-server/handlers/regalias"
	"mdgkb/mdgkb-server/handlers/timetables"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.Head) error {
	err := fileinfos.CreateService(s.helper).Create(item.Photo)
	if err != nil {
		return err
	}
	err = contactinfo.CreateService(s.helper).Create(item.ContactInfo)
	if err != nil {
		return err
	}
	err = human.CreateService(s.helper).Create(item.Human)
	if err != nil {
		return err
	}
	err = timetables.CreateService(s.helper).Create(item.Timetable)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	err = regalias.CreateService(s.helper).CreateMany(item.Regalias)
	if err != nil {
		return err
	}
	err = departments.CreateService(s.helper).CreateMany(item.Departments)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Head) error {
	err := fileinfos.CreateService(s.helper).Upsert(item.Photo)
	if err != nil {
		return err
	}
	err = contactinfo.CreateService(s.helper).Upsert(item.ContactInfo)
	if err != nil {
		return err
	}
	err = human.CreateService(s.helper).Update(item.Human)
	if err != nil {
		return err
	}
	err = timetables.CreateService(s.helper).Upsert(item.Timetable)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	regaliasService := regalias.CreateService(s.helper)
	err = regaliasService.UpsertMany(item.Regalias)
	if err != nil {
		return err
	}
	err = regaliasService.DeleteMany(item.RegaliasForDelete)
	if err != nil {
		return err
	}

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
