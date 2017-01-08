#!/bin/bash

set -x

docker build -t docker.io/surajd/telegrambot .
docker push docker.io/surajd/telegrambot

