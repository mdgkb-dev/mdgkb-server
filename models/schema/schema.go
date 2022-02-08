package schema

type Schema struct {
	Human          map[string]string `json:"human"`
	Comment        map[string]string `json:"comment"`
	Doctors        map[string]string `json:"doctor"`
	MedicalProfile map[string]string `json:"medicalProfile"`
	Division       map[string]string `json:"division"`
	DoctorUser     map[string]string `json:"doctorUser"`
}

func CreateSchema() Schema {
	return Schema{
		Human:          createHumanSchema(),
		Comment:        createCommentsSchema(),
		Doctors:        createDoctorsSchema(),
		Division:       createDivisionSchema(),
		MedicalProfile: createMedicalProfileSchema(),
		DoctorUser:     createDoctorUserSchema(),
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
		"publishedOn": "published_on",
		"positive":    "positive",
		"userId":      "user_id",
	}
}

func createDoctorsSchema() map[string]string {
	return map[string]string{
		"tableName":          "doctors_view",
		"key":                "doctor",
		"id":                 "id",
		"favouriteTableName": "doctors_users",
		"fullName":           "full_name",
		"divisionId":         "division_id",
		"medicalProfileId":   "medical_profile_id",
		"mosDoctorLink":      "mos_doctor_link",
		"onlineDoctorId":     "online_doctor_id",
		"commentsCount":      "comments_count",
	}
}

func createDivisionSchema() map[string]string {
	return map[string]string{
		"tableName": "divisions",
		"key":       "division",
		"value":     "id",
		"label":     "name",
	}
}

func createMedicalProfileSchema() map[string]string {
	return map[string]string{
		"tableName": "medical_profiles",
		"key":       "medicalProfile",
		"value":     "id",
		"label":     "name",
	}
}

func createDoctorUserSchema() map[string]string {
	return map[string]string{
		"tableName": "doctors_users",
		"id":        "id",
		"userId":    "user_id",
		"doctorId":  "doctor_id",
	}
}
