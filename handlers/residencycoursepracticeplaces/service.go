package residencycoursepracticeplaces

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (s *Service) CreateMany(items models.ResidencyCoursePracticePlaces) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.ResidencyCoursePracticePlaces) error {
	if len(items) == 0 {
		return nil
	}
	items.SetForeignKeys()
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.ResidencyCoursePracticePlace)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
