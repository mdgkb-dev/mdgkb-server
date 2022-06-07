package heads

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/handlers/contactInfo"
	"mdgkb/mdgkb-server/handlers/departments"
	"mdgkb/mdgkb-server/handlers/fileInfos"
	"mdgkb/mdgkb-server/handlers/human"
	"mdgkb/mdgkb-server/handlers/regalias"
	"mdgkb/mdgkb-server/handlers/timetables"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Head) error {
	err := fileInfos.CreateService(s.repository.GetDB()).Create(item.Photo)
	if err != nil {
		return err
	}
	err = contactInfo.CreateService(s.repository.GetDB()).Create(item.ContactInfo)
	if err != nil {
		return err
	}
	err = human.CreateService(s.repository.GetDB(), s.helper).Create(item.Human)
	if err != nil {
		return err
	}
	err = timetables.CreateService(s.repository.GetDB()).Create(item.Timetable)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	err = regalias.CreateService(s.repository.GetDB()).CreateMany(item.Regalias)
	if err != nil {
		return err
	}
	err = departments.CreateService(s.repository.GetDB()).CreateMany(item.Departments)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Head) error {
	err := fileInfos.CreateService(s.repository.GetDB()).Upsert(item.Photo)
	if err != nil {
		return err
	}
	err = contactInfo.CreateService(s.repository.GetDB()).Upsert(item.ContactInfo)
	if err != nil {
		return err
	}
	err = human.CreateService(s.repository.GetDB(), s.helper).Update(item.Human)
	if err != nil {
		return err
	}
	err = timetables.CreateService(s.repository.GetDB()).Upsert(item.Timetable)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	regaliasService := regalias.CreateService(s.repository.GetDB())
	err = regaliasService.UpsertMany(item.Regalias)
	if err != nil {
		return err
	}
	err = regaliasService.DeleteMany(item.RegaliasForDelete)
	if err != nil {
		return err
	}

	departmentsService := departments.CreateService(s.repository.GetDB())
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
