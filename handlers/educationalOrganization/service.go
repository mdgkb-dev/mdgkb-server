package educationalOrganization

import (
	"fmt"
	"mdgkb/mdgkb-server/handlers/educationalOrganizationAcademics"
	"mdgkb/mdgkb-server/handlers/educationalOrganizationDocumentTypes"
	"mdgkb/mdgkb-server/handlers/educationalOrganizationManagers"
	"mdgkb/mdgkb-server/handlers/educationalOrganizationProperties"
	"mdgkb/mdgkb-server/handlers/teachers"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Get() (*models.EducationalOrganization, error) {
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
	teachersService := teachers.CreateService(s.repository.getDB(), s.helper)
	item.EducationalOrganizationTeachers, err = teachersService.GetAll()
	if err != nil {
		return nil, err
	}
	academicsService := educationalOrganizationAcademics.CreateService(s.repository.getDB())
	item.EducationalOrganizationAcademics, err = academicsService.GetAll()
	if err != nil {
		return nil, err
	}
	//educationalOrganizationDocumentTypesService := educationalOrganizationDocumentTypes.CreateService(s.repository.getDB())
	//item.EducationalOrganizationDocumentTypes, err = educationalOrganizationDocumentTypesService.GetAll()
	if err != nil {
		return nil, err
	}
	//item.EducationalOrganizationPages, err = educationalOrganizationPages.CreateService(s.repository.getDB()).GetAll()
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *Service) Update(item *models.EducationalOrganization) error {
	propertiesService := educationalOrganizationProperties.CreateService(s.repository.getDB())
	err := propertiesService.DeleteMany(item.EducationalOrganizationPropertiesForDelete)
	if err != nil {
		return err
	}
	err = propertiesService.UpsertMany(item.EducationalOrganizationProperties)
	if err != nil {
		return err
	}

	managersService := educationalOrganizationManagers.CreateService(s.repository.getDB())
	err = managersService.DeleteMany(item.EducationalOrganizationManagersForDelete)
	if err != nil {
		return err
	}
	err = managersService.UpsertMany(item.EducationalOrganizationManagers)
	if err != nil {
		return err
	}

	teachersService := teachers.CreateService(s.repository.getDB(), s.helper)
	fmt.Println(item.TeachersForDelete)
	fmt.Println(item.TeachersForDelete)
	fmt.Println(item.TeachersForDelete)
	fmt.Println(item.TeachersForDelete)
	fmt.Println(item.TeachersForDelete)
	fmt.Println(item.TeachersForDelete)
	err = teachersService.DeleteMany(item.TeachersForDelete)
	if err != nil {
		return err
	}
	err = teachersService.UpsertMany(item.EducationalOrganizationTeachers)
	if err != nil {
		return err
	}

	educationalOrganizationAcademicsService := educationalOrganizationAcademics.CreateService(s.repository.getDB())
	err = educationalOrganizationAcademicsService.DeleteMany(item.EducationalOrganizationAcademicsForDelete)
	if err != nil {
		return err
	}
	err = educationalOrganizationAcademicsService.UpsertMany(item.EducationalOrganizationAcademics)
	if err != nil {
		return err
	}

	educationalOrganizationDocumentTypesService := educationalOrganizationDocumentTypes.CreateService(s.repository.getDB())
	err = educationalOrganizationDocumentTypesService.DeleteMany(item.EducationalOrganizationDocumentTypesForDelete)
	if err != nil {
		return err
	}

	err = educationalOrganizationDocumentTypesService.UpsertMany(item.EducationalOrganizationDocumentTypes)
	if err != nil {
		return err
	}
	return nil
}
