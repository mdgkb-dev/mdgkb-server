package routing

import (
	"mdgkb/mdgkb-server/config"
	"mdgkb/mdgkb-server/handlers/doctors"
	"mdgkb/mdgkb-server/handlers/documentTypes"
	"mdgkb/mdgkb-server/handlers/news"
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/handlers/vacancies"
	"mdgkb/mdgkb-server/handlers/vacancyResponse"
	"mdgkb/mdgkb-server/handlers/valueTypes"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
	"mdgkb/mdgkb-server/routing/auth"
	"mdgkb/mdgkb-server/routing/banners"
	"mdgkb/mdgkb-server/handlers/comments"
	"mdgkb/mdgkb-server/handlers/newsSlides"
	commentsRouter "mdgkb/mdgkb-server/routing/comments"
	"mdgkb/mdgkb-server/routing/buildings"
	"mdgkb/mdgkb-server/routing/divisions"
	hospitalizationRouter "mdgkb/mdgkb-server/routing/hospitalization"
	doctorsRouter "mdgkb/mdgkb-server/routing/doctors"
	documentTypesRouter "mdgkb/mdgkb-server/routing/document-types"
	"mdgkb/mdgkb-server/routing/educationalOraganization"
	"mdgkb/mdgkb-server/routing/menu"
	newsRouter "mdgkb/mdgkb-server/routing/news"
	"mdgkb/mdgkb-server/routing/normativeDocumentTypes"
	"mdgkb/mdgkb-server/routing/normativeDocuments"
	"mdgkb/mdgkb-server/routing/pages"
	"mdgkb/mdgkb-server/routing/sideOrganizations"
	"mdgkb/mdgkb-server/routing/tags"
	"mdgkb/mdgkb-server/routing/timetables"
	usersRouter "mdgkb/mdgkb-server/routing/users"
	vacanciesRouter "mdgkb/mdgkb-server/routing/vacancies"
	vacancyResponseRouter "mdgkb/mdgkb-server/routing/vacancyResponse"
	valueTypesRouter "mdgkb/mdgkb-server/routing/valueTypes"
	newsSlidesRouter "mdgkb/mdgkb-server/routing/newsSlides"

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

	auth.Init(api.Group("/auth"), db, redisClient)
	banners.Init(api.Group("/banners"), db, localUploader)
	buildings.Init(api.Group("/buildings"), db, localUploader)
	doctorsRouter.Init(api.Group("/doctors"), doctors.CreateHandler(db, localUploaderNew))
	hospitalizationRouter.Init(api.Group("/hospitalizations"), db, helper)

	divisions.Init(api.Group("/divisions"), db, localUploaderNew)

	commentsRouter.Init(api.Group("/comments"), comments.CreateHandler(db))
	newsRouter.Init(api.Group("/news"), news.CreateHandler(db, localUploaderNew))
	normativeDocumentTypes.Init(api.Group("/normative-document-types"), db, localUploader)
	normativeDocuments.Init(api.Group("/normative-documents"), db, localUploader)
	sideOrganizations.Init(api.Group("/side-organizations"), db, localUploader)
	tags.Init(api.Group("/tags"), db, localUploader)
	usersRouter.Init(api.Group("/users"), users.CreateHandler(db, localUploaderNew))
	timetables.Init(api.Group("/timetables"), db)

	educationalOraganization.Init(api.Group("/educational-organization"), db, localUploaderNew)
	menu.Init(api.Group("/menus"), db, localUploaderNew)
	pages.Init(api.Group("/pages"), db, localUploaderNew)
	vacanciesRouter.Init(api.Group("/vacancies"), vacancies.CreateHandler(db, localUploaderNew))
	vacancyResponseRouter.Init(api.Group("/vacancy-responses"), vacancyResponse.CreateHandler(db, helper))
	documentTypesRouter.Init(api.Group("/document-types"), documentTypes.CreateHandler(db))
	valueTypesRouter.Init(api.Group("/value-types"), valueTypes.CreateHandler(db))
	newsSlidesRouter.Init(api.Group("/news-slides"), newsSlides.CreateHandler(db, helper))
}
