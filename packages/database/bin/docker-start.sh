#!/bin/bash

docker info > /dev/null 2>&1
if [ $? = 1 ]; then
  echo "Error: Docker is not running. Please start the docker daemon."
  exit
fi

PROJECT_BASE=`git rev-parse --show-toplevel`
docker-compose --log-level ERROR -p roadtrip-postgres-offline \
  -f $PROJECT_BASE/packages/database/docker/docker-compose.yml up -d

echo ""
echo -n "Polling for Postgres to become available..."

end="$((SECONDS+180))"
while true; do
  nc -z localhost 5433
  [[ $? = 0 ]] && break
  [[ "${SECONDS}" -ge "${end}" ]] && echo "failed" && exit 1
  sleep 1
done

sleep 5

echo "Connection to localhost port 5433 [tcp/*] succeeded!"
echo ""

echo "Postgres is running in docker on port 5433."
