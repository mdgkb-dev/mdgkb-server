package models

type EducationalOrganization struct {
	EducationalOrganizationManagers          EducationalManagers `json:"educationalOrganizationManagers"`
	EducationalOrganizationManagersForDelete []string            `json:"educationalOrganizationManagersForDelete"`

	EducationalOrganizationTeachers Teachers `json:"teachers"`
	TeachersForDelete               []string `json:"teachersForDelete"`

	EducationalOrganizationAcademics          EducationalAcademics `json:"educationalOrganizationAcademics"`
	EducationalOrganizationAcademicsForDelete []string             `json:"educationalOrganizationAcademicsForDelete"`
}

func (item *EducationalOrganization) SetFilePath(fileID *string) *string {
	return nil
}
