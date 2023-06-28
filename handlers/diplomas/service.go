package diplomas

import (
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Upsert(item *models.Diploma) error {
	err := s.repository.upsert(item)
	if err != nil {
		return err
	}
	return nil
}
