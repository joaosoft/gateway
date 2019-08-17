env:
	docker-compose up -d

run:
	go run ./bin/launcher/main.go

build:
	docker build -t gateway:1.0 .

push:
	docker login --username joaosoft
	docker tag gateway:1.0 joaosoft/gateway
	docker push joaosoft/gateway

fmt:
	go fmt ./...

vet:
	go vet ./*

gometalinter:
	gometalinter ./*