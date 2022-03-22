package routing

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/pro-assistance/pro-assister/config"
	"mdgkb/mdgkb-server/handlers/applicationsCars"
	"mdgkb/mdgkb-server/handlers/appointments"
	"mdgkb/mdgkb-server/handlers/auth"
	"mdgkb/mdgkb-server/handlers/banners"
	"mdgkb/mdgkb-server/handlers/callbackRequests"
	"mdgkb/mdgkb-server/handlers/centers"
	certificates "mdgkb/mdgkb-server/handlers/certificates"
	"mdgkb/mdgkb-server/handlers/children"
	"mdgkb/mdgkb-server/handlers/comments"
	"mdgkb/mdgkb-server/handlers/divisions"
	"mdgkb/mdgkb-server/handlers/doctors"
	"mdgkb/mdgkb-server/handlers/documentTypes"
	"mdgkb/mdgkb-server/handlers/donorRules"
	"mdgkb/mdgkb-server/handlers/dpoApplications"
	"mdgkb/mdgkb-server/handlers/dpoCourses"
	"mdgkb/mdgkb-server/handlers/educationalManagers"
	"mdgkb/mdgkb-server/handlers/educationalOrganization"
	"mdgkb/mdgkb-server/handlers/entrances"
	"mdgkb/mdgkb-server/handlers/events"
	"mdgkb/mdgkb-server/handlers/faqs"
	"mdgkb/mdgkb-server/handlers/gates"
	"mdgkb/mdgkb-server/handlers/heads"
	"mdgkb/mdgkb-server/handlers/medicalProfiles"
	"mdgkb/mdgkb-server/handlers/menus"
	"mdgkb/mdgkb-server/handlers/meta"
	"mdgkb/mdgkb-server/handlers/news"
	"mdgkb/mdgkb-server/handlers/newsSlides"
	"mdgkb/mdgkb-server/handlers/pages"
	"mdgkb/mdgkb-server/handlers/paidPrograms"
	paidProgramsGroups "mdgkb/mdgkb-server/handlers/paidProgramsGroups"
	"mdgkb/mdgkb-server/handlers/paidServices"
	"mdgkb/mdgkb-server/handlers/partnerTypes"
	"mdgkb/mdgkb-server/handlers/partners"
	"mdgkb/mdgkb-server/handlers/preparations"
	"mdgkb/mdgkb-server/handlers/projects"
	"mdgkb/mdgkb-server/handlers/publicDocumentTypes"
	"mdgkb/mdgkb-server/handlers/questions"
	"mdgkb/mdgkb-server/handlers/search"
	"mdgkb/mdgkb-server/handlers/specializations"
	"mdgkb/mdgkb-server/handlers/teachers"
	"mdgkb/mdgkb-server/handlers/timetablePatterns"
	"mdgkb/mdgkb-server/handlers/formPatterns"
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/handlers/vacancies"
	"mdgkb/mdgkb-server/handlers/vacancyResponse"
	"mdgkb/mdgkb-server/handlers/valueTypes"
	"mdgkb/mdgkb-server/handlers/visitingRules"
	"mdgkb/mdgkb-server/middleware"
	applicationsCarsRouter "mdgkb/mdgkb-server/routing/applicationsCars"
	appointmentsRouter "mdgkb/mdgkb-server/routing/appointments"
	authRouter "mdgkb/mdgkb-server/routing/auth"
	bannersRouter "mdgkb/mdgkb-server/routing/banners"
	"mdgkb/mdgkb-server/routing/buildings"
	callbackRequestsRouter "mdgkb/mdgkb-server/routing/callbackRequests"
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
	educationalManagersRouter "mdgkb/mdgkb-server/routing/educationalManagers"
	educationalOraganizationRouter "mdgkb/mdgkb-server/routing/educationalOraganization"
	entrancesRouter "mdgkb/mdgkb-server/routing/entrances"
	eventsRouter "mdgkb/mdgkb-server/routing/events"
	faqRouter "mdgkb/mdgkb-server/routing/faqs"
	gatesRouter "mdgkb/mdgkb-server/routing/gates"
	headsRouter "mdgkb/mdgkb-server/routing/heads"
	hospitalizationRouter "mdgkb/mdgkb-server/routing/hospitalization"
	medicalProfilesRouter "mdgkb/mdgkb-server/routing/medicalProfiles"
	menusRouter "mdgkb/mdgkb-server/routing/menus"
	metaRouter "mdgkb/mdgkb-server/routing/meta"
	newsRouter "mdgkb/mdgkb-server/routing/news"
	newsSlidesRouter "mdgkb/mdgkb-server/routing/newsSlides"
	"mdgkb/mdgkb-server/routing/normativeDocumentTypes"
	"mdgkb/mdgkb-server/routing/normativeDocuments"
	pagesRouter "mdgkb/mdgkb-server/routing/pages"
	paidProgramsRouter "mdgkb/mdgkb-server/routing/paidPrograms"
	paidProgramsGroupsRouter "mdgkb/mdgkb-server/routing/paidProgramsGroups"
	paidServicesRouter "mdgkb/mdgkb-server/routing/paidServices"
	partnerTypesRouter "mdgkb/mdgkb-server/routing/partnerTypes"
	partnersRouter "mdgkb/mdgkb-server/routing/partners"
	preparationsRouter "mdgkb/mdgkb-server/routing/preparations"
	projectsRouter "mdgkb/mdgkb-server/routing/projects"
	publicDocumentTypesRouter "mdgkb/mdgkb-server/routing/publicDocumentTypes"
	questionsRouter "mdgkb/mdgkb-server/routing/questions"
	searchRouter "mdgkb/mdgkb-server/routing/search"
	"mdgkb/mdgkb-server/routing/sideOrganizations"
	specializationsRouter "mdgkb/mdgkb-server/routing/specializations"
	"mdgkb/mdgkb-server/routing/tags"
	teachersRouter "mdgkb/mdgkb-server/routing/teachers"
	timetablePatternsRouter "mdgkb/mdgkb-server/routing/timetablePatterns"
	formPatternsRouter "mdgkb/mdgkb-server/routing/formPatterns"
	"mdgkb/mdgkb-server/routing/timetables"
	usersRouter "mdgkb/mdgkb-server/routing/users"
	vacanciesRouter "mdgkb/mdgkb-server/routing/vacancies"
	vacancyResponseRouter "mdgkb/mdgkb-server/routing/vacancyResponse"
	valueTypesRouter "mdgkb/mdgkb-server/routing/valueTypes"
	visitingRulesRouter "mdgkb/mdgkb-server/routing/visitingRules"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/go-redis/redis/v7"
	"github.com/uptrace/bun"

	helperPack "github.com/pro-assistance/pro-assister/helper"
)

func Init(r *gin.Engine, db *bun.DB, redisClient *redis.Client, elasticSearchClient *elasticsearch.Client, config config.Config) {
	localUploader := helperPack.NewLocalUploader(&config.UploadPath)
	helper := helperPack.NewHelper(config)

	r.Static("/static", "./static/")
	authGroup := r.Group("/api/v1/auth")
	authRouter.Init(authGroup.Group(""), auth.CreateHandler(db, helper))

	api := r.Group("/api/v1")
	m := middleware.CreateMiddleware(helper)
	api.Use(m.Authentication())
	bannersRouter.Init(api.Group("/banners"), banners.CreateHandler(db, helper))
	buildings.Init(api.Group("/buildings"), db, localUploader)
	doctorsRouter.Init(api.Group("/doctors"), doctors.CreateHandler(db, helper))
	hospitalizationRouter.Init(api.Group("/hospitalizations"), db, helper)

	divisionsRouter.Init(api.Group("/divisions"), divisions.CreateHandler(db, helper))
	headsRouter.Init(api.Group("/heads"), heads.CreateHandler(db, helper))

	commentsRouter.Init(api.Group("/comments"), comments.CreateHandler(db, helper))
	newsRouter.Init(api.Group("/news"), news.CreateHandler(db, helper))
	normativeDocumentTypes.Init(api.Group("/normative-document-types"), db, localUploader)
	normativeDocuments.Init(api.Group("/normative-documents"), db, localUploader)
	sideOrganizations.Init(api.Group("/side-organizations"), db, localUploader)
	tags.Init(api.Group("/tags"), db, localUploader)
	usersRouter.Init(api.Group("/users"), users.CreateHandler(db, helper))
	timetables.Init(api.Group("/timetables"), db)

	educationalOraganizationRouter.Init(api.Group("/educational-organization"), educationalOrganization.CreateHandler(db, helper))
	menusRouter.Init(api.Group("/menus"), menus.CreateHandler(db, helper))
	pagesRouter.Init(api.Group("/pages"), pages.CreateHandler(db, helper))
	projectsRouter.Init(api.Group("/projects"), projects.CreateHandler(db, helper))
	entrancesRouter.Init(api.Group("/entrances"), entrances.CreateHandler(db, helper))
	vacanciesRouter.Init(api.Group("/vacancies"), vacancies.CreateHandler(db, helper))
	vacancyResponseRouter.Init(api.Group("/vacancy-responses"), vacancyResponse.CreateHandler(db, helper))
	documentTypesRouter.Init(api.Group("/document-types"), documentTypes.CreateHandler(db, helper))
	valueTypesRouter.Init(api.Group("/value-types"), valueTypes.CreateHandler(db, helper))
	searchRouter.Init(api.Group("/search"), search.CreateHandler(db, helper, elasticSearchClient))
	faqRouter.Init(api.Group("/faqs"), faqs.CreateHandler(db, helper))
	visitingRulesRouter.Init(api.Group("/visiting-rules"), visitingRules.CreateHandler(db, helper))
	newsSlidesRouter.Init(api.Group("/news-slides"), newsSlides.CreateHandler(db, helper))
	questionsRouter.Init(api.Group("/questions"), questions.CreateHandler(db, helper))
	eventsRouter.Init(api.Group("/events"), events.CreateHandler(db, helper))
	timetablePatternsRouter.Init(api.Group("/timetable-patterns"), timetablePatterns.CreateHandler(db, helper))
	formPatternsRouter.Init(api.Group("/form-patterns"), formPatterns.CreateHandler(db, helper))
	paidProgramsRouter.Init(api.Group("/paid-programs"), paidPrograms.CreateHandler(db, helper))
	paidProgramsGroupsRouter.Init(api.Group("/paid-programs-groups"), paidProgramsGroups.CreateHandler(db, helper))
	partnerTypesRouter.Init(api.Group("/partner-types"), partnerTypes.CreateHandler(db, helper))
	publicDocumentTypesRouter.Init(api.Group("/public-document-types"), publicDocumentTypes.CreateHandler(db, helper))
	partnersRouter.Init(api.Group("/partners"), partners.CreateHandler(db, helper))
	preparationsRouter.Init(api.Group("/preparations"), preparations.CreateHandler(db, helper))
	donorRulesRouter.Init(api.Group("/donor-rules"), donorRules.CreateHandler(db, helper))
	certificatesRouter.Init(api.Group("/certificates"), certificates.CreateHandler(db, helper))
	metaRouter.Init(api.Group("/meta"), meta.CreateHandler(db, helper))
	paidServicesRouter.Init(api.Group("/paid-services"), paidServices.CreateHandler(db, helper))
	medicalProfilesRouter.Init(api.Group("/medical-profiles"), medicalProfiles.CreateHandler(db, helper))
	callbackRequestsRouter.Init(api.Group("/callback-requests"), callbackRequests.CreateHandler(db, helper))
	applicationsCarsRouter.Init(api.Group("/applications-cars"), applicationsCars.CreateHandler(db, helper))
	centersRouter.Init(api.Group("/centers"), centers.CreateHandler(db, helper))
	dpoCoursesRouter.Init(api.Group("/dpo-courses"), dpoCourses.CreateHandler(db, helper))
	dpoApplicationsRouter.Init(api.Group("/dpo-applications"), dpoApplications.CreateHandler(db, helper))
	teachersRouter.Init(api.Group("/teachers"), teachers.CreateHandler(db, helper))
	educationalManagersRouter.Init(api.Group("/educational-managers"), educationalManagers.CreateHandler(db, helper))
	appointmentsRouter.Init(api.Group("/appointments"), appointments.CreateHandler(db, helper))
	childrenRouter.Init(api.Group("/children"), children.CreateHandler(db, helper))
	gatesRouter.Init(api.Group("/gates"), gates.CreateHandler(db, helper))
	specializationsRouter.Init(api.Group("/specializations"), specializations.CreateHandler(db, helper))
}
