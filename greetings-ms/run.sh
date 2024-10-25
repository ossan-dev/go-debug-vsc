#!/bin/bash

docker build . --tag greeting-ms --file Dockerfile.debug
docker run -p 8080:8080 -p 2345:2345 greeting-ms