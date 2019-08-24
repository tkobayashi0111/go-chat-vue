#!/bin/bash

# Exit on any error
set -e

# get TAG
TAG=$(git rev-parse --short HEAD)

# docker build
docker build -t gcr.io/${GOOGLE_PROJECT_ID}/proxy:${TAG} -t gcr.io/${GOOGLE_PROJECT_ID}/proxy:latest -f docker/proxy/Dockerfile-prd .
docker build -t gcr.io/${GOOGLE_PROJECT_ID}/server:${TAG} -t gcr.io/${GOOGLE_PROJECT_ID}/server:latest -f docker/server/Dockerfile-prd .
docker build -t gcr.io/${GOOGLE_PROJECT_ID}/client:${TAG} -t gcr.io/${GOOGLE_PROJECT_ID}/client:latest -f docker/client/Dockerfile-prd .

# docker push
docker push gcr.io/${GOOGLE_PROJECT_ID}/proxy:${TAG}
docker push gcr.io/${GOOGLE_PROJECT_ID}/proxy:latest
docker push gcr.io/${GOOGLE_PROJECT_ID}/server:${TAG}
docker push gcr.io/${GOOGLE_PROJECT_ID}/server:latest
docker push gcr.io/${GOOGLE_PROJECT_ID}/client:${TAG}
docker push gcr.io/${GOOGLE_PROJECT_ID}/client:latest

# update deployment
REPLICAS=$(kubectl get deployment/go-chat -o jsonpath='{.spec.replicas}')
kubectl scale --replicas=0 deployment go-chat
kubectl scale --replicas=${REPLICAS} deployment go-chat
