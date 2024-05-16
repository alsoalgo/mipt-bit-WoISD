#!/bin/bash

curl http://webapp:8080/statistics >> /shared/statistics.log
echo "" >> /shared/statistics.log

