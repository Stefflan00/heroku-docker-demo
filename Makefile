image-name=demo-app

build:
	docker build --tag=${image-name} .

start:
	docker run -it --rm -p 9090:8080 ${image-name}

stop:
	docker stop `docker ps --filter ancestor=${image-name} --format="{{.ID}}"`

ping:
	curl -i ${DOCKER_URL}:9090/ping
