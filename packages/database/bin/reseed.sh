#!/bin/bash

PROJECT_BASE=`git rev-parse --show-toplevel`
cd ${PROJECT_BASE}/packages/database

psql -h localhost -p 5433 -U postgres <<EOF
drop database roadtrip;
drop schema public cascade;
create schema public;
create database roadtrip;
EOF

yarn migrate local up

cat prisma/reseed.prisma.header prisma/schema.prisma.base > prisma/schema.prisma
yarn prisma generate
yarn prisma db seed

rm prisma/schema.prisma

