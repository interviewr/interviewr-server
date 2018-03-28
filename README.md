# interviewr-server
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
Clean artifacts under `bin/` folder:
```sh
make clean
```
Update service dependencies:
```sh
make dep
```
Build docker image:
```sh
make docker-image-build
```
Run docker container:
```sh
make docker-image-test
```
Publish docker image to DockerHub registry:
```sh
make docker-image-publish
```

## Docker image
Also you can grab prepared docker image from DockerHub
```sh
$ docker pull ok2ju/interviewr-server
```

## Notes
To keep docker image lightweight and simple it contains only resulting binary. So we build app locally and just copy binary.
