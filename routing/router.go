package routing

import (
	"github.com/elastic/go-elasticsearch/v8"
	"mdgkb/mdgkb-server/handlers/admissionCommitteeDocumentTypes"
	"mdgkb/mdgkb-server/handlers/appointments"
	"mdgkb/mdgkb-server/handlers/auth"
	"mdgkb/mdgkb-server/handlers/banners"
	"mdgkb/mdgkb-server/handlers/buildings"
	"mdgkb/mdgkb-server/handlers/callbackRequests"
	"mdgkb/mdgkb-server/handlers/candidateApplications"
	"mdgkb/mdgkb-server/handlers/candidateDocumentTypes"
	"mdgkb/mdgkb-server/handlers/candidateExams"
	"mdgkb/mdgkb-server/handlers/centers"
	"mdgkb/mdgkb-server/handlers/certificates"
	"mdgkb/mdgkb-server/handlers/children"
	"mdgkb/mdgkb-server/handlers/comments"
	"mdgkb/mdgkb-server/handlers/divisions"
	"mdgkb/mdgkb-server/handlers/doctors"
	"mdgkb/mdgkb-server/handlers/documentTypes"
	"mdgkb/mdgkb-server/handlers/donorRules"
	"mdgkb/mdgkb-server/handlers/dpoApplications"
	"mdgkb/mdgkb-server/handlers/dpoCourses"
	"mdgkb/mdgkb-server/handlers/dpoDocumentTypes"
	"mdgkb/mdgkb-server/handlers/educationPublicDocumentTypes"
	"mdgkb/mdgkb-server/handlers/educationYears"
	"mdgkb/mdgkb-server/handlers/educationalManagers"
	"mdgkb/mdgkb-server/handlers/educationalOrganization"
	"mdgkb/mdgkb-server/handlers/educationalOrganizationAcademics"
	"mdgkb/mdgkb-server/handlers/entrances"
	"mdgkb/mdgkb-server/handlers/events"
	"mdgkb/mdgkb-server/handlers/faqs"
	"mdgkb/mdgkb-server/handlers/formPatterns"
	"mdgkb/mdgkb-server/handlers/formStatusGroups"
	"mdgkb/mdgkb-server/handlers/formStatuses"
	"mdgkb/mdgkb-server/handlers/formValues"
	"mdgkb/mdgkb-server/handlers/gates"
	"mdgkb/mdgkb-server/handlers/heads"
	"mdgkb/mdgkb-server/handlers/medicalProfiles"
	"mdgkb/mdgkb-server/handlers/menus"
	"mdgkb/mdgkb-server/handlers/meta"
	"mdgkb/mdgkb-server/handlers/news"
	"mdgkb/mdgkb-server/handlers/newsSlides"
	"mdgkb/mdgkb-server/handlers/pages"
	"mdgkb/mdgkb-server/handlers/paidPrograms"
	"mdgkb/mdgkb-server/handlers/paidProgramsGroups"
	"mdgkb/mdgkb-server/handlers/paidServices"
	"mdgkb/mdgkb-server/handlers/partnerTypes"
	"mdgkb/mdgkb-server/handlers/partners"
	"mdgkb/mdgkb-server/handlers/pointsAchievements"
	"mdgkb/mdgkb-server/handlers/postgraduateApplications"
	"mdgkb/mdgkb-server/handlers/postgraduateCourses"
	"mdgkb/mdgkb-server/handlers/postgraduateDocumentTypes"
	"mdgkb/mdgkb-server/handlers/preparations"
	"mdgkb/mdgkb-server/handlers/projects"
	"mdgkb/mdgkb-server/handlers/publicDocumentTypes"
	"mdgkb/mdgkb-server/handlers/questions"
	"mdgkb/mdgkb-server/handlers/residencyApplications"
	"mdgkb/mdgkb-server/handlers/residencyCourses"
	"mdgkb/mdgkb-server/handlers/residencyDocumentTypes"
	"mdgkb/mdgkb-server/handlers/roles"
	"mdgkb/mdgkb-server/handlers/search"
	"mdgkb/mdgkb-server/handlers/sideOrganizations"
	"mdgkb/mdgkb-server/handlers/specializations"
	"mdgkb/mdgkb-server/handlers/tags"
	"mdgkb/mdgkb-server/handlers/teachers"
	"mdgkb/mdgkb-server/handlers/timetablePatterns"
	"mdgkb/mdgkb-server/handlers/timetables"
	"mdgkb/mdgkb-server/handlers/treatDirections"
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/handlers/vacancies"
	"mdgkb/mdgkb-server/handlers/vacancyResponse"
	"mdgkb/mdgkb-server/handlers/valueTypes"
	"mdgkb/mdgkb-server/handlers/visitingRules"
	"mdgkb/mdgkb-server/handlers/visitsApplications"
	"mdgkb/mdgkb-server/middleware"
	admissionCommitteeDocumentTypesRouter "mdgkb/mdgkb-server/routing/admissionCommitteeDocumentTypes"
	appointmentsRouter "mdgkb/mdgkb-server/routing/appointments"
	authRouter "mdgkb/mdgkb-server/routing/auth"
	bannersRouter "mdgkb/mdgkb-server/routing/banners"
	buildingsRouter "mdgkb/mdgkb-server/routing/buildings"
	callbackRequestsRouter "mdgkb/mdgkb-server/routing/callbackRequests"
	candidateApplicationsRouter "mdgkb/mdgkb-server/routing/candidateApplications"
	candidateDocumentTypesRouter "mdgkb/mdgkb-server/routing/candidateDocumentTypes"
	candidateExamsRouter "mdgkb/mdgkb-server/routing/candidateExams"
	centersRouter "mdgkb/mdgkb-server/routing/centers"
	certificatesRouter "mdgkb/mdgkb-server/routing/certificates"
	childrenRouter "mdgkb/mdgkb-server/routing/children"
	commentsRouter "mdgkb/mdgkb-server/routing/comments"
	divisionsRouter "mdgkb/mdgkb-server/routing/divisions"
	doctorsRouter "mdgkb/mdgkb-server/routing/doctors"
	documentTypesRouter "mdgkb/mdgkb-server/routing/document-types"
	donorRulesRouter "mdgkb/mdgkb-server/routing/donorRules"
	dpoApplicationsRouter "mdgkb/mdgkb-server/routing/dpoApplications"
	dpoCoursesRouter "mdgkb/mdgkb-server/routing/dpoCourses"
	dpoDocumentTypesRouter "mdgkb/mdgkb-server/routing/dpoDocumentTypes"
	educationPublicDocumentTypesRouter "mdgkb/mdgkb-server/routing/educationPublicDocumentTypes"
	educationYearsRouter "mdgkb/mdgkb-server/routing/educationYears"
	educationalManagersRouter "mdgkb/mdgkb-server/routing/educationalManagers"
	educationalOraganizationRouter "mdgkb/mdgkb-server/routing/educationalOraganization"
	educationalOrganizationAcademicsRouter "mdgkb/mdgkb-server/routing/educationalOrganizationAcademics"
	entrancesRouter "mdgkb/mdgkb-server/routing/entrances"
	eventsRouter "mdgkb/mdgkb-server/routing/events"
	faqRouter "mdgkb/mdgkb-server/routing/faqs"
	formPatternsRouter "mdgkb/mdgkb-server/routing/formPatterns"
	formStatusGroupsRouter "mdgkb/mdgkb-server/routing/formStatusGroups"
	formStatusesRouter "mdgkb/mdgkb-server/routing/formStatuses"
	formValuesRouter "mdgkb/mdgkb-server/routing/formValues"
	gatesRouter "mdgkb/mdgkb-server/routing/gates"
	headsRouter "mdgkb/mdgkb-server/routing/heads"
	hospitalizationRouter "mdgkb/mdgkb-server/routing/hospitalization"
	medicalProfilesRouter "mdgkb/mdgkb-server/routing/medicalProfiles"
	menusRouter "mdgkb/mdgkb-server/routing/menus"
	metaRouter "mdgkb/mdgkb-server/routing/meta"
	newsRouter "mdgkb/mdgkb-server/routing/news"
	newsSlidesRouter "mdgkb/mdgkb-server/routing/newsSlides"
	pagesRouter "mdgkb/mdgkb-server/routing/pages"
	paidProgramsRouter "mdgkb/mdgkb-server/routing/paidPrograms"
	paidProgramsGroupsRouter "mdgkb/mdgkb-server/routing/paidProgramsGroups"
	paidServicesRouter "mdgkb/mdgkb-server/routing/paidServices"
	partnerTypesRouter "mdgkb/mdgkb-server/routing/partnerTypes"
	partnersRouter "mdgkb/mdgkb-server/routing/partners"
	pointsAchievementsRouter "mdgkb/mdgkb-server/routing/pointsAchievements"
	postgraduateApplicationsRouter "mdgkb/mdgkb-server/routing/postgraduateApplications"
	postgraduateCoursesRouter "mdgkb/mdgkb-server/routing/postgraduateCourses"
	postgraduateDocumentTypesRouter "mdgkb/mdgkb-server/routing/postgraduateDocumentTypes"
	preparationsRouter "mdgkb/mdgkb-server/routing/preparations"
	projectsRouter "mdgkb/mdgkb-server/routing/projects"
	publicDocumentTypesRouter "mdgkb/mdgkb-server/routing/publicDocumentTypes"
	questionsRouter "mdgkb/mdgkb-server/routing/questions"
	residencyApplicationsRouter "mdgkb/mdgkb-server/routing/residencyApplications"
	residencyCoursesRouter "mdgkb/mdgkb-server/routing/residencyCourses"
	residencyDocumentTypesRouter "mdgkb/mdgkb-server/routing/residencyDocumentTypes"
	rolesRouter "mdgkb/mdgkb-server/routing/roles"
	searchRouter "mdgkb/mdgkb-server/routing/search"
	sideOrganizationsRouter "mdgkb/mdgkb-server/routing/sideOrganizations"
	specializationsRouter "mdgkb/mdgkb-server/routing/specializations"
	tagsRouter "mdgkb/mdgkb-server/routing/tags"
	teachersRouter "mdgkb/mdgkb-server/routing/teachers"
	timetablePatternsRouter "mdgkb/mdgkb-server/routing/timetablePatterns"
	timetablesRouter "mdgkb/mdgkb-server/routing/timetables"
	treatDirectionsRouter "mdgkb/mdgkb-server/routing/treatDirections"
	usersRouter "mdgkb/mdgkb-server/routing/users"
	vacanciesRouter "mdgkb/mdgkb-server/routing/vacancies"
	vacancyResponseRouter "mdgkb/mdgkb-server/routing/vacancyResponse"
	valueTypesRouter "mdgkb/mdgkb-server/routing/valueTypes"
	visitingRulesRouter "mdgkb/mdgkb-server/routing/visitingRules"
	visitsApplicationsRouter "mdgkb/mdgkb-server/routing/visitsApplications"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	helperPack "github.com/pro-assistance/pro-assister/helper"
)

func Init(r *gin.Engine, helper *helperPack.Helper, elasticSearchClient elasticsearch.Client) {
	m := middleware.CreateMiddleware(helper)

	r.Use(m.CORSMiddleware())
	//r.Use(m.CheckPermission())
	r.Use(gin.Logger())

	r.Static("/api/v1/static", "./static/")
	authGroup := r.Group("/api/v1/auth")
	authRouter.Init(authGroup.Group(""), auth.CreateHandler(helper))

	api := r.Group("/api/v1")
	//api.Use(m.Authentication())
	api.Use(m.CORSMiddleware())
	api.GET("/subscribe/:channel", helper.Broker.ServeHTTP)

	bannersRouter.Init(api.Group("/banners"), banners.CreateHandler(helper))
	buildingsRouter.Init(api.Group("/buildings"), buildings.CreateHandler(helper))
	doctorsRouter.Init(api.Group("/doctors"), doctors.CreateHandler(helper))
	hospitalizationRouter.Init(api.Group("/hospitalizations"), helper)
	divisionsRouter.Init(api.Group("/divisions"), divisions.CreateHandler(helper))
	headsRouter.Init(api.Group("/heads"), heads.CreateHandler(helper))
	commentsRouter.Init(api.Group("/comments"), comments.CreateHandler(helper))
	newsRouter.Init(api.Group("/news"), news.CreateHandler(helper))
	sideOrganizationsRouter.Init(api.Group("/side-organizations"), sideOrganizations.CreateHandler(helper))
	tagsRouter.Init(api.Group("/tags"), tags.CreateHandler(helper))
	usersRouter.Init(api.Group("/users"), users.CreateHandler(helper))
	timetablesRouter.Init(api.Group("/timetables"), timetables.CreateHandler(helper))
	educationalOraganizationRouter.Init(api.Group("/educational-organization"), educationalOrganization.CreateHandler(helper))
	menusRouter.Init(api.Group("/menus"), menus.CreateHandler(helper))
	pagesRouter.Init(api.Group("/pages"), pages.CreateHandler(helper))
	projectsRouter.Init(api.Group("/projects"), projects.CreateHandler(helper))
	entrancesRouter.Init(api.Group("/entrances"), entrances.CreateHandler(helper))
	vacanciesRouter.Init(api.Group("/vacancies"), vacancies.CreateHandler(helper))
	vacancyResponseRouter.Init(api.Group("/vacancy-responses"), vacancyResponse.CreateHandler(helper))
	documentTypesRouter.Init(api.Group("/document-types"), documentTypes.CreateHandler(helper))
	valueTypesRouter.Init(api.Group("/value-types"), valueTypes.CreateHandler(helper))
	searchRouter.Init(api.Group("/search"), search.CreateHandler(helper))
	faqRouter.Init(api.Group("/faqs"), faqs.CreateHandler(helper))
	visitingRulesRouter.Init(api.Group("/visiting-rules"), visitingRules.CreateHandler(helper))
	newsSlidesRouter.Init(api.Group("/news-slides"), newsSlides.CreateHandler(helper))
	questionsRouter.Init(api.Group("/questions"), questions.CreateHandler(helper))
	eventsRouter.Init(api.Group("/events"), events.CreateHandler(helper))
	timetablePatternsRouter.Init(api.Group("/timetable-patterns"), timetablePatterns.CreateHandler(helper))
	formPatternsRouter.Init(api.Group("/form-patterns"), formPatterns.CreateHandler(helper))
	paidProgramsRouter.Init(api.Group("/paid-programs"), paidPrograms.CreateHandler(helper))
	paidProgramsGroupsRouter.Init(api.Group("/paid-programs-groups"), paidProgramsGroups.CreateHandler(helper))
	partnerTypesRouter.Init(api.Group("/partner-types"), partnerTypes.CreateHandler(helper))
	publicDocumentTypesRouter.Init(api.Group("/public-document-types"), publicDocumentTypes.CreateHandler(helper))
	partnersRouter.Init(api.Group("/partners"), partners.CreateHandler(helper))
	preparationsRouter.Init(api.Group("/preparations"), preparations.CreateHandler(helper))
	donorRulesRouter.Init(api.Group("/donor-rules"), donorRules.CreateHandler(helper))
	certificatesRouter.Init(api.Group("/certificates"), certificates.CreateHandler(helper))
	metaRouter.Init(api.Group("/meta"), meta.CreateHandler(helper))
	paidServicesRouter.Init(api.Group("/paid-services"), paidServices.CreateHandler(helper))
	medicalProfilesRouter.Init(api.Group("/medical-profiles"), medicalProfiles.CreateHandler(helper))
	treatDirectionsRouter.Init(api.Group("/treat-directions"), treatDirections.CreateHandler(helper))
	callbackRequestsRouter.Init(api.Group("/callback-requests"), callbackRequests.CreateHandler(helper))
	visitsApplicationsRouter.Init(api.Group("/visits-applications"), visitsApplications.CreateHandler(helper))
	centersRouter.Init(api.Group("/centers"), centers.CreateHandler(helper))
	dpoCoursesRouter.Init(api.Group("/dpo-courses"), dpoCourses.CreateHandler(helper))
	postgraduateCoursesRouter.Init(api.Group("/postgraduate-courses"), postgraduateCourses.CreateHandler(helper))
	dpoApplicationsRouter.Init(api.Group("/dpo-applications"), dpoApplications.CreateHandler(helper))
	educationalOrganizationAcademicsRouter.Init(api.Group("/educational-organization-academics"), educationalOrganizationAcademics.CreateHandler(helper))
	residencyApplicationsRouter.Init(api.Group("/residency-applications"), residencyApplications.CreateHandler(helper))
	formValuesRouter.Init(api.Group("/form-values"), formValues.CreateHandler(helper))
	formStatusesRouter.Init(api.Group("/form-statuses"), formStatuses.CreateHandler(helper))
	formStatusGroupsRouter.Init(api.Group("/form-status-groups"), formStatusGroups.CreateHandler(helper))
	postgraduateApplicationsRouter.Init(api.Group("/postgraduate-applications"), postgraduateApplications.CreateHandler(helper))
	teachersRouter.Init(api.Group("/teachers"), teachers.CreateHandler(helper))
	educationalManagersRouter.Init(api.Group("/educational-managers"), educationalManagers.CreateHandler(helper))
	appointmentsRouter.Init(api.Group("/appointments"), appointments.CreateHandler(helper))
	childrenRouter.Init(api.Group("/children"), children.CreateHandler(helper))
	gatesRouter.Init(api.Group("/gates"), gates.CreateHandler(helper))
	specializationsRouter.Init(api.Group("/specializations"), specializations.CreateHandler(helper))
	candidateApplicationsRouter.Init(api.Group("/candidate-applications"), candidateApplications.CreateHandler(helper))
	candidateExamsRouter.Init(api.Group("/candidate-exams"), candidateExams.CreateHandler(helper))
	postgraduateDocumentTypesRouter.Init(api.Group("/postgraduate-document-types"), postgraduateDocumentTypes.CreateHandler(helper))
	dpoDocumentTypesRouter.Init(api.Group("/dpo-document-types"), dpoDocumentTypes.CreateHandler(helper))
	candidateDocumentTypesRouter.Init(api.Group("/candidate-document-types"), candidateDocumentTypes.CreateHandler(helper))
	rolesRouter.Init(api.Group("/roles"), roles.CreateHandler(helper))
	residencyCoursesRouter.Init(api.Group("/residency-courses"), residencyCourses.CreateHandler(helper))
	residencyDocumentTypesRouter.Init(api.Group("/residency-document-types"), residencyDocumentTypes.CreateHandler(helper))
	educationYearsRouter.Init(api.Group("/education-years"), educationYears.CreateHandler(helper))
	educationPublicDocumentTypesRouter.Init(api.Group("/education-public-document-types"), educationPublicDocumentTypes.CreateHandler(helper))
	admissionCommitteeDocumentTypesRouter.Init(api.Group("/admission-committee-document-types"), admissionCommitteeDocumentTypes.CreateHandler(helper))
	pointsAchievementsRouter.Init(api.Group("/points-achievements"), pointsAchievements.CreateHandler(helper))
}
