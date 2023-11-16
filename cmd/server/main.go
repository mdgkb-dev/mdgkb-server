package main

import (
	"log"
	"mdgkb/mdgkb-server/loggerhelper"
	"mdgkb/mdgkb-server/migrations"
	"mdgkb/mdgkb-server/routing"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/oiime/logrusbun"
	"github.com/pro-assistance/pro-assister/config"
	"github.com/pro-assistance/pro-assister/cronHelper"
	helperPack "github.com/pro-assistance/pro-assister/helper"
	"github.com/sirupsen/logrus"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	gin.SetMode(gin.ReleaseMode)
	logger := loggerhelper.NewLogger()
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(loggerhelper.LoggingMiddleware(logger))
	helper := helperPack.NewHelper(*conf)

	routing.Init(router, helper)
	helper.DB.DB.AddQueryHook(logrusbun.NewQueryHook(logrusbun.QueryHookOptions{Logger: logger, ErrorLevel: logrus.ErrorLevel, QueryLevel: logrus.DebugLevel}))
	updateJob := &cronHelper.Job{Schedule: "*/1 * * * *", Function: updateSearchElementsTable(helper.DB.DB)}
	err = helper.Cron.AddJobs(cronHelper.Jobs{updateJob})
	if err != nil {
		log.Fatal("cannot add cron jobs:", err)
	}
	helper.Run(migrations.Init(), router)
}
