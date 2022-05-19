package schema

type Schema struct {
	Human                            map[string]string `json:"human"`
	Comment                          map[string]string `json:"comment"`
	Doctors                          map[string]string `json:"doctor"`
	MedicalProfile                   map[string]string `json:"medicalProfile"`
	Division                         map[string]string `json:"division"`
	DoctorUser                       map[string]string `json:"doctorUser"`
	Center                           map[string]string `json:"center"`
	Teacher                          map[string]string `json:"teacher"`
	DpoCourse                        map[string]string `json:"dpoCourse"`
	DpoBaseCourse                    map[string]string `json:"dpoBaseCourse"`
	EducationalManager               map[string]string `json:"educationalManager"`
	Specialization                   map[string]string `json:"specialization"`
	Vacancy                          map[string]string `json:"vacancy"`
	DpoCourseSpecialization          map[string]string `json:"dpoCourseSpecialization"`
	DpoApplication                   map[string]string `json:"dpoApplication"`
	PostgraduateCourse               map[string]string `json:"postgraduateCourse"`
	ResidencyCourse                  map[string]string `json:"residencyCourse"`
	EducationPublicDocumentType      map[string]string `json:"educationPublicDocumentType"`
	PublicDocumentType               map[string]string `json:"publicDocumentType"`
	EducationYear                    map[string]string `json:"educationYear"`
	PostgraduateCourseSpecialization map[string]string `json:"postgraduateCourseSpecialization"`
	EducationalOrganizationAcademic  map[string]string `json:"educationalOrganizationAcademic"`
	Role                             map[string]string `json:"role"`
	News                             map[string]string `json:"news"`
	PathPermission                   map[string]string `json:"pathPermission"`
	NewsToTag                        map[string]string `json:"newsToTag"`
}

func CreateSchema() Schema {
	return Schema{
		Human:                            createHumanSchema(),
		Comment:                          createCommentsSchema(),
		Doctors:                          createDoctorsSchema(),
		Division:                         createDivisionSchema(),
		MedicalProfile:                   createMedicalProfileSchema(),
		DoctorUser:                       createDoctorUserSchema(),
		Center:                           createCenterSchema(),
		Teacher:                          createTeacherSchema(),
		DpoCourse:                        createDpoCourseSchema(),
		DpoBaseCourse:                    createDpoBaseCourseSchema(),
		EducationalManager:               createEducationalManagerSchema(),
		Specialization:                   createSpecializationSchema(),
		Vacancy:                          createVacancySchema(),
		DpoCourseSpecialization:          createDpoCourseSpecializationSchema(),
		DpoApplication:                   createDpoApplicationsSchema(),
		PostgraduateCourse:               createPostgraduateCourseSchema(),
		ResidencyCourse:                  createResidencyCourseSchema(),
		EducationPublicDocumentType:      createEducationPublicDocumentTypeSchema(),
		PublicDocumentType:               createPublicDocumentTypeSchema(),
		EducationYear:                    createEducationYearSchema(),
		PostgraduateCourseSpecialization: createPostgraduateCourseSpecializationSchema(),
		EducationalOrganizationAcademic:  createEducationalOrganizationAcademicsSchema(),
		Role:                             createRolesSchema(),
		PathPermission:                   createPathPermissionsSchema(),
		News:                             createNewsSchema(),
		NewsToTag:                        createNewsToTagSchema(),
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
		"key":         "text",
		"modChecked":  "mod_checked",
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
		"tableName":        "dpo_courses_view",
		"key":              "dpoCourse",
		"id":               "id",
		"name":             "name",
		"slug":             "slug",
		"cost":             "cost",
		"isNmo":            "is_nmo",
		"specializationId": "specialization_id",
		"hours":            "hours",
		"value":            "id",
		"label":            "name",
		"sortColumn":       "name",
		"teacherId":        "teacher_id",
		"listeners":        "listeners",
		"start":            "dpo_course_start",
		"minStart":         "min_dpo_course_start",
		"minEnd":           "min_dpo_course_end",
	}
}

func createDpoApplicationsSchema() map[string]string {
	return map[string]string{
		"tableName": "dpo_applications_view",
		"key":       "dpoApplication",
		"id":        "id",
		"createdAt": "created_at",
		"isNmo":     "is_nmo",
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

func createEducationalManagerSchema() map[string]string {
	return map[string]string{
		"tableName": "educational_managers_view",
		"key":       "educationalManager",
		"order":     "educational_managers_view",
		"id":        "id",
		"fullName":  "fullName",
	}
}

func createDpoCourseSpecializationSchema() map[string]string {
	return map[string]string{
		"tableName":        "dpo_courses_specializations",
		"key":              "dpoCourseSpecialization",
		"id":               "id",
		"dpoCourseId":      "dpo_course_id",
		"specializationId": "specialization_id",
	}
}

func createSpecializationSchema() map[string]string {
	return map[string]string{
		"tableName":  "specializations",
		"key":        "specialization",
		"order":      "name",
		"id":         "id",
		"name":       "name",
		"value":      "id",
		"sortColumn": "name",
		"label":      "name",
	}
}

func createVacancySchema() map[string]string {
	return map[string]string{
		"tableName":         "vacancies_view",
		"key":               "vacancy",
		"title":             "title",
		"minSalary":         "min_salary",
		"divisionId":        "division_id",
		"slug":              "slug",
		"date":              "vacancy_date",
		"value":             "slug",
		"maxSalary":         "max_salary",
		"sortColumn":        "title",
		"responsesCount":    "responses_count",
		"newResponsesCount": "new_responses_count",
		"active":            "active",
	}
}

func createPostgraduateCourseSchema() map[string]string {
	return map[string]string{
		"tableName":     "postgraduate_courses_view",
		"value":         "id",
		"key":           "postgraduateCourse",
		"name":          "name",
		"cost":          "cost",
		"code":          "code",
		"years":         "years",
		"educationForm": "education_form",
	}
}

func createPostgraduateCourseSpecializationSchema() map[string]string {
	return map[string]string{
		"tableName":            "postgraduate_courses_specializations",
		"key":                  "postgraduateCourseSpecialization",
		"id":                   "id",
		"postgraduateCourseId": "postgraduate_course_id",
		"specializationId":     "specialization_id",
	}
}

func createResidencyCourseSchema() map[string]string {
	return map[string]string{
		"tableName":     "residency_courses_view",
		"value":         "id",
		"key":           "residencyCourse",
		"name":          "name",
		"code":          "code",
		"slug":          "slug",
		"cost":          "cost",
		"freePlaces":    "free_places",
		"paidPlaces":    "paid_places",
		"startYear":     "start_year",
		"startYearId":   "start_year_id",
		"endYear":       "end_year",
		"endYearId":     "end_year_id",
		"educationForm": "education_form",
	}
}

func createEducationPublicDocumentTypeSchema() map[string]string {
	return map[string]string{
		"tableName":            "education_public_document_types",
		"value":                "id",
		"id":                   "id",
		"key":                  "educationPublicDocumentType",
		"publicDocumentTypeId": "public_document_type_id",
	}
}

func createPublicDocumentTypeSchema() map[string]string {
	return map[string]string{
		"tableName": "public_document_types",
		"value":     "id",
		"id":        "id",
		"key":       "publicDocumentTypeSchema",
	}
}

func createEducationYearSchema() map[string]string {
	return map[string]string{
		"tableName":  "education_years",
		"value":      "id",
		"sortColumn": "year",
		"label":      "year",
		"id":         "id",
		"key":        "educationYear",
	}
}

func createEducationalOrganizationAcademicsSchema() map[string]string {
	return map[string]string{
		"tableName":  "educational_organization_academics_view",
		"key":        "educationalOrganizationAcademic",
		"id":         "id",
		"fullName":   "full_name",
		"value":      "id",
		"sortColumn": "full_name",
	}
}

func createRolesSchema() map[string]string {
	return map[string]string{
		"tableName":  "roles",
		"key":        "role",
		"name":       "name",
		"id":         "id",
		"value":      "id",
		"label":      "label",
		"sortColumn": "label",
	}
}

func createPathPermissionsSchema() map[string]string {
	return map[string]string{
		"tableName":  "path_permissions",
		"key":        "pathPermission",
		"id":         "id",
		"value":      "id",
		"resource":   "resource",
		"sortColumn": "resource",
	}
}

func createNewsSchema() map[string]string {
	return map[string]string{
		"tableName":   "news_view",
		"key":         "news",
		"id":          "id",
		"value":       "id",
		"sortColumn":  "title",
		"title":       "title",
		"status":      "status",
		"previewText": "preview_text",
		"slug":        "slug",
		"publishedOn": "published_on",
		"createdAt":   "created_at",
		"description": "description",
		"main":        "main",
		"subMain":     "sub_main",
		"articleLink": "article_link",
		"isArticle":   "is_article",
		"viewsCount":  "views_count",
		"isDraft":     "is_draft",
	}
}

func createNewsToTagSchema() map[string]string {
	return map[string]string{
		"tableName": "news_to_tags",
		"key":       "newsToTag",
		"id":        "id",
		"value":     "id",
		"newsId":    "news_id",
		"tagId":     "tag_id",
	}
}
