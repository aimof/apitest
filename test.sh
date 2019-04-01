#!/bin/sh
cd ./test && \
docker-compose build && \
docker-compose run apitest