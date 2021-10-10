#!/bin/bash

PROJECT_BASE=`git rev-parse --show-toplevel`
cd ${PROJECT_BASE}/packages/roadtrip-api

psql -h localhost -p 5433 -U postgres <<EOF
drop database roadtrip;
drop schema public cascade;
create schema public;
create database roadtrip;
EOF

cd $PROJECT_BASE/packages/database
yarn migrate local up

cd $PROJECT_BASE/packages/roadtrip-api
yarn prisma generate
yarn prisma db seed
