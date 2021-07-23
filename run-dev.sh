#!/usr/bin/env bash

if [ ! -f ./src/.env.local ]; then
  cp ./src/.env ./src/.env.local
fi

docker-compose stop
docker-compose build --pull --force-rm
docker-compose up -d --force-recreate