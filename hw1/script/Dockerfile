FROM ubuntu:latest

RUN apt-get update \
 && apt-get install -y \
      curl \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /script

COPY ./script.sh ./script.sh

CMD ./script.sh