#!/bin/bash

IMAGE_NAME=noraclock
CONTAINER_NAME=noraclock

VOLUME_DIR=$HOME/docker/volumes/$CONTAINER_NAME

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
  --net host \
  --volume $VOLUME_DIR/app-conf:/etc/app \
  --volume $VOLUME_DIR/app-logs:/var/log \
  --env CONF_FILE_PATH=/etc/app/conf.json \
  $IMAGE_NAME:"$1"; then
  echo "Failed to run container."
  exit 1
fi


