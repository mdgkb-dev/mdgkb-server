package paidProgramOptionsGroups

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/handlers/paidProgramOptions"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(items models.PaidProgramOptionsGroups) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	items.SetIdForChildren()
	if err != nil {
		return err
	}
	err = paidProgramOptions.CreateService(s.repository.getDB()).CreateMany(items.GetPaidProgramOptions())
	if err != nil {
		return err
	}
	return err
}

func (s *Service) UpsertMany(items models.PaidProgramOptionsGroups) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	paidProgramOptionsService := paidProgramOptions.CreateService(s.repository.getDB())
	err = paidProgramOptionsService.UpsertMany(items.GetPaidProgramOptions())
	if err != nil {
		return err
	}
	err = paidProgramOptionsService.DeleteMany(items.PaidProgramOptionsForDelete())
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
