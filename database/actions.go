package main

import (
	"context"
	"fmt"
	"github.com/uptrace/bun/migrate"
	"log"
)

func createMigrationSql(migrator *migrate.Migrator, name *string) {
	mf, err := migrator.CreateSQL(context.TODO(), *name)
	if err != nil {
		panic(err)
	}
	fmt.Printf("created migration %s (%s)\n", mf.FileName, mf.FilePath)
}

func initMigration(migrator *migrate.Migrator) {
	err := migrator.Init(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	_, err = migrator.DB().Exec("create sequence bun_migration_locks_id_seq;")
	_, err = migrator.DB().Exec("create sequence bun_migrations_id_seq;")
	_, err = migrator.DB().Exec("alter table bun_migration_locks alter column id set default nextval('public.bun_migration_locks_id_seq');")
	_, err = migrator.DB().Exec("alter table bun_migrations alter column id set default nextval('public.bun_migrations_id_seq');")
	_, err = migrator.DB().Exec("alter sequence bun_migration_locks_id_seq owned by bun_migration_locks.id;")
	_, err = migrator.DB().Exec("alter sequence bun_migrations_id_seq owned by bun_migrations.id;")

	if err != nil {
		fmt.Println(err)
	}
}

func runMigration(migrator *migrate.Migrator) {
	group, err := migrator.Migrate(context.TODO())
	if err != nil {
		log.Fatalf("fail migrate: %s", err)
	}

	if group == nil || group.ID == 0 {
		fmt.Printf("there are no new migrations to run\n")
		return
	}

	fmt.Printf("migrated to %s\n", group)
}
