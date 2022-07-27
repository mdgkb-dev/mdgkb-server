package main

import (
	"log"
	"mdgkb/mdgkb-server/migrations"
	"mdgkb/mdgkb-server/routing"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/pro-assistance/pro-assister/config"
	helperPack "github.com/pro-assistance/pro-assister/helper"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	router := gin.Default()
	helper := helperPack.NewHelper(*conf)

	routing.Init(router, helper)
	helper.Run(migrations.Init(), router)
}
