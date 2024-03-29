package schema

//
//type Schema struct {
//	Human                            map[string]string `json:"human"`
//	Comment                          map[string]string `json:"comment"`
//	Doctors                          map[string]string `json:"doctor"`
//	MedicalProfile                   map[string]string `json:"medicalProfile"`
//	Division                         map[string]string `json:"division"`
//	DoctorUser                       map[string]string `json:"doctorUser"`
//	User                             map[string]string `json:"user"`
//	Center                           map[string]string `json:"center"`
//	Teacher                          map[string]string `json:"teacher"`
//	NmoCourse                        map[string]string `json:"nmoCourse"`
//	DpoBaseCourse                    map[string]string `json:"dpoBaseCourse"`
//	EducationalManager               map[string]string `json:"educationalManager"`
//	Specialization                   map[string]string `json:"specialization"`
//	Vacancy                          map[string]string `json:"vacancy"`
//	VacancyResponse                  map[string]string `json:"vacancyResponse"`
//	NmoCourseSpecialization          map[string]string `json:"nmoCourseSpecialization"`
//	VisitsApplication                map[string]string `json:"visitsApplication"`
//	DpoApplication                   map[string]string `json:"dpoApplication"`
//	ResidencyApplication             map[string]string `json:"residencyApplication"`
//	PostgraduateApplication          map[string]string `json:"postgraduateApplication"`
//	PostgraduateCourse               map[string]string `json:"postgraduateCourse"`
//	ResidencyCourse                  map[string]string `json:"residencyCourse"`
//	EducationPublicDocumentType      map[string]string `json:"educationPublicDocumentType"`
//	PublicDocumentType               map[string]string `json:"publicDocumentType"`
//	EducationYear                    map[string]string `json:"educationYear"`
//	PostgraduateCourseSpecialization map[string]string `json:"postgraduateCourseSpecialization"`
//	EducationalOrganizationAcademic  map[string]string `json:"educationalOrganizationAcademic"`
//	Role                             map[string]string `json:"role"`
//	News                             map[string]string `json:"news"`
//	PathPermission                   map[string]string `json:"pathPermission"`
//	NewsToTag                        map[string]string `json:"newsToTag"`
//	FormStatus                       map[string]string `json:"formStatus"`
//	FormPattern                      map[string]string `json:"formPattern"`
//	Question                         map[string]string `json:"question"`
//	TreatDirection                   map[string]string `json:"treatDirection"`
//	DoctorComment                    map[string]string `json:"doctorComment"`
//	DivisionComment                  map[string]string `json:"divisionComment"`
//	NewsComment                      map[string]string `json:"newsComment"`
//	Building                         map[string]string `json:"building"`
//	Diet                             map[string]string `json:"diet"`
//	AgePeriod                        map[string]string `json:"agePeriod"`
//	DoctorDivision                   map[string]string `json:"doctorDivision"`
//	Employee                         map[string]string `json:"employee"`
//	Hospitalization                  map[string]string `json:"hospitalization"`
//	SupportMessage                   map[string]string `json:"supportMessage"`
//	DailyMenu                        map[string]string `json:"dailyMenu"`
//	DishesGroup                      map[string]string `json:"dishesGroup"`
//	DailyMenuOrder                   map[string]string `json:"dailyMenuOrder"`
//}
//
//func CreateSchema() Schema {
//	return Schema{
//		Human:                            createHumanSchema(),
//		Comment:                          createCommentsSchema(),
//		Doctors:                          createDoctorsSchema(),
//		Division:                         createDivisionSchema(),
//		MedicalProfile:                   createMedicalProfileSchema(),
//		DoctorUser:                       createDoctorUserSchema(),
//		User:                             createUserSchema(),
//		Center:                           createCenterSchema(),
//		Teacher:                          createTeacherSchema(),
//		NmoCourse:                        createNmoCourseSchema(),
//		DpoBaseCourse:                    createDpoBaseCourseSchema(),
//		EducationalManager:               createEducationalManagerSchema(),
//		Specialization:                   createSpecializationSchema(),
//		Vacancy:                          createVacancySchema(),
//		VacancyResponse:                  createVacancyResponseSchema(),
//		NmoCourseSpecialization:          createNmoCourseSpecializationSchema(),
//		VisitsApplication:                createVisitsApplicationsSchema(),
//		DpoApplication:                   createDpoApplicationsSchema(),
//		PostgraduateApplication:          createPostgraduateApplicationsSchema(),
//		ResidencyApplication:             createResidencyApplicationsSchema(),
//		PostgraduateCourse:               createPostgraduateCourseSchema(),
//		ResidencyCourse:                  createResidencyCourseSchema(),
//		EducationPublicDocumentType:      createEducationPublicDocumentTypeSchema(),
//		PublicDocumentType:               createPublicDocumentTypeSchema(),
//		EducationYear:                    createEducationYearSchema(),
//		PostgraduateCourseSpecialization: createPostgraduateCourseSpecializationSchema(),
//		EducationalOrganizationAcademic:  createEducationalOrganizationAcademicsSchema(),
//		Role:                             createRolesSchema(),
//		PathPermission:                   createPathPermissionsSchema(),
//		News:                             createNewsSchema(),
//		NewsToTag:                        createNewsToTagSchema(),
//		FormStatus:                       createFormStatusSchema(),
//		Question:                         createQuestionSchema(),
//		TreatDirection:                   createTreatDirectionSchema(),
//		DoctorComment:                    createDoctorCommentSchema(),
//		DivisionComment:                  createDivisionCommentSchema(),
//		NewsComment:                      createNewsCommentSchema(),
//		Building:                         createBuildingSchema(),
//		Diet:                             createDietSchema(),
//		AgePeriod:                        createAgePeriodSchema(),
//		DoctorDivision:                   createDoctorDivisionSchema(),
//		Hospitalization:                  createHospitalizationSchema(),
//		SupportMessage:                   createSupportMessageSchema(),
//		DailyMenu:                        createDailyMenuSchema(),
//		FormPattern:                      createFormPatternSchema(),
//		DishesGroup:                      createDishesGroupSchema(),
//		DailyMenuOrder:                   createDailyMenuOrderSchema(),
//	}
//}
//
//func createHumanSchema() map[string]string {
//	return map[string]string{
//		"tableName": "human",
//		"dateBirth": "date_birth",
//		"fullName":  "full_name",
//		"isMale":    "is_male",
//	}
//}
//
//func createCommentsSchema() map[string]string {
//	return map[string]string{
//		"tableName":   "comments",
//		"publishedOn": "published_on",
//		"positive":    "positive",
//		"key":         "text",
//		"modChecked":  "mod_checked",
//		"rating":      "rating",
//		"userId":      "user_id",
//	}
//}
//
//func createDoctorsSchema() map[string]string {
//	return map[string]string{
//		"tableName":          "doctors_view",
//		"key":                "doctor",
//		"id":                 "id",
//		"favouriteTableName": "doctors_users",
//		"fullName":           "full_name",
//		"divisionId":         "division_id",
//		"medicalProfileId":   "medical_profile_id",
//		"mosDoctorLink":      "mos_doctor_link",
//		"onlineDoctorId":     "online_doctor_id",
//		"commentsCount":      "comments_count",
//		"isMale":             "is_male",
//		"divisionName":       "division_name",
//		"dateBirth":          "date_birth",
//	}
//}
//
//func createDivisionSchema() map[string]string {
//	return map[string]string{
//		"tableName":                    "divisions_view",
//		"id":                           "id",
//		"sortColumn":                   "name",
//		"key":                          "division",
//		"slug":                         "slug",
//		"name":                         "name",
//		"value":                        "id",
//		"label":                        "name",
//		"isCenter":                     "is_center",
//		"commentsCount":                "comments_count",
//		"hospitalizationContactInfoId": "hospitalization_contact_info_id",
//		"treatDirectionId":             "treat_direction_id",
//		"buildingId":                   "building_id",
//		"hasAmbulatory":                "has_ambulatory",
//		"hasDiagnostic":                "has_diagnostic",
//	}
//}
//
//func createMedicalProfileSchema() map[string]string {
//	return map[string]string{
//		"tableName":  "medical_profiles",
//		"key":        "medicalProfile",
//		"sortColumn": "name",
//		"value":      "id",
//		"label":      "name",
//	}
//}
//
//func createDoctorUserSchema() map[string]string {
//	return map[string]string{
//		"tableName": "doctors_users",
//		"id":        "id",
//		"userId":    "user_id",
//		"doctorId":  "doctor_id",
//	}
//}
//
//func createCenterSchema() map[string]string {
//	return map[string]string{
//		"tableName":  "centers",
//		"sortColumn": "name",
//		"key":        "center",
//		"value":      "id",
//		"label":      "name",
//		"name":       "name",
//	}
//}
//
//func createTeacherSchema() map[string]string {
//	return map[string]string{
//		"tableName": "teachers_view",
//		"key":       "teacher",
//		"id":        "id",
//		"fullName":  "full_name",
//		"dateBirth": "date_birth",
//	}
//}
//
//func createNmoCourseSchema() map[string]string {
//	return map[string]string{
//		"tableName":        "nmo_courses_view",
//		"key":              "nmoCourse",
//		"id":               "id",
//		"name":             "name",
//		"slug":             "slug",
//		"cost":             "cost",
//		"isNmo":            "is_nmo",
//		"specializationId": "specialization_id",
//		"hours":            "hours",
//		"value":            "id",
//		"label":            "name",
//		"sortColumn":       "name",
//		"teacherId":        "teacher_id",
//		"listeners":        "listeners",
//		"start":            "dpo_course_start",
//		"minStart":         "min_dpo_course_start",
//		"minEnd":           "min_dpo_course_end",
//	}
//}
//
//func createVisitsApplicationsSchema() map[string]string {
//	return map[string]string{
//		"tableName":     "visits_applications_view",
//		"key":           "visitsApplication",
//		"id":            "id",
//		"createdAt":     "created_at",
//		"formStatusId":  "form_status_id",
//		"email":         "email",
//		"childFullName": "child_full_name",
//		"gateName":      "gate_name",
//		"divisionName":  "division_name",
//		"withCar":       "with_car",
//	}
//}
//
//func createDpoApplicationsSchema() map[string]string {
//	return map[string]string{
//		"tableName":    "dpo_applications_view",
//		"key":          "dpoApplication",
//		"id":           "id",
//		"createdAt":    "created_at",
//		"isNmo":        "is_nmo",
//		"formStatusId": "form_status_id",
//		"email":        "email",
//		"fullName":     "full_name",
//		"courseName":   "course_name",
//	}
//}
//
//func createResidencyApplicationsSchema() map[string]string {
//	return map[string]string{
//		"tableName":          "residency_applications_view",
//		"key":                "residencyApplication",
//		"id":                 "id",
//		"createdAt":          "created_at",
//		"approvingDate":      "approving_date",
//		"formStatusId":       "form_status_id",
//		"email":              "email",
//		"pointsAchievements": "points_achievements",
//		"pointsEntrance":     "points_entrance",
//		"pointsSum":          "points_sum",
//		"startYear":          "start_year",
//		"endYear":            "end_year",
//		"fullName":           "full_name",
//		"main":               "main",
//		"paid":               "paid",
//		"courseName":         "course_name",
//		"admissionCommittee": "admission_committee",
//	}
//}
//
//func createPostgraduateApplicationsSchema() map[string]string {
//	return map[string]string{
//		"tableName":    "postgraduate_applications_view",
//		"key":          "postgraduateApplication",
//		"id":           "id",
//		"createdAt":    "created_at",
//		"formStatusId": "form_status_id",
//		"email":        "email",
//		"fullName":     "full_name",
//		"courseName":   "course_name",
//	}
//}
//
//func createDpoBaseCourseSchema() map[string]string {
//	return map[string]string{
//		"tableName": "dpo_base_courses",
//		"key":       "nmoCourse",
//		"id":        "id",
//		"name":      "name",
//		"hours":     "hours",
//		"teacherId": "teacher_id",
//		"listeners": "listeners",
//		"start":     "dpo_course_start",
//	}
//}
//
//func createEducationalManagerSchema() map[string]string {
//	return map[string]string{
//		"tableName": "educational_managers_view",
//		"key":       "educationalManager",
//		"order":     "educational_managers_view",
//		"id":        "id",
//		"fullName":  "fullName",
//	}
//}
//
//func createNmoCourseSpecializationSchema() map[string]string {
//	return map[string]string{
//		"tableName":        "nmo_courses_specializations",
//		"key":              "nmoCourseSpecialization",
//		"id":               "id",
//		"nmoCourseId":      "dpo_course_id",
//		"specializationId": "specialization_id",
//	}
//}
//
//func createSpecializationSchema() map[string]string {
//	return map[string]string{
//		"tableName":  "specializations",
//		"key":        "specialization",
//		"order":      "name",
//		"id":         "id",
//		"name":       "name",
//		"value":      "id",
//		"sortColumn": "name",
//		"label":      "name",
//	}
//}
//
//func createVacancySchema() map[string]string {
//	return map[string]string{
//		"tableName":         "vacancies_view",
//		"key":               "vacancy",
//		"title":             "title",
//		"minSalary":         "min_salary",
//		"divisionId":        "division_id",
//		"slug":              "slug",
//		"date":              "vacancy_date",
//		"value":             "slug",
//		"maxSalary":         "max_salary",
//		"sortColumn":        "title",
//		"responsesCount":    "responses_count",
//		"newResponsesCount": "new_responses_count",
//		"active":            "active",
//	}
//}
//
//func createVacancyResponseSchema() map[string]string {
//	return map[string]string{
//		"tableName":    "vacancy_responses_view",
//		"key":          "vacancyResponse",
//		"title":        "title",
//		"date":         "created_at",
//		"value":        "id",
//		"sortColumn":   "created_at",
//		"formStatusId": "form_status_id",
//		"email":        "email",
//		"fullName":     "full_name",
//	}
//}
//
//func createPostgraduateCourseSchema() map[string]string {
//	return map[string]string{
//		"tableName":     "postgraduate_courses_view",
//		"value":         "id",
//		"key":           "postgraduateCourse",
//		"name":          "name",
//		"cost":          "cost",
//		"code":          "code",
//		"years":         "years",
//		"educationForm": "education_form",
//	}
//}
//
//func createPostgraduateCourseSpecializationSchema() map[string]string {
//	return map[string]string{
//		"tableName":            "postgraduate_courses_specializations",
//		"key":                  "postgraduateCourseSpecialization",
//		"id":                   "id",
//		"postgraduateCourseId": "postgraduate_course_id",
//		"specializationId":     "specialization_id",
//	}
//}
//
//func createResidencyCourseSchema() map[string]string {
//	return map[string]string{
//		"id":            "id",
//		"tableName":     "residency_courses_view",
//		"value":         "id",
//		"key":           "residencyCourse",
//		"name":          "name",
//		"code":          "code",
//		"slug":          "slug",
//		"cost":          "cost",
//		"freePlaces":    "free_places",
//		"paidPlaces":    "paid_places",
//		"startYear":     "start_year",
//		"startYearId":   "start_year_id",
//		"endYear":       "end_year",
//		"endYearId":     "end_year_id",
//		"educationForm": "education_form",
//	}
//}
//
//func createEducationPublicDocumentTypeSchema() map[string]string {
//	return map[string]string{
//		"tableName":            "education_public_document_types",
//		"value":                "id",
//		"id":                   "id",
//		"key":                  "educationPublicDocumentType",
//		"publicDocumentTypeId": "public_document_type_id",
//	}
//}
//
//func createPublicDocumentTypeSchema() map[string]string {
//	return map[string]string{
//		"tableName": "public_document_types",
//		"value":     "id",
//		"id":        "id",
//		"key":       "publicDocumentTypeSchema",
//	}
//}
//
//func createEducationYearSchema() map[string]string {
//	return map[string]string{
//		"tableName":  "education_years",
//		"value":      "id",
//		"sortColumn": "year",
//		"label":      "year",
//		"id":         "id",
//		"key":        "educationYear",
//	}
//}
//
//func createEducationalOrganizationAcademicsSchema() map[string]string {
//	return map[string]string{
//		"tableName":  "educational_organization_academics_view",
//		"key":        "educationalOrganizationAcademic",
//		"id":         "id",
//		"fullName":   "full_name",
//		"value":      "id",
//		"sortColumn": "full_name",
//	}
//}
//
//func createRolesSchema() map[string]string {
//	return map[string]string{
//		"tableName":  "roles",
//		"key":        "role",
//		"name":       "name",
//		"id":         "id",
//		"value":      "id",
//		"label":      "label",
//		"sortColumn": "label",
//	}
//}
//
//func createPathPermissionsSchema() map[string]string {
//	return map[string]string{
//		"tableName":  "path_permissions",
//		"key":        "pathPermission",
//		"id":         "id",
//		"value":      "id",
//		"resource":   "resource",
//		"sortColumn": "resource",
//	}
//}
//
//func createNewsSchema() map[string]string {
//	return map[string]string{
//		"tableName":   "news_view",
//		"key":         "news",
//		"id":          "id",
//		"value":       "id",
//		"sortColumn":  "title",
//		"title":       "title",
//		"status":      "status",
//		"previewText": "preview_text",
//		"slug":        "slug",
//		"publishedOn": "published_on",
//		"createdAt":   "created_at",
//		"description": "description",
//		"main":        "main",
//		"subMain":     "sub_main",
//		"articleLink": "article_link",
//		"isArticle":   "is_article",
//		"viewsCount":  "views_count",
//		"isDraft":     "is_draft",
//	}
//}
//
//func createNewsToTagSchema() map[string]string {
//	return map[string]string{
//		"tableName": "news_to_tags",
//		"key":       "newsToTag",
//		"id":        "id",
//		"value":     "id",
//		"newsId":    "news_id",
//		"tagId":     "tag_id",
//	}
//}
//
//func createFormStatusSchema() map[string]string {
//	return map[string]string{
//		"tableName":         "form_statuses_view",
//		"key":               "formStatus",
//		"id":                "id",
//		"value":             "id",
//		"formStatusGroupId": "form_status_group_id",
//		"code":              "code",
//		"label":             "label",
//		"sortColumn":        "label",
//	}
//}
//
//func createQuestionSchema() map[string]string {
//	return map[string]string{
//		"tableName":    "questions",
//		"key":          "question",
//		"id":           "id",
//		"value":        "id",
//		"date":         "question_date",
//		"published":    "published",
//		"answered":     "answered",
//		"isNew":        "is_new",
//		"answerIsRead": "answer_is_read",
//		"theme":        "theme",
//	}
//}
//
//func createTreatDirectionSchema() map[string]string {
//	return map[string]string{
//		"tableName":  "treat_directions",
//		"key":        "treatDirection",
//		"id":         "id",
//		"value":      "id",
//		"name":       "name",
//		"sortColumn": "name",
//		"label":      "name",
//	}
//}
//
//func createDoctorCommentSchema() map[string]string {
//	return map[string]string{
//		"tableName": "doctor_comments",
//		"key":       "doctorComment",
//		"id":        "id",
//		"commentId": "comment_id",
//		"doctorId":  "doctor_id",
//	}
//}
//
//func createDivisionCommentSchema() map[string]string {
//	return map[string]string{
//		"tableName": "division_comments",
//		"key":       "divisionComment",
//		"id":        "id",
//		"commentId": "comment_id",
//		"doctorId":  "division_id",
//	}
//}
//
//func createNewsCommentSchema() map[string]string {
//	return map[string]string{
//		"tableName": "news_comments",
//		"key":       "newsComment",
//		"id":        "id",
//		"commentId": "comment_id",
//		"doctorId":  "doctor_id",
//	}
//}
//
//func createUserSchema() map[string]string {
//	return map[string]string{
//		"tableName": "users_view",
//		"key":       "users",
//		"id":        "id",
//		"email":     "email",
//		"fullName":  "full_name",
//	}
//}
//
//func createBuildingSchema() map[string]string {
//	return map[string]string{
//		"tableName":  "buildings",
//		"key":        "building",
//		"id":         "id",
//		"name":       "name",
//		"address":    "address",
//		"number":     "number",
//		"value":      "id",
//		"label":      "name",
//		"sortColumn": "name",
//	}
//}
//
//func createDietSchema() map[string]string {
//	return map[string]string{
//		"tableName":   "diets",
//		"key":         "diet",
//		"id":          "id",
//		"diabetes":    "diabetes",
//		"agePeriodId": "age_period_id",
//	}
//}
//
//func createAgePeriodSchema() map[string]string {
//	return map[string]string{
//		"tableName":  "age_periods",
//		"key":        "age_period",
//		"id":         "id",
//		"name":       "name",
//		"value":      "id",
//		"label":      "name",
//		"sortColumn": "name",
//	}
//}
//
//func createDoctorDivisionSchema() map[string]string {
//	return map[string]string{
//		"tableName":  "doctors_divisions",
//		"key":        "doctorDivision",
//		"id":         "id",
//		"value":      "id",
//		"divisionId": "division_id",
//		"doctorId":   "doctor_id",
//	}
//}
//
//// Переведено на анализ моделей
////func createEmployeesSchema() map[string]string {
////	return map[string]string{
////		"tableName": "employees_view",
////		"key":       "employee",
////		"id":        "id",
////		"fullName":  "full_name",
////		"isMale":    "is_male",
////		"dateBirth": "date_birth",
////	}
////}
//
//func createHospitalizationSchema() map[string]string {
//	return map[string]string{
//		"tableName":     "hospitalizations_view",
//		"key":           "hospitalization",
//		"id":            "id",
//		"date":          "hospitalization_date",
//		"divisionId":    "division_id",
//		"is_new":        "is_new",
//		"createdAt":     "created_at",
//		"approvingDate": "approving_date",
//		"formStatusId":  "form_status_id",
//		"email":         "email",
//		"fullName":      "full_name",
//		"policyType":    "policy_type",
//		"treatmentType": "treatment_type",
//		"stayType":      "stay_type",
//		"referralType":  "referral_type",
//	}
//}
//
//func createSupportMessageSchema() map[string]string {
//	return map[string]string{
//		"tableName": "support_messages",
//		"key":       "supportMessage",
//		"id":        "id",
//		"value":     "id",
//		"date":      "support_message_date",
//		"isNew":     "is_new",
//		"theme":     "theme",
//		"question":  "question",
//		"answer":    "answer",
//	}
//}
//
//func createDailyMenuSchema() map[string]string {
//	return map[string]string{
//		"tableName": "daily_menus",
//		"key":       "dailyMenu",
//		"id":        "id",
//		"value":     "id",
//		"date":      "item_date",
//		"active":    "active",
//		"order":     "item_order",
//		"name":      "name",
//	}
//}
//
//func createFormPatternSchema() map[string]string {
//	return map[string]string{
//		"tableName": "form_patterns",
//		"id":        "id",
//		"code":      "code",
//		"name":      "name",
//	}
//}
//
//func createDishesGroupSchema() map[string]string {
//	return map[string]string{
//		"tableName": "dishes_groups",
//		"id":        "id",
//		"name":      "name",
//		"order":     "dishes_group_order",
//	}
//}
//
//func createDailyMenuOrderSchema() map[string]string {
//	return map[string]string{
//		"tableName":    "daily_menu_orders_view",
//		"key":          "dailyMenuOrder",
//		"id":           "id",
//		"createdAt":    "created_at",
//		"formStatusId": "form_status_id",
//		"email":        "email",
//		"date":         "item_date",
//		"boxNumber":    "box_number",
//		"number":       "number",
//		"formValueId":  "formValueId",
//	}
//}
