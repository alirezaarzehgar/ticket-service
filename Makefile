APP_ENVS := MYSQL_HOST=localhost

all: prod

prod: stop .env swag
	go mod vendor
	docker-compose build
	docker-compose up -d

dev: .env swag
	docker-compose up db -d
	$(APP_ENVS) go run .

.env:
	cp .env.example .env
	mkdir logs assets

stop:
	yes | docker-compose rm
	docker-compose stop

update:
	git pull -f
	make prod

test-compile: swag
	cp .env.example .env
	go mod vendor
	go build .


test:
	docker-compose -f docker-compose-test.yml up --exit-code-from unit-tests unit-tests
	docker-compose -f docker-compose-test.yml logs unit-tests
	docker-compose -f docker-compose-test.yml down unit-tests --remove-orphans

swag:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag fmt
	swag init
