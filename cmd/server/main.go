package main

import (
	"log"
	"mdgkb/mdgkb-server/migrations"
	"mdgkb/mdgkb-server/routing"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/pro-assistance/pro-assister/config"
	"github.com/pro-assistance/pro-assister/cronHelper"
	helperPack "github.com/pro-assistance/pro-assister/helper"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	gin.SetMode(gin.ReleaseMode)
	helper := helperPack.NewHelper(*conf)

	updateJob := &cronHelper.Job{Schedule: "*/1 * * * *", Function: updateSearchElementsTable(helper.DB.DB)}
	err = helper.Cron.AddJobs(cronHelper.Jobs{updateJob})
	if err != nil {
		log.Fatal("cannot add cron jobs:", err)
	}
	helper.Run(migrations.Init(), routing.Init)
}
