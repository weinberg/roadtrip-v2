#!/bin/bash

PROJECT_BASE=`git rev-parse --show-toplevel`
cd ${PROJECT_BASE}/packages/go

cat prisma/schema.prisma.header ../database/prisma/schema.prisma.base > prisma/schema.prisma

go run github.com/prisma/prisma-client-go generate

rm prisma/schema.prisma
