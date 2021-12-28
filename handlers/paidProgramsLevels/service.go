package paidProgramsLevels

import (
	"mdgkb/mdgkb-server/handlers/paidProgramServices"
	"mdgkb/mdgkb-server/handlers/paidPrograms"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.PaidProgramsGroup) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	err = paidPrograms.CreateService(s.repository.getDB()).CreateMany(item.PaidPrograms)
	if err != nil {
		return err
	}

	err = paidProgramServices.CreateService(s.repository.getDB()).CreateMany(item.PaidProgramServices)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.PaidProgramsGroup) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	paidProgramsService := paidPrograms.CreateService(s.repository.getDB())
	err = paidProgramsService.UpsertMany(item.PaidPrograms)
	if err != nil {
		return err
	}
	err = paidProgramsService.DeleteMany(item.PaidProgramServicesForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.PaidProgramsGroups, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.PaidProgramsGroup, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}
