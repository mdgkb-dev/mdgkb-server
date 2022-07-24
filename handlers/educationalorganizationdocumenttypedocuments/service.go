package educationalorganizationdocumenttypedocuments

import (
	document "mdgkb/mdgkb-server/handlers/documents"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(items models.EducationalOrganizationDocumentTypeDocuments) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpsertMany(items models.EducationalOrganizationDocumentTypeDocuments) error {
	if len(items) == 0 {
		return nil
	}
	documentsService := document.CreateService(s.helper)
	err := documentsService.UpsertMany(items.GetDocuments())
	if err != nil {
		return err
	}
	items.SetForeignKeys()
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
