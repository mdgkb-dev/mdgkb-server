package schema

type Schema struct {
	HumanSchema          map[string]string `json:"human"`
	CommentsSchema       map[string]string `json:"comments"`
	DoctorsSchema        map[string]string `json:"doctor"`
	MedicalProfileSchema map[string]string `json:"medicalProfile"`
	DivisionSchema       map[string]string `json:"division"`
}

func CreateSchema() Schema {
	return Schema{
		HumanSchema:          createHumanSchema(),
		CommentsSchema:       createCommentsSchema(),
		DoctorsSchema:        createDoctorsSchema(),
		DivisionSchema:       createDivisionSchema(),
		MedicalProfileSchema: createMedicalProfileSchema(),
	}
}

func createHumanSchema() map[string]string {
	return map[string]string{
		"tableName": "human",
		"dateBirth": "date_birth",
		"fullName":  "full_name",
		"isMale":    "is_male",
	}
}

func createCommentsSchema() map[string]string {
	return map[string]string{
		"tableName":   "comment",
		"publishedOn": "publishedOn",
		"positive":    "positive",
	}
}

func createDoctorsSchema() map[string]string {
	return map[string]string{
		"tableName":        "doctors_view",
		"fullName":         "full_name",
		"divisionId":       "division_id",
		"medicalProfileId": "medical_profile_id",
		"mosDoctorLink":    "mos_doctor_link",
		"onlineDoctorId":   "online_doctor_id",
		"commentsCount":    "comments_count",
	}
}

func createDivisionSchema() map[string]string {
	return map[string]string{
		"tableName": "divisions",
		"value":     "id",
		"label":     "name",
	}
}

func createMedicalProfileSchema() map[string]string {
	return map[string]string{
		"tableName": "medical_profiles",
		"value":     "id",
		"label":     "name",
	}
}
