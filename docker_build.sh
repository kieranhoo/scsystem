#!/usr/bin/env bash

if [ "$#" -lt 2 ]; then
    echo "use: $0 with args [DOCKER_USERNAME DOCKER_IMAGE_NAME]"
    exit 1
fi

DOCKER_USERNAME=$1
DOCKER_IMAGE_NAME=$2
COMMIT_ID=$(git rev-parse --short HEAD)

