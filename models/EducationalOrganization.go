package models

type EducationalOrganization struct {
	EducationalOrganizationProperties EducationalOrganizationProperties `json:"educationalOrganizationProperties"`
	EducationalOrganizationPropertiesForDelete []string `json:"educationalOrganizationPropertiesForDelete"`

	EducationalOrganizationManagers EducationalOrganizationManagers `json:"educationalOrganizationManagers"`
	EducationalOrganizationManagersForDelete []string `json:"educationalOrganizationManagersForDelete"`

	EducationalOrganizationTeachers EducationalOrganizationTeachers `json:"educationalOrganizationTeachers"`
	EducationalOrganizationTeachersForDelete []string `json:"educationalOrganizationTeachersForDelete"`
}

