FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive

ARG ID_USER=1000
ARG ID_GROUP=1000
RUN groupadd -g ${ID_GROUP} app &&\
    useradd -u ${ID_USER} -m app -g app &&\
    usermod -a -G www-data app &&\
    apt-get update &&\
    apt-get install curl -y

ARG GOLANG_VERSION=1.20.4
RUN curl -OL https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz &&\
    tar -xvzf go${GOLANG_VERSION}.linux-amd64.tar.gz &&\
    mv go /usr/local/go

ENV PATH /usr/local/go/bin:${PATH}
