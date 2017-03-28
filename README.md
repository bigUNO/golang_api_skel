# golang_api_skel
[![Build Status](https://travis-ci.org/bigUNO/golang_api_skel.svg?branch=master)](https://travis-ci.org/bigUNO/golang_api_skel)
Basic framework for Go API

* [Build](#build)
 * [Docker](##docker)
* [Testing](#testing)
* [Usage](#usage)

# Build
Creates an executable called api in the `build` directory.

```
go build -o build/golang_api_skel src/*
```

## Docker
Designed to run inside a docker container.

Docker healthchecks curl `/healthcheck` to see if app is alive.

# Testing
Use docker to test

```sh
docker build -t golang_api_skel .
docker run -d --publish 6060:8080 --name test --rm golang_api_skel
```
Test the healthcheck
```sh
curl http://localhost:6060/healthcheck
```

Clean up
```sh
docker stop test
```

# Usage
Return `{"alive": true`. Great for testing
```sh
curl http://localhost:8080/healthcheck
```

Return all items
```sh
curl http://localhost:8080/items
```

Return single item
```sh
curl http://localhost:8080/items/b379fdbaqj868olu1gi0
```

Note the example above will not work, you need to grab the UUID of one of the
test items.

Add item
```sh
curl -H "Content-Type: application/json" -d '{"name":"soda"}' http://localhost:8080/items
```
