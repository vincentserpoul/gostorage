#!/bin/bash

# --vendor flag doesn't seem to work
gometalinter --exclude=vendor --deadline 5m  ./...

docker stack deploy --compose-file=infra/compose.yml storage;

# wait for the containers to be ready
sleep 20s

go test ./... -cover -bench=. -test.benchtime=3s -test.benchmem;

docker stack rm storage;