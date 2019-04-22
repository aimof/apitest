#!/bin/sh
docker-compose -f ./test/docker-compose.yml build && \
docker-compose -f ./test/docker-compose.yml run apitest