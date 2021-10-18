#!/bin/bash

PROJECT_BASE=`git rev-parse --show-toplevel`
cd ${PROJECT_BASE}/packages/roadtrip-api

cat prisma/schema.prisma.header ../database/prisma/schema.prisma.base > prisma/schema.prisma

yarn prisma generate

rm prisma/schema.prisma
