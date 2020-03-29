GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o beego_blog main.go

.PHONY: docker
docker:
	docker build . -t beego_blog:v0.0.1

.PHONY: run
run:
	go mod download
	go run main.go



