package models

type EducationalOrganization struct {
	EducationalOrganizationProperties          EducationalOrganizationProperties `json:"educationalOrganizationProperties"`
	EducationalOrganizationPropertiesForDelete []string                          `json:"educationalOrganizationPropertiesForDelete"`

	EducationalOrganizationManagers          EducationalManagers `json:"educationalOrganizationManagers"`
	EducationalOrganizationManagersForDelete []string            `json:"educationalOrganizationManagersForDelete"`

	EducationalOrganizationTeachers Teachers `json:"teachers"`
	TeachersForDelete               []string `json:"teachersForDelete"`

	EducationalOrganizationAcademics          EducationalOrganizationAcademics `json:"educationalOrganizationAcademics"`
	EducationalOrganizationAcademicsForDelete []string                         `json:"educationalOrganizationAcademicsForDelete"`
}

func (item *EducationalOrganization) SetFilePath(fileID *string) *string {
	return nil
}
