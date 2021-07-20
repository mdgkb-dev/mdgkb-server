run:
	reflex -r '\.go' -s -- sh -c "go run cmd/server/main.go"

migrate_init:
	go run database/*.go -action=init

migrate:
	go run database/*.go -mode=migration -action=migrate

migrate_create:
	go run database/*.go -mode=migration -action=create -name=${name}

seed:
	go run database/*.go -mode=seed -action=migrate

seed_create:
	go run database/*.go -mode=seed -action=create -name=${name}

migrate_rollback:
	go run database/migrations/*.go rollback



