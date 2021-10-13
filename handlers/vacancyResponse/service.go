package vacancyResponse

import (
	"mdgkb/mdgkb-server/handlers/contactInfo"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.VacancyResponse) error {
	err := contactInfo.CreateService(s.repository.getDB()).Create(item.ContactInfo)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	return s.repository.create(item)
}

func (s *Service) GetAll() (models.VacancyResponses, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.VacancyResponse, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.VacancyResponse) error {
	//err := contactInfo.CreateService(s.repository.getDB()).Update(item.ContactInfo)
	//if err != nil {
	//	return err
	//}
	item.SetForeignKeys()
	return s.repository.update(item)
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}
