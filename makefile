include .env

ifeq ($(OS),Windows_NT)
	database := .\database
	migrations := .\database\migrations
	cli := .\cmd\cli
	main := .\cmd\server\main.go
else
	database := database/*.go
	migrations := database/migrations/*.go
	cli := cmd/cli/*.go
	main := cmd/server/main.go
endif

run: migrate
	reflex -r '\.go' -s -- sh -c "go run cmd/server/main.go"

run_cold: migrate
	go run $(main)	

migrate_init:
	go run $(database) -action=init

migrate:
	go run $(database) -mode=migration -action=migrate

migrate_create:
	go run $(database) -mode=migration -action=create -name=${name}

seed:
	go run $(database) -mode=seed -action=migrate

seed_create:
	go run $(database) -mode=seed -action=create -name=${name}

migrate_rollback:
	go run $(migrations) rollback

drop_database:
	go run $(database) -action=dropDatabase

dump_from_remote:
	 @./database/dump_pg.sh $(DB_PASSWORD)

dump: dump_from_remote migrate

deploy:
	./cmd/server/deploy.sh DEPLOY_PATH=$(DEPLOY_PATH) DEPLOY_BRANCH=$(DEPLOY_BRANCH)

kill:
	kill -9 `lsof -t -i:$(SERVER_PORT)`


#####
#GIT#
#####

git_push: git_commit
	git push -u origin HEAD

git_commit:
	git pull origin develop
	git add .
	git commit -m "$m"

git_merge: git_push
	git checkout develop
	git pull
	git merge @{-1}
	git push

# example: make git_feature n=1
git_feature:
	git flow feature start PORTAL-$n