package paidprogramoptionsgroups

import (
	"mdgkb/mdgkb-server/handlers/paidprogramoptions"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items models.PaidProgramOptionsGroups) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	items.SetIDForChildren()
	if err != nil {
		return err
	}
	err = paidprogramoptions.CreateService(s.helper).CreateMany(items.GetPaidProgramOptions())
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
	items.SetIDForChildren()
	paidProgramOptionsService := paidprogramoptions.CreateService(s.helper)
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
