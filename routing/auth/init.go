package auth

import (
	handler "mdgkb/mdgkb-server/handlers/auth"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/uptrace/bun"

	_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB, redisClient *redis.Client) {
	var h = handler.NewHandler(handler.NewRepository(db), redisClient)
	r.POST("/login", h.Login)
	r.POST("/register", h.Register)
	r.POST("/refresh", h.Refresh)
	r.POST("/logout", h.Logout)
	//r.POST("/check-email", handler.CheckEmail)
	//r.GET("/logout", handler.Logout)
}
