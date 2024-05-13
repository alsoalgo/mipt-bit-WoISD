#!/bin/bash


while true; do
  curl http://webapp:8080/statistics >> statistics.log
  echo "" >> statistics.log
  sleep 5
done
