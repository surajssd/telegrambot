FROM centos:centos7

MAINTAINER "Suraj Deshmukh <surajd@redhat.com>"

# install golang and packaged dependecies, setup go env
RUN yum update -y && \
    yum install -y golang git && \
    yum clean all && \
    mkdir /go

ENV GOPATH=/go
ENV GOBIN=/go/bin
