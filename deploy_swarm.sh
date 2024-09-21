#!/bin/bash

SWARM_INIT=$(docker info | grep 'Swarm: active')

if [ -z "$SWARM_INIT" ]; then
  echo "Init Docker Swarm..."
  docker swarm init
else
  echo "Docker Swarm is active"
fi

docker stack deploy -c docker-compose.python-cluster.yml python_cluster
docker service ls
