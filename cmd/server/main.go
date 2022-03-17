package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/pro-assistance/pro-assister/config"
	"log"
	"mdgkb/mdgkb-server/database/connect"
	"mdgkb/mdgkb-server/routing"
	"net/http"

	_ "github.com/go-pg/pg/v10/orm"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	router := gin.Default()
	router.Use(CORSMiddleware())

	db := connect.InitDB(conf.DB)
	defer db.Close()
	redis := connect.InitRedis(conf)
	elasticSearch := connect.InitElasticSearch(conf)
	routing.Init(router, db, redis, elasticSearch, *conf)

	err = http.ListenAndServe(fmt.Sprintf(":%s", conf.ServerPort), router)
	if err != nil {
		panic(err)
	}
}
