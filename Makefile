APP_ENVS := MYSQL_HOST=localhost

all: prod

help: ## Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

prod: stop .env swag ## Deploy project
	go mod vendor
	docker-compose build
	docker-compose up -d

dev: .env swag ## Run database on docker and code on local machine
	docker-compose up db -d
	$(APP_ENVS) go run .

.env: ## Create default .env and log, assets directories
	cp .env.example .env
	mkdir -p log assets

stop: ## Stop and remove all containers
	yes | docker-compose rm
	docker-compose stop

test-compile: swag ## Compile project
	cp .env.example .env
	go mod vendor
	go build .


test: ## Run tests and return status code of unit tests
	docker-compose -f docker-compose-test.yml up --exit-code-from unit-tests unit-tests
	docker-compose -f docker-compose-test.yml logs unit-tests
	docker-compose -f docker-compose-test.yml down unit-tests --remove-orphans


test-dev: .env swag ## Run database on container and tests on local machine
	docker-compose up db -d
	$(APP_ENVS) go test -v ./...

swag: ## Install swagger and format/init comments
	go install github.com/swaggo/swag/cmd/swag@latest
	swag fmt
	swag init
