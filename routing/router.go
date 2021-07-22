package routing

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/go-redis/redis/v7"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/config"
	"mdgkb/mdgkb-server/helpers"
	"mdgkb/mdgkb-server/routing/auth"
	"mdgkb/mdgkb-server/routing/buildings"
	"mdgkb/mdgkb-server/routing/divisions"
	"mdgkb/mdgkb-server/routing/news"
	"mdgkb/mdgkb-server/routing/users"
)

func Init(r *gin.Engine, db *bun.DB, redisClient *redis.Client, config config.Config) {
	localUploader := helpers.NewLocalUploader(&config.UploadPath)
	r.Static("/static", "./static")
	api := r.Group("/api/v1")
	auth.Init(api.Group("/auth"), db, redisClient)
	news.Init(api.Group("/news"), db, localUploader)
	buildings.Init(api.Group("/buildings"), db, localUploader)
	divisions.Init(api.Group("/divisions"), db, localUploader)
	users.Init(api.Group("/users"), db, localUploader)
}
