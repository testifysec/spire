ARG GO_VER=1.18.1

FROM golang:$GO_VER
RUN apt update && apt install -y \
    ca-certificates \
    curl \
    gnupg2 \
    software-properties-common \
    git \
    curl \
    podman 

RUN apt install git docker curl -y

COPY ./witness /usr/local/bin/witness

RUN mkdir -p /code

WORKDIR /code
