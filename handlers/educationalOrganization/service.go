package educationalOrganization

import (
	educationalOrganizationPages "mdgkb/mdgkb-server/handlers/educationalOrganizationDocumentPages"
	"mdgkb/mdgkb-server/handlers/educationalOrganizationDocumentTypes"
	"mdgkb/mdgkb-server/handlers/educationalOrganizationManagers"
	"mdgkb/mdgkb-server/handlers/educationalOrganizationProperties"
	"mdgkb/mdgkb-server/handlers/educationalOrganizationTeachers"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Get() (*models.EducationalOrganization,  error) {
	item := models.EducationalOrganization{}
	var err error
	propertiesService := educationalOrganizationProperties.CreateService(s.repository.getDB())
	item.EducationalOrganizationProperties, err = propertiesService.GetAll()
	if err != nil {
		return nil, err
	}
	managersService := educationalOrganizationManagers.CreateService(s.repository.getDB())
	item.EducationalOrganizationManagers, err = managersService.GetAll()
	if err != nil {
		return nil, err
	}
	teachersService := educationalOrganizationTeachers.CreateService(s.repository.getDB())
	item.EducationalOrganizationTeachers, err = teachersService.GetAll()
	if err != nil {
		return nil, err
	}
	educationalOrganizationDocumentTypesService := educationalOrganizationDocumentTypes.CreateService(s.repository.getDB())
	item.EducationalOrganizationDocumentTypes, err = educationalOrganizationDocumentTypesService.GetAll()
	if err != nil {
		return nil, err
	}
	item.EducationalOrganizationPages, err = educationalOrganizationPages.CreateService(s.repository.getDB()).GetAll()
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *Service) Update(item *models.EducationalOrganization) error {
	propertiesService := educationalOrganizationProperties.CreateService(s.repository.getDB())
	err := propertiesService.DeleteMany(item.EducationalOrganizationPropertiesForDelete)
	if err != nil {
		return  err
	}
	err = propertiesService.UpsertMany(item.EducationalOrganizationProperties)
	if err != nil {
		return  err
	}

	managersService := educationalOrganizationManagers.CreateService(s.repository.getDB())
	err = managersService.DeleteMany(item.EducationalOrganizationManagersForDelete)
	if err != nil {
		return  err
	}
	err = managersService.UpsertMany(item.EducationalOrganizationManagers)
	if err != nil {
		return  err
	}

	teachersService := educationalOrganizationTeachers.CreateService(s.repository.getDB())
	err = teachersService.DeleteMany(item.EducationalOrganizationTeachersForDelete)
	if err != nil {
		return  err
	}
	err = teachersService.UpsertMany(item.EducationalOrganizationTeachers)
	if err != nil {
		return  err
	}
	educationalOrganizationDocumentTypesService := educationalOrganizationDocumentTypes.CreateService(s.repository.getDB())
	err = educationalOrganizationDocumentTypesService.DeleteMany(item.EducationalOrganizationDocumentTypesForDelete)
	if err != nil {
		return  err
	}
	err = educationalOrganizationDocumentTypesService.UpsertMany(item.EducationalOrganizationDocumentTypes)
	if err != nil {
		return  err
	}
	return nil
}
