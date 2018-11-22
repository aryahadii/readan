GO ?= go
GO_VARS ?=
GIT ?= git

ROOT := github.com/aryahadii/readan
DOCKER_IMAGE := aryaha/readan

COMMIT := $(shell $(GIT) rev-parse HEAD)
VERSION ?= $(shell $(GIT) describe --tags ${COMMIT} 2> /dev/null || echo "$(COMMIT)")
BUILD_TIME := $(shell LANG=en_US date)

.PHONY: clear docker-build docker-push dependencies rm-docker-containers

readan: dependencies
	$(GO_VARS) $(GO) build -i -o="readan" $(ROOT)/cmd/readan

dependencies: vendor

clean: rm-docker-containers
	rm -rf ./readan

docker: Dockerfile
	docker build -t $(DOCKER_IMAGE):$(VERSION) .
	docker tag $(DOCKER_IMAGE):$(VERSION) $(DOCKER_IMAGE):latest

docker-push:
	docker push $(DOCKER_IMAGE):$(VERSION)
	docker push $(DOCKER_IMAGE):latest
