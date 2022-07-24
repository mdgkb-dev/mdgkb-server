package migrations

import "github.com/uptrace/bun/migrate"

func Init() *migrate.Migrations {
	var migrations = migrate.NewMigrations()
	if err := migrations.DiscoverCaller(); err != nil {
		panic(err)
	}
	return migrations
}
