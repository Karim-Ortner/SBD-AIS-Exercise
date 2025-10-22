#!/bin/sh
set -e

docker stop orderservice postgres18 2>/dev/null || true
docker rm orderservice postgres18 2>/dev/null || true

# create network
docker network create orderapp-network 2>/dev/null || true

# todo
# docker build
docker build -t orderservice .
echo "build orderservice"

# docker run db

docker run -d \
  --name postgres18 \
  --env-file ./debug.env \
  --network orderapp-network \
  -v pgdata:/var/lib/postgresql/18/docker \
  -p 5432:5432 \
  postgres:18

echo "run db"
# docker run orderservice
sleep 10

#  -p 8080:8080 \
docker run -d \
  --name orderservice \
  --env-file ./debug.env \
  --network orderapp-network \
  -p 3000:3000 \
  orderservice:latest
echo "run orderservice"


echo "All services are up and running."
echo "Orderservice is accessible at http://localhost:3000"