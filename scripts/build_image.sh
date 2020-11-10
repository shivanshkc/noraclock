#!/bin/bash

if [[ "$#" -ne 1 ]]; then
  echo "Usage: <script-name> <tag>"
  exit 1
fi

IMAGE_NAME=noraclock

echo "############### Building Image ###############"
if ! docker build -t $IMAGE_NAME:"$1" .; then
  echo "Failed to build image."
  exit 1
fi
echo "Image successfully built."
