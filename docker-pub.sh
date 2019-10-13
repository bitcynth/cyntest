#!/bin/bash
tag=$(git rev-parse --short HEAD)

docker build -f Dockerfile -t bitcynth/cyntest:$tag .
docker push bitcynth/cyntest:$tag