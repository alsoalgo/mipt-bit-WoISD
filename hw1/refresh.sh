#!/bin/bash

minikube delete
minikube start
eval $(minikube docker-env)
./run.sh