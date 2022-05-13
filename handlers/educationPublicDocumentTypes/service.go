package educationPublicDocumentTypes

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) GetAll() (models.EducationPublicDocumentTypes, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.EducationPublicDocumentType, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.EducationPublicDocumentType) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.EducationPublicDocumentType) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Upsert(item *models.EducationPublicDocumentType) error {
	if item == nil {
		return nil
	}
	err := s.repository.upsert(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteByPublicDocumentTypeID(id uuid.NullUUID) error {
	if !id.Valid {
		return nil
	}
	return s.repository.deleteByPublicDocumentTypeID(id)
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}
