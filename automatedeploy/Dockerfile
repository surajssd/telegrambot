FROM centos:centos7

MAINTAINER "Suraj Deshmukh <surajd@redhat.com>"

# install golang and packaged dependecies, setup go env
RUN yum install -y golang git && \
    mkdir /go
ENV GOPATH=/go

WORKDIR /app
COPY ./ /app

# install dependencies
RUN go get github.com/Sirupsen/logrus && \
    go get gopkg.in/yaml.v2 

# build code and remove all installed things
RUN go build telegrambot.go && \
    yum remove -y golang git && \
    yum clean all && \
    rm -rf /go

ENTRYPOINT [ "/app/telegrambot" ]

