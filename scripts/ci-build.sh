#!/bin/bash
tag="$(git rev-parse --short HEAD)-$BUILD_NUMBER"

docker build -f Dockerfile -t bitcynth/cyntest:$tag .
docker push bitcynth/cyntest:$tag

kubectl set image deployment cyntest cyntest=bitcynth/cyntest:$tag