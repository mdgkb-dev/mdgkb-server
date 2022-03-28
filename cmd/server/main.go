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

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	router := gin.Default()
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
