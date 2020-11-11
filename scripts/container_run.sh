#!/bin/bash

IMAGE_NAME=noraclock
CONTAINER_NAME=noraclock

if [[ "$#" -ne 1 ]]; then
  echo "Usage: <script-name> <tag>"
  exit 1
fi

echo "############### Removing Old Containers ###############"
docker rm -f $CONTAINER_NAME

echo "################ Running New Container ################"
if ! docker run \
  --detach \
  --name $CONTAINER_NAME \
  --restart unless-stopped \
  --env-file env/prod.env \
  --net host \
  --volume "$HOME"/service-logs/noraclock:/var/log \
  $IMAGE_NAME:"$1"; then
  echo "Failed to run container."
  exit 1
fi