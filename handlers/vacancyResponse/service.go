package vacancyResponse

import (
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/handlers/vacancyResponsesToDocuments"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.VacancyResponse) error {
	err := users.CreateService(s.repository.getDB(), s.helper).UpsertEmail(item.User)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = vacancyResponsesToDocuments.CreateService(s.repository.getDB()).CreateMany(item.VacancyResponsesToDocuments)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.VacancyResponses, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.VacancyResponse, error) {
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

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}
