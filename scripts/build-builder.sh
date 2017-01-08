#!/bin/bash

set -x

docker build -t docker.io/surajd/telegrambotbuilder -f Dockerfile.builder .
docker push docker.io/surajd/telegrambotbuilder
