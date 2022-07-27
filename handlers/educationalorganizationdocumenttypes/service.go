package educationalorganizationdocumenttypes

import (
	"mdgkb/mdgkb-server/handlers/educationalorganizationdocumenttypedocuments"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) CreateMany(items models.EducationalOrganizationDocumentTypes) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.createMany(items)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAll() (models.EducationalOrganizationDocumentTypes, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) UpsertMany(items models.EducationalOrganizationDocumentTypes) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetChildrenForeignKeys()
	educationalOrganizationDocumentTypeDocumentsService := educationalorganizationdocumenttypedocuments.CreateService(s.helper)
	err = educationalOrganizationDocumentTypeDocumentsService.DeleteMany(items.GetIDForDelete())
	if err != nil {
		return err
	}
	err = educationalOrganizationDocumentTypeDocumentsService.UpsertMany(items.GetEducationalOrganizationDocumentTypeDocuments())
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
