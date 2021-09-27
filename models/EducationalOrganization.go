package models

type EducationalOrganization struct {
	EducationalOrganizationProperties EducationalOrganizationProperties `json:"educationalOrganizationProperties"`
	EducationalOrganizationPropertiesForDelete []string `json:"educationalOrganizationPropertiesForDelete"`

	EducationalOrganizationManagers EducationalOrganizationManagers `json:"educationalOrganizationManagers"`
	EducationalOrganizationManagersForDelete []string `json:"educationalOrganizationManagersForDelete"`

	EducationalOrganizationTeachers EducationalOrganizationTeachers `json:"educationalOrganizationTeachers"`
	EducationalOrganizationTeachersForDelete []string `json:"educationalOrganizationTeachersForDelete"`

	EducationalOrganizationDocumentTypes EducationalOrganizationDocumentTypes `json:"educationalOrganizationDocumentTypes"`
	EducationalOrganizationDocumentTypesForDelete []string `json:"educationalOrganizationDocumentTypesForDelete"`
}

func (item *EducationalOrganization) SetFilePath(fileId *string) *string {
	path := item.EducationalOrganizationDocumentTypes.SetFilePath(fileId)
	if path != nil {
		return path
	}
	return nil
}