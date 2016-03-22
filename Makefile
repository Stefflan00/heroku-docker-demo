image-name=demo-app

setup:
	heroku apps:create golangdemo
	heroku docker:init

build:
	# godep save # needed to change "GoVersion": "go1.6beta2"
	docker-compose build

start:
	docker-compose up web

stop:
	docker-compose down

ping:
	curl -i ${DOCKER_URL}:9090/ping

deploy:
	heroku docker:release

prod_open:
	heroku open

prod_logs:
	heroku logs --tail
