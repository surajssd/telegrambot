FROM docker.io/surajd/telegrambotbuilder


RUN go get -u -v gopkg.in/yaml.v2 && \
    go get -u -v github.com/Sirupsen/logrus && \
    go get -u -v github.com/spf13/cobra && \
    go get -u -v github.com/spf13/viper


RUN git clone https://github.com/surajssd/telegrambot $GOPATH/src/github.com/surajssd/telegrambot && \
    cd $GOPATH/src/github.com/surajssd/telegrambot && \
    go install

ENTRYPOINT [ "/go/bin/telegrambot" ]
