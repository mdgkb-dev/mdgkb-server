package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mdgkb/mdgkb-server/config"
	"mdgkb/mdgkb-server/database/connect"
	"mdgkb/mdgkb-server/database/migrations"
	"mdgkb/mdgkb-server/database/seeding"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func main() {
	mode := flag.String("mode", "migration", "init/create")
	action := flag.String("action", "", "init/create/createSql/run/rollback")
	name := flag.String("name", "", "init/create/createSql/run/rollback")
	flag.Parse()

	db := getDb()
	defer db.Close()
	var migrator *migrate.Migrator
	if *mode == "migration" {
		migrator = migrate.NewMigrator(db, migrations.Migrations)
	} else {
		migrator = migrate.NewMigrator(db, seeding.Migrations)
	}
	doAction(migrator, action, name)
}

func doAction(migrator *migrate.Migrator, action *string, name *string) {
	switch *action {
	case "init":
		initMigration(migrator)
	case "dropDatabase":
		dropDatabase(migrator)
	case "create":
		createMigrationSql(migrator, name)
	case "migrate":
		runMigration(migrator)
	case "status":
		ms, err := migrator.MigrationsWithStatus(context.TODO())
		if err != nil {
			panic(err)
		}
		fmt.Printf("migrations: %s\n", ms)
		fmt.Printf("unapplied migrations: %s\n", ms.Unapplied())
		fmt.Printf("last migration group: %s\n", ms.LastGroup())
	default:
		log.Fatal("cannot parse action")
	}
}

func getDb() *bun.DB {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	return connect.InitDB(conf)
}
