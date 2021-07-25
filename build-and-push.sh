#!/usr/bin/env bash

docker build -t mihaershv/go-media-fetcher -f Dockerfile.prod ./src
docker image push mihaershv/go-media-fetcher