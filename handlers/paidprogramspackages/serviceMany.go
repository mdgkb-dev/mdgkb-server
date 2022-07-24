package paidprogramspackages

import (
	"mdgkb/mdgkb-server/handlers/paidprogramservicesgroups"
	"mdgkb/mdgkb-server/handlers/paidprogramspackagesoptions"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items models.PaidProgramPackages) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	paidProgramsPackagesOptionsService := paidprogramspackagesoptions.CreateService(s.helper)
	err = paidProgramsPackagesOptionsService.UpsertMany(items.GetPaidProgramPackagesOptions())
	if err != nil {
		return err
	}
	err = paidProgramsPackagesOptionsService.DeleteMany(items.GetPaidProgramPackagesOptionsForDelete())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.PaidProgramPackages) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	paidProgramsPackagesOptionsService := paidprogramspackagesoptions.CreateService(s.helper)
	err = paidProgramsPackagesOptionsService.UpsertMany(items.GetPaidProgramPackagesOptions())
	if err != nil {
		return err
	}
	err = paidProgramsPackagesOptionsService.DeleteMany(items.GetPaidProgramPackagesOptionsForDelete())
	if err != nil {
		return err
	}

	paidProgramServicesGroupsService := paidprogramservicesgroups.CreateService(s.helper)
	err = paidProgramServicesGroupsService.UpsertMany(items.GetPaidProgramServicesGroups())
	if err != nil {
		return err
	}
	err = paidProgramServicesGroupsService.DeleteMany(items.GetPaidProgramServicesGroupsForDelete())
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
