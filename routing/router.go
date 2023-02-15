package routing

import (
	"mdgkb/mdgkb-server/handlers/appointmenstypes"
	"mdgkb/mdgkb-server/handlers/appointments"
	"mdgkb/mdgkb-server/handlers/auth"
	"mdgkb/mdgkb-server/handlers/banners"
	"mdgkb/mdgkb-server/handlers/buildings"
	"mdgkb/mdgkb-server/handlers/callbackrequests"
	"mdgkb/mdgkb-server/handlers/candidateapplications"
	"mdgkb/mdgkb-server/handlers/candidateexams"
	"mdgkb/mdgkb-server/handlers/certificates"
	"mdgkb/mdgkb-server/handlers/children"
	"mdgkb/mdgkb-server/handlers/comments"
	"mdgkb/mdgkb-server/handlers/dailymenuitems"
	"mdgkb/mdgkb-server/handlers/dailymenuorders"
	"mdgkb/mdgkb-server/handlers/dailymenus"
	"mdgkb/mdgkb-server/handlers/diets"
	"mdgkb/mdgkb-server/handlers/dietsgroups"
	"mdgkb/mdgkb-server/handlers/dishesgroups"
	"mdgkb/mdgkb-server/handlers/dishessamples"
	"mdgkb/mdgkb-server/handlers/divisions"
	"mdgkb/mdgkb-server/handlers/doctors"
	"mdgkb/mdgkb-server/handlers/donorrules"
	"mdgkb/mdgkb-server/handlers/dpoapplications"
	"mdgkb/mdgkb-server/handlers/dpocourses"
	"mdgkb/mdgkb-server/handlers/educationalacademics"
	"mdgkb/mdgkb-server/handlers/educationalmanagers"
	"mdgkb/mdgkb-server/handlers/educationyears"
	"mdgkb/mdgkb-server/handlers/employees"
	"mdgkb/mdgkb-server/handlers/entrances"
	"mdgkb/mdgkb-server/handlers/events"
	"mdgkb/mdgkb-server/handlers/faqs"
	"mdgkb/mdgkb-server/handlers/formpatterns"
	"mdgkb/mdgkb-server/handlers/formstatuses"
	"mdgkb/mdgkb-server/handlers/formstatusgroups"
	"mdgkb/mdgkb-server/handlers/formvalues"
	"mdgkb/mdgkb-server/handlers/gates"
	"mdgkb/mdgkb-server/handlers/heads"
	"mdgkb/mdgkb-server/handlers/hospitalizations"
	"mdgkb/mdgkb-server/handlers/hospitalizationstypes"
	"mdgkb/mdgkb-server/handlers/medicalprofiles"
	"mdgkb/mdgkb-server/handlers/menus"
	"mdgkb/mdgkb-server/handlers/meta"
	"mdgkb/mdgkb-server/handlers/news"
	"mdgkb/mdgkb-server/handlers/newsslides"
	"mdgkb/mdgkb-server/handlers/pages"
	"mdgkb/mdgkb-server/handlers/pagesections"
	"mdgkb/mdgkb-server/handlers/pagesidemenus"
	"mdgkb/mdgkb-server/handlers/paidprograms"
	"mdgkb/mdgkb-server/handlers/paidprogramsgroups"
	"mdgkb/mdgkb-server/handlers/paidservices"
	"mdgkb/mdgkb-server/handlers/partners"
	"mdgkb/mdgkb-server/handlers/partnertypes"
	"mdgkb/mdgkb-server/handlers/pointsachievements"
	"mdgkb/mdgkb-server/handlers/postgraduateapplications"
	"mdgkb/mdgkb-server/handlers/postgraduatecourses"
	"mdgkb/mdgkb-server/handlers/preparations"
	"mdgkb/mdgkb-server/handlers/projects"
	"mdgkb/mdgkb-server/handlers/questions"
	"mdgkb/mdgkb-server/handlers/residencyapplications"
	"mdgkb/mdgkb-server/handlers/residencycourses"
	"mdgkb/mdgkb-server/handlers/roles"
	"mdgkb/mdgkb-server/handlers/search"
	"mdgkb/mdgkb-server/handlers/sideorganizations"
	"mdgkb/mdgkb-server/handlers/specializations"
	"mdgkb/mdgkb-server/handlers/supportmessages"
	"mdgkb/mdgkb-server/handlers/tags"
	"mdgkb/mdgkb-server/handlers/teachers"
	"mdgkb/mdgkb-server/handlers/timetablepatterns"
	"mdgkb/mdgkb-server/handlers/timetables"
	"mdgkb/mdgkb-server/handlers/treatdirections"
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/handlers/vacancies"
	"mdgkb/mdgkb-server/handlers/vacancyresponse"
	"mdgkb/mdgkb-server/handlers/valuetypes"
	"mdgkb/mdgkb-server/handlers/visitingrules"
	"mdgkb/mdgkb-server/handlers/visitsapplications"
	appointmentsRouter "mdgkb/mdgkb-server/routing/appointments"
	appointmentsTypesRouter "mdgkb/mdgkb-server/routing/appointmentstypes"
	authRouter "mdgkb/mdgkb-server/routing/auth"
	bannersRouter "mdgkb/mdgkb-server/routing/banners"
	buildingsRouter "mdgkb/mdgkb-server/routing/buildings"
	callbackRequestsRouter "mdgkb/mdgkb-server/routing/callbackrequests"
	candidateApplicationsRouter "mdgkb/mdgkb-server/routing/candidateapplications"
	candidateExamsRouter "mdgkb/mdgkb-server/routing/candidateexams"
	certificatesRouter "mdgkb/mdgkb-server/routing/certificates"
	childrenRouter "mdgkb/mdgkb-server/routing/children"
	commentsRouter "mdgkb/mdgkb-server/routing/comments"
	dailyMenuItemsRouter "mdgkb/mdgkb-server/routing/dailymenuitems"
	dailyMenuOrdersRouter "mdgkb/mdgkb-server/routing/dailymenuorders"
	dailyMenusRouter "mdgkb/mdgkb-server/routing/dailymenus"
	dietsRouter "mdgkb/mdgkb-server/routing/diets"
	dietsGroupsRouter "mdgkb/mdgkb-server/routing/dietsgroups"
	dishesGroupsRouter "mdgkb/mdgkb-server/routing/dishesgroups"
	dishesSamplesRouter "mdgkb/mdgkb-server/routing/dishessamples"
	divisionsRouter "mdgkb/mdgkb-server/routing/divisions"
	doctorsRouter "mdgkb/mdgkb-server/routing/doctors"
	donorRulesRouter "mdgkb/mdgkb-server/routing/donorrules"
	dpoApplicationsRouter "mdgkb/mdgkb-server/routing/dpoapplications"
	dpoCoursesRouter "mdgkb/mdgkb-server/routing/dpocourses"
	educationalAcademicsRouter "mdgkb/mdgkb-server/routing/educationalacademics"
	educationalManagersRouter "mdgkb/mdgkb-server/routing/educationalmanagers"
	educationYearsRouter "mdgkb/mdgkb-server/routing/educationyears"
	employeesRouter "mdgkb/mdgkb-server/routing/employees"
	entrancesRouter "mdgkb/mdgkb-server/routing/entrances"
	eventsRouter "mdgkb/mdgkb-server/routing/events"
	faqRouter "mdgkb/mdgkb-server/routing/faqs"
	formPatternsRouter "mdgkb/mdgkb-server/routing/formpatterns"
	formStatusesRouter "mdgkb/mdgkb-server/routing/formstatuses"
	formStatusGroupsRouter "mdgkb/mdgkb-server/routing/formstatusgroups"
	formValuesRouter "mdgkb/mdgkb-server/routing/formvalues"
	gatesRouter "mdgkb/mdgkb-server/routing/gates"
	headsRouter "mdgkb/mdgkb-server/routing/heads"
	hospitalizationRouter "mdgkb/mdgkb-server/routing/hospitalizations"
	hospitalizationsTypesRouter "mdgkb/mdgkb-server/routing/hospitalizationstypes"
	medicalProfilesRouter "mdgkb/mdgkb-server/routing/medicalprofiles"
	menusRouter "mdgkb/mdgkb-server/routing/menus"
	metaRouter "mdgkb/mdgkb-server/routing/meta"
	newsRouter "mdgkb/mdgkb-server/routing/news"
	newsSlidesRouter "mdgkb/mdgkb-server/routing/newsslides"
	pagesRouter "mdgkb/mdgkb-server/routing/pages"
	pageSectionsRouter "mdgkb/mdgkb-server/routing/pagesections"
	pageSideMenusRouter "mdgkb/mdgkb-server/routing/pagesidemenus"
	paidProgramsRouter "mdgkb/mdgkb-server/routing/paidprograms"
	paidProgramsGroupsRouter "mdgkb/mdgkb-server/routing/paidprogramsgroups"
	paidServicesRouter "mdgkb/mdgkb-server/routing/paidservices"
	partnersRouter "mdgkb/mdgkb-server/routing/partners"
	partnerTypesRouter "mdgkb/mdgkb-server/routing/partnertypes"
	pointsAchievementsRouter "mdgkb/mdgkb-server/routing/pointsachievements"
	postgraduateApplicationsRouter "mdgkb/mdgkb-server/routing/postgraduateapplications"
	postgraduateCoursesRouter "mdgkb/mdgkb-server/routing/postgraduatecourses"
	preparationsRouter "mdgkb/mdgkb-server/routing/preparations"
	projectsRouter "mdgkb/mdgkb-server/routing/projects"
	questionsRouter "mdgkb/mdgkb-server/routing/questions"
	residencyApplicationsRouter "mdgkb/mdgkb-server/routing/residencyapplications"
	residencyCoursesRouter "mdgkb/mdgkb-server/routing/residencycourses"
	rolesRouter "mdgkb/mdgkb-server/routing/roles"
	searchRouter "mdgkb/mdgkb-server/routing/search"
	sideOrganizationsRouter "mdgkb/mdgkb-server/routing/sideorganizations"
	specializationsRouter "mdgkb/mdgkb-server/routing/specializations"
	supportMessagesRouter "mdgkb/mdgkb-server/routing/supportmessages"
	tagsRouter "mdgkb/mdgkb-server/routing/tags"
	teachersRouter "mdgkb/mdgkb-server/routing/teachers"
	timetablePatternsRouter "mdgkb/mdgkb-server/routing/timetablepatterns"
	timetablesRouter "mdgkb/mdgkb-server/routing/timetables"
	treatDirectionsRouter "mdgkb/mdgkb-server/routing/treatdirections"
	usersRouter "mdgkb/mdgkb-server/routing/users"
	vacanciesRouter "mdgkb/mdgkb-server/routing/vacancies"
	vacancyResponseRouter "mdgkb/mdgkb-server/routing/vacancyresponse"
	valueTypesRouter "mdgkb/mdgkb-server/routing/valuetypes"
	visitingRulesRouter "mdgkb/mdgkb-server/routing/visitingrules"
	visitsApplicationsRouter "mdgkb/mdgkb-server/routing/visitsapplications"

	"github.com/gin-gonic/gin"

	helperPack "github.com/pro-assistance/pro-assister/helper"
)

func Init(r *gin.Engine, helper *helperPack.Helper) {
	//m := middleware.CreateMiddleware(helper)

	//r.Use(m.CORSMiddleware())
	//r.Use(m.CheckPermission())
	r.Use(gin.Logger())

	r.Static("/api/v1/static", "./static/")
	authGroup := r.Group("/api/v1/auth")
	authRouter.Init(authGroup.Group(""), auth.CreateHandler(helper))

	api := r.Group("/api/v1")
	ws := r.Group("/ws")

	//api.Use(m.Authentication())
	//api.Use(m.CORSMiddleware())
	//ws.Use(m.CORSMiddleware())
	api.GET("/subscribe/:channel", helper.Broker.ServeHTTP)

	bannersRouter.Init(api.Group("/banners"), banners.CreateHandler(helper))
	buildingsRouter.Init(api.Group("/buildings"), buildings.CreateHandler(helper))
	doctorsRouter.Init(api.Group("/doctors"), doctors.CreateHandler(helper))
	hospitalizationRouter.Init(api.Group("/hospitalizations"), hospitalizations.CreateHandler(helper))
	hospitalizationsTypesRouter.Init(api.Group("/hospitalizations-types"), hospitalizationstypes.CreateHandler(helper))
	divisionsRouter.Init(api.Group("/divisions"), divisions.CreateHandler(helper))
	headsRouter.Init(api.Group("/heads"), heads.CreateHandler(helper))
	commentsRouter.Init(api.Group("/comments"), comments.CreateHandler(helper))
	newsRouter.Init(api.Group("/news"), news.CreateHandler(helper))
	sideOrganizationsRouter.Init(api.Group("/side-organizations"), sideorganizations.CreateHandler(helper))
	tagsRouter.Init(api.Group("/tags"), tags.CreateHandler(helper))
	usersRouter.Init(api.Group("/users"), users.CreateHandler(helper))
	timetablesRouter.Init(api.Group("/timetables"), timetables.CreateHandler(helper))
	menusRouter.Init(api.Group("/menus"), menus.CreateHandler(helper))
	pagesRouter.Init(api.Group("/pages"), pages.CreateHandler(helper))
	projectsRouter.Init(api.Group("/projects"), projects.CreateHandler(helper))
	entrancesRouter.Init(api.Group("/entrances"), entrances.CreateHandler(helper))
	vacanciesRouter.Init(api.Group("/vacancies"), vacancies.CreateHandler(helper))
	vacancyResponseRouter.Init(api.Group("/vacancy-responses"), vacancyresponse.CreateHandler(helper))
	pageSectionsRouter.Init(api.Group("/page-sections"), pagesections.CreateHandler(helper))
	valueTypesRouter.Init(api.Group("/value-types"), valuetypes.CreateHandler(helper))
	searchRouter.Init(api.Group("/search"), search.CreateHandler(helper))
	faqRouter.Init(api.Group("/faqs"), faqs.CreateHandler(helper))
	visitingRulesRouter.Init(api.Group("/visiting-rules"), visitingrules.CreateHandler(helper))
	newsSlidesRouter.Init(api.Group("/news-slides"), newsslides.CreateHandler(helper))
	questionsRouter.Init(api.Group("/questions"), questions.CreateHandler(helper))
	eventsRouter.Init(api.Group("/events"), events.CreateHandler(helper))
	timetablePatternsRouter.Init(api.Group("/timetable-patterns"), timetablepatterns.CreateHandler(helper))
	formPatternsRouter.Init(api.Group("/form-patterns"), formpatterns.CreateHandler(helper))
	paidProgramsRouter.Init(api.Group("/paid-programs"), paidprograms.CreateHandler(helper))
	paidProgramsGroupsRouter.Init(api.Group("/paid-programs-groups"), paidprogramsgroups.CreateHandler(helper))
	partnerTypesRouter.Init(api.Group("/partner-types"), partnertypes.CreateHandler(helper))
	pageSideMenusRouter.Init(api.Group("/page-side-menus"), pagesidemenus.CreateHandler(helper))
	partnersRouter.Init(api.Group("/partners"), partners.CreateHandler(helper))
	preparationsRouter.Init(api.Group("/preparations"), preparations.CreateHandler(helper))
	donorRulesRouter.Init(api.Group("/donor-rules"), donorrules.CreateHandler(helper))
	certificatesRouter.Init(api.Group("/certificates"), certificates.CreateHandler(helper))
	metaRouter.Init(api.Group("/meta"), meta.CreateHandler(helper))
	paidServicesRouter.Init(api.Group("/paid-services"), paidservices.CreateHandler(helper))
	medicalProfilesRouter.Init(api.Group("/medical-profiles"), medicalprofiles.CreateHandler(helper))
	treatDirectionsRouter.Init(api.Group("/treat-directions"), treatdirections.CreateHandler(helper))
	callbackRequestsRouter.Init(api.Group("/callback-requests"), callbackrequests.CreateHandler(helper))
	visitsApplicationsRouter.Init(api.Group("/visits-applications"), visitsapplications.CreateHandler(helper))
	dpoCoursesRouter.Init(api.Group("/dpo-courses"), dpocourses.CreateHandler(helper))
	postgraduateCoursesRouter.Init(api.Group("/postgraduate-courses"), postgraduatecourses.CreateHandler(helper))
	dpoApplicationsRouter.Init(api.Group("/dpo-applications"), dpoapplications.CreateHandler(helper))
	educationalAcademicsRouter.Init(api.Group("/educational-academics"), educationalacademics.CreateHandler(helper))
	residencyApplicationsRouter.Init(api.Group("/residency-applications"), residencyapplications.CreateHandler(helper))
	formValuesRouter.Init(api.Group("/form-values"), formvalues.CreateHandler(helper))
	formStatusesRouter.Init(api.Group("/form-statuses"), formstatuses.CreateHandler(helper))
	formStatusGroupsRouter.Init(api.Group("/form-status-groups"), formstatusgroups.CreateHandler(helper))
	postgraduateApplicationsRouter.Init(api.Group("/postgraduate-applications"), postgraduateapplications.CreateHandler(helper))
	teachersRouter.Init(api.Group("/teachers"), teachers.CreateHandler(helper))
	educationalManagersRouter.Init(api.Group("/educational-managers"), educationalmanagers.CreateHandler(helper))
	appointmentsRouter.Init(api.Group("/appointments"), appointments.CreateHandler(helper))
	childrenRouter.Init(api.Group("/children"), children.CreateHandler(helper))
	gatesRouter.Init(api.Group("/gates"), gates.CreateHandler(helper))
	specializationsRouter.Init(api.Group("/specializations"), specializations.CreateHandler(helper))
	candidateApplicationsRouter.Init(api.Group("/candidate-applications"), candidateapplications.CreateHandler(helper))
	candidateExamsRouter.Init(api.Group("/candidate-exams"), candidateexams.CreateHandler(helper))
	rolesRouter.Init(api.Group("/roles"), roles.CreateHandler(helper))
	residencyCoursesRouter.Init(api.Group("/residency-courses"), residencycourses.CreateHandler(helper))
	educationYearsRouter.Init(api.Group("/education-years"), educationyears.CreateHandler(helper))
	pointsAchievementsRouter.Init(api.Group("/points-achievements"), pointsachievements.CreateHandler(helper))
	dietsRouter.Init(api.Group("/diets"), diets.CreateHandler(helper))
	dietsGroupsRouter.Init(api.Group("/diets-groups"), dietsgroups.CreateHandler(helper))
	employeesRouter.Init(api.Group("/employees"), employees.CreateHandler(helper))
	dishesGroupsRouter.Init(api.Group("/dishes-groups"), dishesgroups.CreateHandler(helper))
	dishesSamplesRouter.Init(api.Group("/dishes-samples"), dishessamples.CreateHandler(helper))

	dailyMenusHandler := dailymenus.CreateHandler(helper)
	dailyMenusRouter.Init(api.Group("/daily-menus"), dailyMenusHandler)
	ws.Group("/daily-menus").GET("/regular-update", dailyMenusHandler.GetWeb)

	dailyMenuItemsRouter.Init(api.Group("/daily-menu-items"), dailymenuitems.CreateHandler(helper))
	supportMessagesRouter.Init(api.Group("/support-messages"), supportmessages.CreateHandler(helper))
	appointmentsTypesRouter.Init(api.Group("/appointments-types"), appointmenstypes.CreateHandler(helper))
	dailyMenuOrdersRouter.Init(api.Group("/daily-menu-orders"), dailymenuorders.CreateHandler(helper))
}
