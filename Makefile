# force to use bash
SHELL = /bin/bash
WORKDIR := $(PWD)

.DEFAULT_GOAL: docker-image-build
.PHONY: default

docker-image-build:
	@ echo "---> Building Docker image ..."
	@ docker build -t interviewr-server .
	@ docker run -it -p 8090:8090 -d  --rm --name interviewr-server interviewr-server
.PHONY: docker-image-build

docker-image-publish:
	@ echo "---> Publishing Docker image ..."
	@ docker push ok2ju/interviewr-server:develop
.PHONY: docker-image-publish