package connect

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"

	"mdgkb/mdgkb-server/config"
)

func InitDB(conf *config.Config) *bun.DB {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", conf.DbDb, conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)
	fmt.Println("postgres://postgres:1@localhost:5432/mdgkb?sslmode=disable")
	fmt.Println(dsn)
	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(conn, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose()))
	_, _ = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	recognizeModels(db)
	return db
}

func recognizeModels(db *bun.DB) {
	db.RegisterModel((*models.NewsToCategory)(nil))
	db.RegisterModel((*models.NewsToTag)(nil))
}

var client *redis.Client

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
