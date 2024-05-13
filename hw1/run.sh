#!/bin/bash

docker build -t webapp-image:latest ./webapp
docker build -t script-image:latest ./script

kubectl create -f services/webapp.yaml

kubectl create -f deployments/webapp.yaml 
kubectl create -f deployments/script.yaml

minikube tunnel