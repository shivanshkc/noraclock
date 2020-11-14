#!/bin/bash

echo "################## Sourcing Dev Configs ###################"
set -a
if ! . env/dev.env; then
  echo "Failed to source the configs."
  set +a
  exit 1
fi
echo "Configs successfully sourced."
set +a

echo "############### Building Application Binary ###############"
if ! CGO_ENABLED=0 GOOS=linux go build -o bin/service; then
  echo "Failed to build application binary."
  exit 1
fi
echo "Application binary successfully built."

echo "############### Running Application Binary ################"
if ! bin/service; then
  echo "Application exited with non-zero status code."
  exit 1
fi


