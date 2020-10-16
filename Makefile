DOCKER_COMPOSE := docker-compose
KILL_PROCESS   := "kill -9 \$$(ps aux | grep app | grep -v grep | awk '{print \$$1}')"

base:
	$(DOCKER_COMPOSE) --file tools/dev/base.yml up --detach
	-docker rm aigateway_db_waiter

codegen:
	$(DOCKER_COMPOSE) --file tools/dev/utility.yml run --rm codegen-api
	$(DOCKER_COMPOSE) --file tools/dev/utility.yml run --rm codegen-web

resolve-deps:
	$(DOCKER_COMPOSE) --file tools/dev/utility.yml run --rm deps-api

init: base codegen resolve-deps
	$(DOCKER_COMPOSE) --file tools/dev/utility.yml run --rm gen-private-key
	$(DOCKER_COMPOSE) --file tools/dev/utility.yml run --rm gen-public-key
	docker build --tag aigateway/api:cache --target cache --file tools/prod/api.Dockerfile .

up: base
	$(DOCKER_COMPOSE) --file tools/dev/runtime.yml up --detach

rebuild: base
	$(DOCKER_COMPOSE) --file tools/dev/runtime.yml up --detach --force-recreate --build

run: up
	-docker exec -t aigateway_api sh -c $(KILL_PROCESS)
	docker cp api/src aigateway_api:/go/src/github.com/xtream-d-labs/ai-gateway/api/
	docker exec -t aigateway_api go build -mod=vendor -o app main.go
	docker exec -t aigateway_api ./app --scheme http --host 0.0.0.0 --port 80

down:
	$(DOCKER_COMPOSE) --file tools/dev/runtime.yml down -v
	$(DOCKER_COMPOSE) --file tools/dev/base.yml down

test/src: base
	$(DOCKER_COMPOSE) --file tools/dev/utility.yml run --rm lint-api
	$(DOCKER_COMPOSE) --file tools/dev/utility.yml run --rm test-api
