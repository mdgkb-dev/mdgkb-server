package paidprogramservices

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items models.PaidProgramServices) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpsertMany(items models.PaidProgramServices) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
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
