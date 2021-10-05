#!/bin/bash

export PGDATABASE=postgres

if [ "$1" = "local" ]; then
  ENV=$1
fi

if [ -z $ENV ]; then
  echo "migrate-env env direction"
  echo "Environment required: [local]"
  exit
fi

DIR=$2
if [ -z $DIR ]; then
  echo "migrate-env env direction"
  echo "Direction required: [up, down]"
  exit
fi

OPTS=$3
if [ "$ENV" = "local" ]; then
  echo "Local migrate"
  DATABASE_URL=postgres://postgres:postgres@localhost:5432 yarn node-pg-migrate $DIR $OPTS
  exit
fi
