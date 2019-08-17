run:
	go run ./bin/launcher/main.go

build:
	docker build -t gateway:1.0 .

fmt:
	go fmt ./...

vet:
	go vet ./*

gometalinter:
	gometalinter ./*