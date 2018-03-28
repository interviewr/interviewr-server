SHELL = /bin/bash
WORKDIR := $(PWD)

MODULE_NAME = interviewr-server
DOCKER_IMAGE_VERSION = develop
# Binaries to build.
TARGETS = bin/entry

default: build
.PHONY: default

build:
	@ echo "---> Building $(TARGETS) binary ..."
	@ rm -f $(TARGETS)
	@ env GOOS=linux GOARCH=386 go build -o $(WORKDIR)/$(TARGETS) ./cmd/$(MODULE_NAME)

dep:
	@ echo "---> Updating dependencies ..."
	@ dep ensure -update
.PHONY: dep

image-build:
	@ echo "---> Building Docker image ..."
	@ docker build -t $(MODULE_NAME):$(DOCKER_IMAGE_VERSION) $(WORKDIR)
.PHONY: image-build

image-publish:
	@ echo "---> Publishing Docker image ..."
	@ docker push ok2ju/interviewr-server:$(DOCKER_IMAGE_VERSION)
.PHONY: image-publish

image-run:
	@ echo "---> Running Docker container ..."
	@ docker run -it -p 8090:8090 -d --rm --name $(MODULE_NAME) $(MODULE_NAME):$(DOCKER_IMAGE_VERSION)
.PHONY: image-test