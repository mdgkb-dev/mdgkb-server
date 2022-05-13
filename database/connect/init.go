package connect

import (
	"database/sql"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/uptrace/bun/extra/bundebug"
	"mdgkb/mdgkb-server/models"

	"github.com/go-redis/redis/v7"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/pro-assistance/pro-assister/config"
)

func InitDB(conf config.DB) *bun.DB {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", conf.DB, conf.User, conf.Password, conf.Host, conf.Port, conf.Name)
	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(conn, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(false),
	))
	_, _ = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	//recognizeModels(db)
	return db
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
		panic(err)
	}
	return client
}
