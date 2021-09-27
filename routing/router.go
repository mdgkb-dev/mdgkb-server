package routing

import (
	"mdgkb/mdgkb-server/config"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
	"mdgkb/mdgkb-server/routing/auth"
	"mdgkb/mdgkb-server/routing/banners"
	"mdgkb/mdgkb-server/routing/buildings"
	"mdgkb/mdgkb-server/routing/carousels"
	"mdgkb/mdgkb-server/routing/divisions"
	"mdgkb/mdgkb-server/routing/doctors"
	"mdgkb/mdgkb-server/routing/educationalOraganization"
	"mdgkb/mdgkb-server/routing/news"
	"mdgkb/mdgkb-server/routing/normativeDocumentTypes"
	"mdgkb/mdgkb-server/routing/normativeDocuments"
	"mdgkb/mdgkb-server/routing/sideOrganizations"
	"mdgkb/mdgkb-server/routing/tags"
	"mdgkb/mdgkb-server/routing/timetables"
	"mdgkb/mdgkb-server/routing/users"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/go-redis/redis/v7"
	"github.com/uptrace/bun"
)

func Init(r *gin.Engine, db *bun.DB, redisClient *redis.Client, config config.Config) {
	localUploader := helpers.NewLocalUploader(&config.UploadPath)
	r.Static("/static", "./static/")
	api := r.Group("/api/v1")

	auth.Init(api.Group("/auth"), db, redisClient)
	banners.Init(api.Group("/banners"), db, localUploader)
	buildings.Init(api.Group("/buildings"), db, localUploader)
	carousels.Init(api.Group("/carousels"), db, localUploader)
	doctors.Init(api.Group("/doctors"), db, localUploader)
	divisions.Init(api.Group("/divisions"), db, localUploader)
	news.Init(api.Group("/news"), db, localUploader)
	normativeDocumentTypes.Init(api.Group("/normative-document-types"), db, localUploader)
	normativeDocuments.Init(api.Group("/normative-documents"), db, localUploader)
	sideOrganizations.Init(api.Group("/side-organizations"), db, localUploader)
	tags.Init(api.Group("/tags"), db, localUploader)
	users.Init(api.Group("/users"), db, localUploader)
	timetables.Init(api.Group("/timetables"), db, localUploader)

	educationalOraganization.Init(api.Group("/educational-organization"), db, uploadHelper.NewLocalUploader(&config.UploadPath))
}
