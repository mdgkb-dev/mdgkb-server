package vacancyResponsesToDocuments

import (
	"mdgkb/mdgkb-server/handlers/documents"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(items models.VacancyResponsesToDocuments) error {
	if len(items) == 0 {
		return nil
	}

	err := documents.CreateService(s.repository.getDB()).CreateMany(items.GetDocuments())
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

func (s *Service) UpsertMany(items models.VacancyResponsesToDocuments) error {
	if len(items) == 0 {
		return nil
	}
	err := documents.CreateService(s.repository.getDB()).UpsertMany(items.GetDocuments())
	if err != nil {
		return err
	}

	err = s.repository.upsertMany(items)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteMany(idPool []string) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
