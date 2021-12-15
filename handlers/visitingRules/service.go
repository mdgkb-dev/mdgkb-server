package visitingRules

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) Create(item *models.VisitingRule) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.VisitingRules, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.VisitingRule, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) UpsertAndDeleteMany(items models.VisitingRulesWithDeleted) error {
	if len(items.VisitingRules) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items.VisitingRules)
	if err != nil {
		return err
	}
	if len(items.VisitingRulesForDelete) > 0 {
		err = s.repository.deleteMany(items.VisitingRulesForDelete)
	}
	return err
}

func (s *Service) UpsertMany(items models.VisitingRules) error {
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
	err := s.repository.deleteMany(idPool)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.VisitingRule) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}
