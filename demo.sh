#!/bin/sh
cd ./demo && \
docker-compose build && \
docker-compose run apitest