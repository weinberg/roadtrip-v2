#!/bin/bash

PROJECT_BASE=`git rev-parse --show-toplevel`
docker-compose --log-level ERROR -p roadtrip-postgres-offline \
  -f $PROJECT_BASE/packages/database/docker/docker-compose.yml down
