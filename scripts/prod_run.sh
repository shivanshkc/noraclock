#!/bin/bash

export CONF_FILE_PATH=conf/prod.json

echo "########### Building Application Binary ############"
if ! GOOS=linux go build -o bin/application; then
  echo "Failed to build application binary."
  exit 1
fi
echo "Application binary built."
echo "####################################################"

echo "############ Running Application Binary ############"
if ! bin/application; then
  echo "Application exited with non-zero status code."
  exit 1
fi
echo "####################################################"
