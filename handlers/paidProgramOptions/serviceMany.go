package paidProgramOptions

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(items models.PaidProgramOptions) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) UpsertMany(items models.PaidProgramOptions) error {
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
