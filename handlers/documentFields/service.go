package documentFields

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(items models.DocumentFields) error {
	return s.repository.createMany(items)
}

func (s *Service) UpsertMany(items models.DocumentFields) error {
	return s.repository.upsertMany(items)
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
