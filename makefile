include .env

ifeq ($(OS),Windows_NT)
	migrations := .\migrations
	cli := .\cmd\cli
	main := .\cmd\server\main.go
else
	migrations := database/*.go
	cli := cmd/cli/*.go
	main := cmd/server/main.go
endif

run: migrate set_git_hooks_dir
	reflex -r '\.go' -s -- sh -c "go run $(main)"

set_git_hooks_dir:
	git config core.hooksPath	 cmd/githooks/

run_cold:
	go run $(main)

migrate:
	go run $(main) -mode=migrate -action=migrate

migrate_create:
	go run $(main) -mode=migrate -action=create -name=${name}

migrate_rollback:
	go run $(migrates) rollback

drop_database:
	go run $(database) -action=dropDatabase

dump_from_remote:
	@./cmd/dump_pg.sh $(DB_NAME) $(DB_USER) $(DB_PASSWORD) $(DB_REMOTE_USER) $(DB_REMOTE_PASSWORD)

dump: dump_from_remote

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

git_deploy:
	git checkout develop
	git pull
	git checkout master
	git merge --no-commit develop
	git push

#######
#TESTS#
#######

test: test_dumb make_test drop_test_database

make_test:
	ENV_LOCATION=$(TEST_ENV_LOCATION) bash -c 'go clean -testcache && go test ./handlers/auth'

test_dumb:
	@./cmd/dump_test.sh $(DB_NAME) $(DB_USER) $(DB_PASSWORD) $(DB_REMOTE_USER) $(DB_REMOTE_PASSWORD)

drop_test_database:
	PGPASSWORD=$(DB_PASSWORD) dropdb -Umdgkb -hlocalhost $(DB_NAME)_test