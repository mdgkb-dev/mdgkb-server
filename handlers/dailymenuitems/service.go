package dailymenuitems

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.DailyMenuItem) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.DailyMenuItem) error {
	//err := timetables.CreateService(s.helper).Upsert(item.Timetable)
	//if err != nil {
	//	return err
	//}
	//item.SetForeignKeys()
	//err = s.repository.update(item)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (s *Service) GetAll() (models.DailyMenuItems, error) {
	return s.repository.getAll()
}

func (s *Service) Get(slug string) (*models.DailyMenuItem, error) {
	item, err := s.repository.get(slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}

func (s *Service) UpsertMany(items models.DailyMenuItems) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	return nil
}
