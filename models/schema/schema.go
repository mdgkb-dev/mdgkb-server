package schema

type Schema struct {
	Human          map[string]string `json:"human"`
	Comment        map[string]string `json:"comment"`
	Doctors        map[string]string `json:"doctor"`
	MedicalProfile map[string]string `json:"medicalProfile"`
	Division       map[string]string `json:"division"`
	DoctorUser     map[string]string `json:"doctorUser"`
	Center         map[string]string `json:"center"`
	Teacher        map[string]string `json:"teacher"`
	DpoCourse      map[string]string `json:"dpoCourse"`
	DpoBaseCourse  map[string]string `json:"dpoBaseCourse"`
}

func CreateSchema() Schema {
	return Schema{
		Human:          createHumanSchema(),
		Comment:        createCommentsSchema(),
		Doctors:        createDoctorsSchema(),
		Division:       createDivisionSchema(),
		MedicalProfile: createMedicalProfileSchema(),
		DoctorUser:     createDoctorUserSchema(),
		Center:         createCenterSchema(),
		Teacher:        createTeacherSchema(),
		DpoCourse:      createDpoCourseSchema(),
		DpoBaseCourse:  createDpoBaseCourseSchema(),
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
		"rating":      "rating",
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
		"tableName":                    "divisions_view",
		"sortColumn":                   "name",
		"key":                          "division",
		"name":                         "name",
		"value":                        "id",
		"label":                        "name",
		"commentsCount":                "comments_count",
		"hospitalizationContactInfoId": "hospitalization_contact_info_id",
	}
}

func createMedicalProfileSchema() map[string]string {
	return map[string]string{
		"tableName":  "medical_profiles",
		"key":        "medicalProfile",
		"sortColumn": "name",
		"value":      "id",
		"label":      "name",
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

func createCenterSchema() map[string]string {
	return map[string]string{
		"tableName":  "centers",
		"sortColumn": "name",
		"key":        "center",
		"value":      "id",
		"label":      "name",
		"name":       "name",
	}
}

func createTeacherSchema() map[string]string {
	return map[string]string{
		"tableName": "teachers_view",
		"key":       "teacher",
		"id":        "id",
		"fullName":  "full_name",
	}
}

func createDpoCourseSchema() map[string]string {
	return map[string]string{
		"tableName": "dpo_courses",
		"key":       "dpoCourse",
		"id":        "id",
		"name":      "name",
		"hours":     "hours",
		"teacherId": "teacher_id",
		"listeners": "listeners",
		"start":     "dpo_course_start",
	}
}

func createDpoBaseCourseSchema() map[string]string {
	return map[string]string{
		"tableName": "dpo_base_courses",
		"key":       "dpoCourse",
		"id":        "id",
		"name":      "name",
		"hours":     "hours",
		"teacherId": "teacher_id",
		"listeners": "listeners",
		"start":     "dpo_course_start",
	}
}
