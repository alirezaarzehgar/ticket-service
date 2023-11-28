APP_ENVS := MYSQL_HOST=localhost

all: prod

prod: stop .env
	go mod vendor
	docker-compose build
	docker-compose up -d

dev: .env
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

test:
	cp .env.example .env
	docker-compose -f docker-compose-test.yml up --exit-code-from unit-tests unit-tests
	docker-compose -f docker-compose-test.yml logs unit-tests