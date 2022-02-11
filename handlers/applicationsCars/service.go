package applicationsCars

import (
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll() (models.ApplicationsCars, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.ApplicationCar, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.ApplicationCar) error {
	err := users.CreateService(s.repository.getDB(), s.helper).Upsert(item.User)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.ApplicationCar) error {
	err := users.CreateService(s.repository.getDB(), s.helper).Upsert(item.User)
	if err != nil {
		return err
	}
	item.SetForeignKeys()
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}
