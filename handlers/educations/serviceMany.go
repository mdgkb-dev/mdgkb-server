package educations

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/handlers/educationAccreditation"
	"mdgkb/mdgkb-server/handlers/educationCertification"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(items models.Educations) error {
	if len(items) == 0 {
		return nil
	}
	err := educationCertification.CreateService(s.repository.getDB()).CreateMany(items.GetEducationCertifications())
	if err != nil {
		return err
	}
	err = educationAccreditation.CreateService(s.repository.getDB()).CreateMany(items.GetEducationQualification())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	err = s.repository.createMany(items)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpsertMany(items models.Educations) error {
	if len(items) == 0 {
		return nil
	}
	err := educationCertification.CreateService(s.repository.getDB()).UpsertMany(items.GetEducationCertifications())
	if err != nil {
		return err
	}
	err = educationAccreditation.CreateService(s.repository.getDB()).UpsertMany(items.GetEducationQualification())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
	return s.repository.upsertMany(items)
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
