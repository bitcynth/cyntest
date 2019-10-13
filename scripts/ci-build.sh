#!/bin/bash

set -e

version="$(git rev-parse --short HEAD)"
tag="$version-$BUILD_NUMBER"
time=$(date +'%Y-%m-%d_%T')

/usr/local/go/bin/go build -o cyntest -ldflags "-X main.buildVersion=$version -X main.buildTime=$time" ./src/

docker build -f Dockerfile -t bitcynth/cyntest:$tag .
docker push bitcynth/cyntest:$tag

kubectl set image deployment cyntest cyntest=bitcynth/cyntest:$tag