all: build

.PHONY: gen
gen:
	protoc --go_out=plugins=micro:. proto/scores/*.proto

.PHONY: build
build:
	CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-w' -i -o scores-srv ./main.go ./plugins.go

.PHONY: docker-build
docker-build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -i -o scores-srv ./main.go ./plugins.go

.PHONY: docker-docker
docker: docker-build
	docker build --no-cache -t micro/scores-srv .