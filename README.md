# telegrambot

### Get telegrambot for usage:

```bash
$ go get github.com/surajssd/telegrambot
```

### Development environment

Fork this repo

```bash
$ git clone https://github.com/<YOUR GITHUB ID>/telegrambot $GOPATH/src/github.com/surajssd/telegrambot
$ cd $GOPATH/src/github.com/surajssd/telegrambot
$ go build
```

### Run:

Required environment variables:
```bash
$ export TOKEN=<your bot token>
$ export WEBHOOK_URL=<your webhook url>
```

Optional environment variables:
```bash
$ export HOUR=10                    # DEFAULT: 12
$ export MINUTE=10                  # DEFAULT: 45
$ export NAMES=<path to names file> # DEFAULT: names.yml
```

Finally run as:
```bash
$ ./telegrambot
```
