FROM docker.io/surajd/telegrambotbuilder


RUN go get -u -v gopkg.in/yaml.v2 && \
    go get -u -v github.com/Sirupsen/logrus


RUN git clone https://github.com/surajssd/telegrambot $GOPATH/src/github.com/surajssd/telegrambot && \
    cd $GOPATH/src/github.com/surajssd/telegrambot && \
    go install telegrambot.go

ENTRYPOINT [ "/go/bin/telegrambot" ]
