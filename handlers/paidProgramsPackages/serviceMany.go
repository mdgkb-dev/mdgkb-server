package paidProgramsPackages

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/handlers/paidProgramServicesGroups"
	"mdgkb/mdgkb-server/handlers/paidProgramsPackagesOptions"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(items models.PaidProgramPackages) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	paidProgramsPackagesOptionsService := paidProgramsPackagesOptions.CreateService(s.repository.getDB())
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
	items.SetIdForChildren()
	paidProgramsPackagesOptionsService := paidProgramsPackagesOptions.CreateService(s.repository.getDB())
	err = paidProgramsPackagesOptionsService.UpsertMany(items.GetPaidProgramPackagesOptions())
	if err != nil {
		return err
	}
	err = paidProgramsPackagesOptionsService.DeleteMany(items.GetPaidProgramPackagesOptionsForDelete())
	if err != nil {
		return err
	}

	paidProgramServicesGroupsService := paidProgramServicesGroups.CreateService(s.repository.getDB())
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
