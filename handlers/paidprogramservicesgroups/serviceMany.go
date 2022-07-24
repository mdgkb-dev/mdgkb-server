package paidprogramservicesgroups

import (
	"mdgkb/mdgkb-server/handlers/paidprogramservices"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items models.PaidProgramServicesGroups) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	paidProgramServicesGroupsService := paidprogramservices.CreateService(s.helper)
	err = paidProgramServicesGroupsService.UpsertMany(items.GetPaidProgramServices())
	if err != nil {
		return err
	}
	err = paidProgramServicesGroupsService.DeleteMany(items.GetPaidProgramServicesForDelete())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.PaidProgramServicesGroups) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	paidProgramServicesGroupsService := paidprogramservices.CreateService(s.helper)
	err = paidProgramServicesGroupsService.UpsertMany(items.GetPaidProgramServices())
	if err != nil {
		return err
	}
	err = paidProgramServicesGroupsService.DeleteMany(items.GetPaidProgramServicesForDelete())
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
