FROM centos:centos7

MAINTAINER "Suraj Deshmukh <surajd@redhat.com>"

COPY ./telegrambot /

ENTRYPOINT [ "/telegrambot" ]

