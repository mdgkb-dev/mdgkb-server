package publicdocumenttypes

import (
	"mdgkb/mdgkb-server/handlers/documenttypes"
	"mdgkb/mdgkb-server/handlers/educationpublicdocumenttypes"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetAll() (models.PublicDocumentTypes, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.PublicDocumentType, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.PublicDocumentType) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	educationPublicDocumentTypesService := educationpublicdocumenttypes.CreateService(s.helper)
	err = educationPublicDocumentTypesService.Upsert(item.EducationPublicDocumentType)
	if err != nil {
		return err
	}
	if item.EducationPublicDocumentType == nil {
		err = educationPublicDocumentTypesService.DeleteByPublicDocumentTypeID(item.ID)
		if err != nil {
			return err
		}
	}
	err = documenttypes.CreateService(s.helper).UpsertMany(item.DocumentTypes)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Update(item *models.PublicDocumentType) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	educationPublicDocumentTypesService := educationpublicdocumenttypes.CreateService(s.helper)
	err = educationPublicDocumentTypesService.Upsert(item.EducationPublicDocumentType)
	if err != nil {
		return err
	}
	if item.EducationPublicDocumentType == nil {
		err = educationPublicDocumentTypesService.DeleteByPublicDocumentTypeID(item.ID)
		if err != nil {
			return err
		}
	}
	documentTypeService := documenttypes.CreateService(s.helper)
	err = documentTypeService.DeleteMany(item.DocumentTypesForDelete)
	if err != nil {
		return err
	}
	err = documentTypeService.UpsertMany(item.DocumentTypes)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}

func (s *Service) UpdateOrder(items models.PublicDocumentTypes) error {
	return s.repository.upsertMany(items)
}
