# golang_api_skel
[![Build Status](https://travis-ci.org/bigUNO/golang_api_skel.svg?branch=master)](https://travis-ci.org/bigUNO/golang_api_skel)
Basic framework for Go API

* [Build](#build)
* [Configuration](#configuration)
 * [Docker](##docker)
* [Testing](#testing)
* [Usage](#usage)

# Build
Creates an executable called api in the `build` directory.

```
go build -o build/golang_api_skel src/*
```

# Config
A config file is needed

```sh
title = "golang_api_skel"

[datebase]
mongodbhosts = "mongodb1.example.com:27017"
authdatabase = "drbipolar"
authusername = "koolkeith"
authpassword = "oakridgecake"
isdrop       = false
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

## Want to contribute?
You can [create a GitHub issue](https://github.com/bigUNO/golang_api_skel/issues/new) for any feature requests, bugs, or documentation improvements.

Where possible, please [submit a pull request](https://help.github.com/articles/creating-a-pull-request-from-a-fork/) for the change.
