# interviewr-server [![Build Status](https://img.shields.io/travis/interviewr/interviewr-server/develop.svg?style=flat-square)](https://travis-ci.org/interviewr/interviewr-server) [![GitHub issues](https://img.shields.io/github/issues/interviewr/interviewr-server.svg?style=flat-square)](https://github.com/interviewr/interviewr-server/issues) [![GitHub contributors](https://img.shields.io/github/contributors/interviewr/interviewr-server.svg?style=flat-square)](https://github.com/interviewr/interviewr-server/graphs/contributors) [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)]( https://github.com/interviewr/interviewr-server/blob/develop/LICENSE)

Interviewr backend service. It's part of [interviewr-docker](https://github.com/interviewr/interviewr-docker).

## Prerequisites
* [dep](https://golang.github.io/dep/) - Dependency management tool for Go;
* [PostgreSQL](https://www.postgresql.org/)

## Running
1. Clone this repo
```sh
$ git clone https://github.com/interviewr/interviewr-server.git
```
2. Install dependencies
```sh
$ dep ensure
```
3. Build binary
```sh
make build
```
4. Run service
```sh
$ ./bin/entry
```

## Development
### Service configuration
You can configure service by editing `config.json` located at the root of project.

### Available commands
Build service:
```sh
make build
```
Update service dependencies:
```sh
make dep
```
Build docker image:
```sh
make image-build
```
Run docker container:
```sh
make image-run
```
Publish docker image to DockerHub registry:
```sh
make image-publish
```

## Docker image
Also you can grab prepared docker image from DockerHub
```sh
$ docker pull ok2ju/interviewr-server
```

## Notes
To keep docker image lightweight and simple it contains only resulting binary. So we build app locally and just copy binary.
