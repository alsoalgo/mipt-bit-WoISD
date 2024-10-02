#!/bin/bash

docker build -t webapp-image:latest ./webapp
docker build -t script-image:latest ./script

# Из https://istio.io/latest/docs/setup/getting-started/
istioctl install -y --set profile=demo --set meshConfig.outboundTrafficPolicy.mode=REGISTRY_ONLY
kubectl label namespace default istio-injection=enabled

kubectl create -f services/webapp.yaml

kubectl create -f deployments/webapp.yaml 
kubectl create -f deployments/script.yaml

kubectl create -f gateways/gateway.yaml

kubectl create -f services/virtual.yaml
kubectl create -f services/external.yaml

minikube tunnel