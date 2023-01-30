package teachers

import (
	"mdgkb/mdgkb-server/handlers/employees"
	"mdgkb/mdgkb-server/handlers/human"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetAll() (models.Teachers, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

// func (s *Service) Get(id *string) (*models.Teacher, error) {
// 	item, err := s.repository.get(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return item, nil
// }

func (s *Service) Get(slug string) (*models.Teacher, error) {
	item, err := s.repository.get(slug)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.Teacher) error {
	// err := employees.CreateService(s.helper).Create(item.Employee)
	// if err != nil {
	// 	return err
	// }
	item.SetForeignKeys()
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.Teacher) error {
	item.SetForeignKeys()
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	err = employees.CreateService(s.helper).Update(item.Employee)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) UpsertMany(items models.Teachers) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.upsertMany(items)
}

func (s *Service) DeleteMany(id []string) error {
	if len(id) == 0 {
		return nil
	}
	return s.repository.deleteMany(id)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
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
