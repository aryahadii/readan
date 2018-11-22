ROOT := github.com/aryahadii/readan
.PHONY: clear docker-build docker-push dependencies rm-docker-containers

GO ?= go
GO_VARS ?=
GIT ?= git

readan: dependencies
	$(GO_VARS) $(GO) build -i -o="readan" $(ROOT)/cmd/readan

dependencies: vendor

clean: rm-docker-containers
	rm -rf ./readan
