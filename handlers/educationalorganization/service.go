package educationalorganization

import (
	"mdgkb/mdgkb-server/handlers/educationalorganizationacademics"
	"mdgkb/mdgkb-server/handlers/educationalorganizationmanagers"
	"mdgkb/mdgkb-server/handlers/teachers"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Get() (*models.EducationalOrganization, error) {
	item := models.EducationalOrganization{}
	var err error
	managersService := educationalorganizationmanagers.CreateService(s.helper)
	item.EducationalOrganizationManagers, err = managersService.GetAll()
	if err != nil {
		return nil, err
	}
	teachersService := teachers.CreateService(s.helper)
	item.EducationalOrganizationTeachers, err = teachersService.GetAll()
	if err != nil {
		return nil, err
	}
	academicsService := educationalorganizationacademics.CreateService(s.helper)
	item.EducationalOrganizationAcademics, err = academicsService.GetAll()
	if err != nil {
		return nil, err
	}
	//educationalOrganizationDocumentTypesService := educationalOrganizationDocumentTypes.CreateService(s.helper)
	//item.EducationalOrganizationDocumentTypes, err = educationalOrganizationDocumentTypesService.GetAll()
	if err != nil {
		return nil, err
	}
	//item.EducationalOrganizationPages, err = educationalOrganizationPages.CreateService(s.helper).GetAll()
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *Service) Update(item *models.EducationalOrganization) error {
	managersService := educationalorganizationmanagers.CreateService(s.helper)
	err := managersService.DeleteMany(item.EducationalOrganizationManagersForDelete)
	if err != nil {
		return err
	}
	err = managersService.UpsertMany(item.EducationalOrganizationManagers)
	if err != nil {
		return err
	}

	teachersService := teachers.CreateService(s.helper)
	err = teachersService.DeleteMany(item.TeachersForDelete)
	if err != nil {
		return err
	}
	err = teachersService.UpsertMany(item.EducationalOrganizationTeachers)
	if err != nil {
		return err
	}

	educationalOrganizationAcademicsService := educationalorganizationacademics.CreateService(s.helper)
	err = educationalOrganizationAcademicsService.DeleteMany(item.EducationalOrganizationAcademicsForDelete)
	if err != nil {
		return err
	}
	err = educationalOrganizationAcademicsService.UpsertMany(item.EducationalOrganizationAcademics)
	if err != nil {
		return err
	}
	return nil
}
