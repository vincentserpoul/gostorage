# gostorage [![Go Report Card](https://goreportcard.com/badge/github.com/vincentserpoul/gostorage)](https://goreportcard.com/report/github.com/vincentserpoul/gostorage) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/ef8b36bfc33e468cae3168b58ee05680)](https://www.codacy.com/app/vincent_11/gostorage?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=vincentserpoul/gostorage&amp;utm_campaign=Badge_Grade)

init of different storage conn pool

## Prerequisites

* [go](https://golang.org/dl/) (tested on 1.9.3)
* [gometalinter](https://github.com/alecthomas/gometalinter)
* [dep](github.com/golang/dep/cmd/dep)
* [docker](https://www.docker.com/get-docker)

## dep issue

It seems there are still some issues with dep and cgo deps.
dep ensure and running the go-ethereum package test leads to this error.

```bash
# github.com/vincentserpoul/gostorage/vendor/github.com/ethereum/go-ethereum/crypto/secp256k1
vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/curve.go:44:10: fatal error: 'libsecp256k1/include/secp256k1.h' file not found
#include "libsecp256k1/include/secp256k1.h"
         ^~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
```


## Run tests

```bash
./run-tests.sh
```

## Ethereum

## MySQL

## Redis
