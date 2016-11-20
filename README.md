# telegrambot


### Build:

```bash
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
$ export HOUR=10    # DEFAULT: 12
$ export MINUTE=10  # DEFAULT: 45
$ export NAMES=<path to names file> # DEFAULT: names.yml
```
