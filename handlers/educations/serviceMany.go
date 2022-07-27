package educations

import (
	"mdgkb/mdgkb-server/handlers/educationaccreditation"
	"mdgkb/mdgkb-server/handlers/educationcertification"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items models.Educations) error {
	if len(items) == 0 {
		return nil
	}
	err := educationcertification.CreateService(s.helper).CreateMany(items.GetEducationCertifications())
	if err != nil {
		return err
	}
	err = educationaccreditation.CreateService(s.helper).CreateMany(items.GetEducationQualification())
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
	err := educationcertification.CreateService(s.helper).UpsertMany(items.GetEducationCertifications())
	if err != nil {
		return err
	}
	err = educationaccreditation.CreateService(s.helper).UpsertMany(items.GetEducationQualification())
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
