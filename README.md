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
$ export HOUR=10                     # DEFAULT: 12
$ export MINUTE=10                   # DEFAULT: 45
$ export NAMES=/some/path/names.yml  # DEFAULT: names.yml
$ export NOPINGDAYS="Sunday,Monday"  # DEFAULT: "Saturday,Sunday"
```

Finally run as:
```bash
$ ./telegrambot
```


### Deploy on docker-compose

```bash
export TOKEN=<token>
export WEBHOOK_URL=<url>
export NOPINGDAYS="Saturday"
export HOUR=18
export MINUTE=19
export NAMES="/names/names"

docker-compose up
```

### Deploy on openshift


Create an openshift project
```bash
oc new-project telegrambot
```

Expose all environments
```bash
export TOKEN=<token>
export WEBHOOK_URL=<url>
export NOPINGDAYS="Saturday"
export HOUR=18
export MINUTE=19
export NAMES="/names/names"
```

Convert all the configs using kompose
```bash
mkdir configs
kompose --provider openshift convert -o configs/ --build-repo https://github.com/surajssd/telegrambot
```

Create configmap for names
```bash
oc create configmap telegrambot --from-file=names=<path to names.yml>
```

Add the configmap info to deploymentconfig

```yaml
volumeMounts:
- name: namevol
  mountPath: "/names"
  readOnly: true
```
Add above to container in dc

```yaml
volumes:
- name: namevol
  configMap:
    name: telegrambot
```
Add above to pod in dc.

