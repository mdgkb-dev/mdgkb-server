package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"log"
	"mdgkb/mdgkb-server/config"
	"mdgkb/mdgkb-server/database/connect"
	"mdgkb/mdgkb-server/routing"
	"net/http"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	router := gin.Default()

	db := connect.InitDB(conf)
	defer db.Close()
	redis := connect.InitRedis(conf)
	routing.Init(router, db, redis, *conf)

	err = http.ListenAndServe(fmt.Sprintf(":%s", conf.ServerPort), router)
	if err != nil {
		panic(err)
	}
}
