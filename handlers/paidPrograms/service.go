package paidPrograms

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/handlers/paidProgramOptionsGroups"
	"mdgkb/mdgkb-server/handlers/paidProgramsPackages"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(items models.PaidPrograms) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	err = paidProgramsPackages.CreateService(s.repository.getDB()).CreateMany(items.GetPaidProgramPackages())
	if err != nil {
		return err
	}

	err = paidProgramOptionsGroups.CreateService(s.repository.getDB()).CreateMany(items.GetPaidProgramOptionsGroups())
	if err != nil {
		return err
	}
	return err
}

func (s *Service) UpsertMany(items models.PaidPrograms) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	err = paidProgramsPackages.CreateService(s.repository.getDB()).UpsertMany(items.GetPaidProgramPackages())
	if err != nil {
		return err
	}
	err = paidProgramOptionsGroups.CreateService(s.repository.getDB()).UpsertMany(items.GetPaidProgramOptionsGroups())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}

func (s *Service) Get(id string) (*models.PaidProgram, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.PaidProgram) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	paidProgramsPackagesService := paidProgramsPackages.CreateService(s.repository.getDB())
	err = paidProgramsPackagesService.UpsertMany(item.PaidProgramPackages)
	if err != nil {
		return err
	}
	err = paidProgramsPackagesService.DeleteMany(item.PaidProgramPackagesForDelete)
	if err != nil {
		return err
	}

	paidProgramOptionsGroupsService := paidProgramOptionsGroups.CreateService(s.repository.getDB())
	err = paidProgramOptionsGroupsService.UpsertMany(item.PaidProgramOptionsGroups)
	if err != nil {
		return err
	}
	err = paidProgramOptionsGroupsService.DeleteMany(item.PaidProgramOptionsGroupsForDelete)
	if err != nil {
		return err
	}
	return nil
}
