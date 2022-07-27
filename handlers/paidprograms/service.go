package paidprograms

import (
	"mdgkb/mdgkb-server/handlers/paidprogramoptionsgroups"
	"mdgkb/mdgkb-server/handlers/paidprogramspackages"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items models.PaidPrograms) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	err = paidprogramspackages.CreateService(s.helper).CreateMany(items.GetPaidProgramPackages())
	if err != nil {
		return err
	}

	err = paidprogramoptionsgroups.CreateService(s.helper).CreateMany(items.GetPaidProgramOptionsGroups())
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
	items.SetIDForChildren()
	err = paidprogramspackages.CreateService(s.helper).UpsertMany(items.GetPaidProgramPackages())
	if err != nil {
		return err
	}
	err = paidprogramoptionsgroups.CreateService(s.helper).UpsertMany(items.GetPaidProgramOptionsGroups())
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
	item.SetIDForChildren()
	paidProgramsPackagesService := paidprogramspackages.CreateService(s.helper)
	err = paidProgramsPackagesService.UpsertMany(item.PaidProgramPackages)
	if err != nil {
		return err
	}
	err = paidProgramsPackagesService.DeleteMany(item.PaidProgramPackagesForDelete)
	if err != nil {
		return err
	}

	paidProgramOptionsGroupsService := paidprogramoptionsgroups.CreateService(s.helper)
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
