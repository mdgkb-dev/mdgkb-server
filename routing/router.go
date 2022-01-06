package routing

import (
	"mdgkb/mdgkb-server/config"
	"mdgkb/mdgkb-server/handlers/auth"
	"mdgkb/mdgkb-server/handlers/comments"
	"mdgkb/mdgkb-server/handlers/divisions"
	"mdgkb/mdgkb-server/handlers/doctors"
	"mdgkb/mdgkb-server/handlers/documentTypes"
	"mdgkb/mdgkb-server/handlers/donorRules"
	"mdgkb/mdgkb-server/handlers/events"
	"mdgkb/mdgkb-server/handlers/faqs"
	"mdgkb/mdgkb-server/handlers/heads"
	"mdgkb/mdgkb-server/handlers/menus"
	"mdgkb/mdgkb-server/handlers/news"
	"mdgkb/mdgkb-server/handlers/newsSlides"
	"mdgkb/mdgkb-server/handlers/pages"
	"mdgkb/mdgkb-server/handlers/paidPrograms"
	paidProgramsGroups "mdgkb/mdgkb-server/handlers/paidProgramsGroups"
	"mdgkb/mdgkb-server/handlers/partnerTypes"
	"mdgkb/mdgkb-server/handlers/partners"
	"mdgkb/mdgkb-server/handlers/preparations"
	"mdgkb/mdgkb-server/handlers/projects"
	"mdgkb/mdgkb-server/handlers/questions"
	"mdgkb/mdgkb-server/handlers/search"
	"mdgkb/mdgkb-server/handlers/timetablePatterns"
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/handlers/vacancies"
	"mdgkb/mdgkb-server/handlers/vacancyResponse"
	"mdgkb/mdgkb-server/handlers/valueTypes"
	"mdgkb/mdgkb-server/handlers/visitingRules"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
	authRouter "mdgkb/mdgkb-server/routing/auth"
	"mdgkb/mdgkb-server/routing/banners"
	"mdgkb/mdgkb-server/routing/buildings"
	commentsRouter "mdgkb/mdgkb-server/routing/comments"
	divisionsRouter "mdgkb/mdgkb-server/routing/divisions"
	doctorsRouter "mdgkb/mdgkb-server/routing/doctors"
	documentTypesRouter "mdgkb/mdgkb-server/routing/document-types"
	donorRulesRouter "mdgkb/mdgkb-server/routing/donorRules"
	"mdgkb/mdgkb-server/routing/educationalOraganization"
	eventsRouter "mdgkb/mdgkb-server/routing/events"
	faqRouter "mdgkb/mdgkb-server/routing/faqs"
	headsRouter "mdgkb/mdgkb-server/routing/heads"
	hospitalizationRouter "mdgkb/mdgkb-server/routing/hospitalization"
	menusRouter "mdgkb/mdgkb-server/routing/menus"
	newsRouter "mdgkb/mdgkb-server/routing/news"
	newsSlidesRouter "mdgkb/mdgkb-server/routing/newsSlides"
	"mdgkb/mdgkb-server/routing/normativeDocumentTypes"
	"mdgkb/mdgkb-server/routing/normativeDocuments"
	pagesRouter "mdgkb/mdgkb-server/routing/pages"
	paidProgramsRouter "mdgkb/mdgkb-server/routing/paidPrograms"
	paidProgramsGroupsRouter "mdgkb/mdgkb-server/routing/paidProgramsGroups"
	partnerTypesRouter "mdgkb/mdgkb-server/routing/partnerTypes"
	partnersRouter "mdgkb/mdgkb-server/routing/partners"
	preparationsRouter "mdgkb/mdgkb-server/routing/preparations"
	projectsRouter "mdgkb/mdgkb-server/routing/projects"
	questionsRouter "mdgkb/mdgkb-server/routing/questions"
	searchRouter "mdgkb/mdgkb-server/routing/search"
	"mdgkb/mdgkb-server/routing/sideOrganizations"
	"mdgkb/mdgkb-server/routing/tags"
	timetablePatternsRouter "mdgkb/mdgkb-server/routing/timetablePatterns"
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
)

func Init(r *gin.Engine, db *bun.DB, redisClient *redis.Client, config config.Config) {
	localUploader := helpers.NewLocalUploader(&config.UploadPath)
	localUploaderNew := uploadHelper.NewLocalUploader(&config.UploadPath)
	helper := helpers.NewHelper(config)
	r.Static("/static", "./static/")
	api := r.Group("/api/v1")

	authRouter.Init(api.Group("/auth"), auth.CreateHandler(db, helper))
	banners.Init(api.Group("/banners"), db, localUploader)
	buildings.Init(api.Group("/buildings"), db, localUploader)
	doctorsRouter.Init(api.Group("/doctors"), doctors.CreateHandler(db, helper))
	hospitalizationRouter.Init(api.Group("/hospitalizations"), db, helper)

	divisionsRouter.Init(api.Group("/divisions"), divisions.CreateHandler(db, helper))
	headsRouter.Init(api.Group("/heads"), heads.CreateHandler(db, helper))

	commentsRouter.Init(api.Group("/comments"), comments.CreateHandler(db))
	newsRouter.Init(api.Group("/news"), news.CreateHandler(db, helper))
	normativeDocumentTypes.Init(api.Group("/normative-document-types"), db, localUploader)
	normativeDocuments.Init(api.Group("/normative-documents"), db, localUploader)
	sideOrganizations.Init(api.Group("/side-organizations"), db, localUploader)
	tags.Init(api.Group("/tags"), db, localUploader)
	usersRouter.Init(api.Group("/users"), users.CreateHandler(db, helper))
	timetables.Init(api.Group("/timetables"), db)

	educationalOraganization.Init(api.Group("/educational-organization"), db, localUploaderNew)
	menusRouter.Init(api.Group("/menus"), menus.CreateHandler(db, helper))
	pagesRouter.Init(api.Group("/pages"), pages.CreateHandler(db, helper))
	projectsRouter.Init(api.Group("/projects"), projects.CreateHandler(db, helper))
	vacanciesRouter.Init(api.Group("/vacancies"), vacancies.CreateHandler(db, helper))
	vacancyResponseRouter.Init(api.Group("/vacancy-responses"), vacancyResponse.CreateHandler(db, helper))
	documentTypesRouter.Init(api.Group("/document-types"), documentTypes.CreateHandler(db))
	valueTypesRouter.Init(api.Group("/value-types"), valueTypes.CreateHandler(db))
	searchRouter.Init(api.Group("/search"), search.CreateHandler(db, helper))
	faqRouter.Init(api.Group("/faqs"), faqs.CreateHandler(db, helper))
	visitingRulesRouter.Init(api.Group("/visiting-rules"), visitingRules.CreateHandler(db, helper))
	newsSlidesRouter.Init(api.Group("/news-slides"), newsSlides.CreateHandler(db, helper))
	questionsRouter.Init(api.Group("/questions"), questions.CreateHandler(db, helper))
	eventsRouter.Init(api.Group("/events"), events.CreateHandler(db, helper))
	timetablePatternsRouter.Init(api.Group("/timetable-patterns"), timetablePatterns.CreateHandler(db, helper))
	paidProgramsRouter.Init(api.Group("/paid-programs"), paidPrograms.CreateHandler(db, helper))
	paidProgramsGroupsRouter.Init(api.Group("/paid-programs-groups"), paidProgramsGroups.CreateHandler(db, helper))
	partnerTypesRouter.Init(api.Group("/partner-types"), partnerTypes.CreateHandler(db, helper))
	partnersRouter.Init(api.Group("/partners"), partners.CreateHandler(db, helper))
	preparationsRouter.Init(api.Group("/preparations"), preparations.CreateHandler(db, helper))
	donorRulesRouter.Init(api.Group("/donor-rules"), donorRules.CreateHandler(db, helper))
}
