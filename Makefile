DOCKER_COMPOSE := docker-compose
KILL_PROCESS   := "kill -9 \$$(ps aux | grep app | grep -v grep | awk '{print \$$1}')"
PWD            := $(shell pwd)

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
	sed -e "s|<replace-your-absolute-path>|${PWD}|g" .vscode/launch-template.json >.vscode/launch.json

up: base
	$(DOCKER_COMPOSE) --file tools/dev/runtime.yml up --detach
	$(DOCKER_COMPOSE) --file tools/dev/runtime.yml logs -f

down:
	$(DOCKER_COMPOSE) --file tools/dev/runtime.yml down -v
	$(DOCKER_COMPOSE) --file tools/dev/base.yml down

test/src: base
	$(DOCKER_COMPOSE) --file tools/dev/utility.yml run --rm lint-api
	$(DOCKER_COMPOSE) --file tools/dev/utility.yml run --rm test-api
