package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/pro-assistance/pro-assister/config"
	helperPack "github.com/pro-assistance/pro-assister/helper"
	"log"
	"mdgkb/mdgkb-server/database/connect"
	"mdgkb/mdgkb-server/routing"
	"net/http"
	"os"
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
	helper := helperPack.NewHelper(*conf)
	ex, _ := os.Executable()
	err = helper.Util.WritePidFile(ex + ".pid")
	if err != nil {
		log.Fatal("cannot write pid:", err)
	}
	routing.Init(router, db, redis, elasticSearch, helper)
	err = http.ListenAndServe(fmt.Sprintf(":%s", conf.ServerPort), router)
	if err != nil {
		panic(err)
	}
}
