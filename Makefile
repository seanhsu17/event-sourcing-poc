PUB_NAME=event-publisher
SUB_NAME=event-subscriber
SCHE_NAME=event-scheduler
VERSION=latest

default: test

.PHONY: run
run: di build
	docker-compose -f deployments/docker-compose.yaml up

.PHONY: build
build:
	docker build -t $(PUB_NAME):$(VERSION) -f ./build/$(PUB_NAME)/Dockerfile .
	docker build -t $(SUB_NAME):$(VERSION) -f ./build/$(SUB_NAME)/Dockerfile .
	docker build -t $(SCHE_NAME):$(VERSION) -f ./build/$(SCHE_NAME)/Dockerfile .

.PHONY: stop
stop:
	docker-compose -f deployments/docker-compose.yaml down

.PHONY: di
di:
	go install github.com/google/wire/cmd/wire@v0.5.0
	wire gen ./internal/injector


.PHONY: test
test:
	go test -race -short ./...
