package vacancies

import (
	"mdgkb/mdgkb-server/handlers/human"
	"mdgkb/mdgkb-server/handlers/vacancyResponsesToDocuments"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.Vacancy) error {
	return s.repository.create(item)
}

func (s *Service) GetAll() (models.Vacancies, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) GetAllWithResponses() (models.Vacancies, error) {
	items, err := s.repository.getAllWithResponses()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.Vacancy, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.Vacancy) error {
	return s.repository.update(item)
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) CreateResponse(item *models.VacancyResponse) error {
	err := human.CreateService(s.repository.getDB(), s.helper).Create(item.Human)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.createResponse(item)
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
