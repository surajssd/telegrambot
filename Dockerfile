FROM centos:7

MAINTAINER "Suraj Deshmukh <surajd@redhat.com>"

COPY ./telegrambot /

ENTRYPOINT [ "/telegrambot" ]

