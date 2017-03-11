# golang_api_skel
Basic framework for Go API

[![Build Status](https://travis-ci.org/bigUNO/golang_api_skel.svg?branch=master)](https://travis-ci.org/bigUNO/golang_api_skel)

Mostly stolen code, pulling together for self-learning.

## Build
Creates an executable called api in the `build` directory.

```
go build -o build/api src/*
```

## Usage:

Return `Welcome!`. Great for testing
```
curl http://localhost:8080
```

Return all items
```
curl http://localhost:8080/items
```

Return single item
```
curl http://localhost:8080/items/2
```

Add item
```
curl -H "Content-Type: application/json" -d '{"name":"New Item"}' http://localhost:8080/items
```
