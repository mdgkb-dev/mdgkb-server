package models

type EducationalOrganization struct {
	EducationalOrganizationProperties          EducationalOrganizationProperties `json:"educationalOrganizationProperties"`
	EducationalOrganizationPropertiesForDelete []string                          `json:"educationalOrganizationPropertiesForDelete"`

	EducationalOrganizationManagers          EducationalOrganizationManagers `json:"educationalOrganizationManagers"`
	EducationalOrganizationManagersForDelete []string                        `json:"educationalOrganizationManagersForDelete"`

	EducationalOrganizationTeachers          Teachers `json:"teachers"`
	EducationalOrganizationTeachersForDelete []string `json:"educationalOrganizationTeachersForDelete"`

	EducationalOrganizationAcademics          EducationalOrganizationAcademics `json:"educationalOrganizationAcademics"`
	EducationalOrganizationAcademicsForDelete []string                         `json:"educationalOrganizationAcademicsForDelete"`

	EducationalOrganizationDocumentTypes          EducationalOrganizationDocumentTypes `json:"educationalOrganizationDocumentTypes"`
	EducationalOrganizationDocumentTypesForDelete []string                             `json:"educationalOrganizationDocumentTypesForDelete"`

	//EducationalOrganizationPages EducationalOrganizationPages `json:"educationalOrganizationPages"`
	//EducationalOrganizationPagesForDelete []string `json:"educationalOrganizationPagesForDelete"`
}

func (item *EducationalOrganization) SetFilePath(fileID *string) *string {
	path := item.EducationalOrganizationDocumentTypes.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}
