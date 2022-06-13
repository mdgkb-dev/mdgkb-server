package connect

import (
	"database/sql"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v7"
	"github.com/oiime/logrusbun"
	"github.com/pro-assistance/pro-assister/config"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
	"mdgkb/mdgkb-server/models"
	"os"
	"time"
)

func InitDB(conf config.DB) *bun.DB {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", conf.DB, conf.User, conf.Password, conf.Host, conf.Port, conf.Name)
	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(conn, sqlitedialect.New())

	_, _ = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	createLogger(conf, db)
	//recognizeModels(db)
	return db
}

func createLogger(conf config.DB, db *bun.DB) {
	log := logrus.New()
	log.Level = logrus.TraceLevel
	if conf.LogPath != "" {
		f, err := os.OpenFile(conf.LogPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			panic(err)
		}
		log.SetOutput(f)
	} else {
		log.SetOutput(os.Stdout)
	}
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		CallerPrettyfier:  nil,
		PrettyPrint:       false,
	})
	//defer f.Close()

	//bun.SetLogger(log)
	db.AddQueryHook(logrusbun.NewQueryHook(logrusbun.QueryHookOptions{
		LogSlow:         time.Second,
		Logger:          log,
		QueryLevel:      logrus.DebugLevel,
		ErrorLevel:      logrus.ErrorLevel,
		SlowLevel:       logrus.WarnLevel,
		MessageTemplate: "{{.Operation}} : {{.Query}}",
		ErrorTemplate:   "{{.Operation}}[{{.Duration}}]: {{.Query}}: {{.Error}}",
	}))
}

func recognizeModels(db *bun.DB) {
	db.RegisterModel((*models.NewsToCategory)(nil))
	db.RegisterModel((*models.NewsToTag)(nil))
}

func InitElasticSearch(conf *config.Config) (client *elasticsearch.Client) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			conf.ElasticSearch.ElasticSearchURL,
		},
	}
	if conf.ElasticSearch.ElasticSearchOn {
		client, err := elasticsearch.NewClient(cfg)
		if err != nil {
			panic(err)
		}
		return client
	}
	return nil
}

func InitRedis(conf *config.Config) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", conf.RedisHost, conf.RedisPort), //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Println(err)
	}
	return client
}
