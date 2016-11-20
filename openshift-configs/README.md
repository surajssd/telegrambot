# Steps to deploy app


Create OpenShift project
```bash
$ oc new-project bot
```

Create `secrets` for `TOKEN` and `WEBHOOK_URL`
```bash
$ oc create secret generic telegrambot --from-literal='token=<your bot token>,webhook=<your web hook url>'
```

Create `configmap` for `NAMES`
```bash
$ oc create configmap telegrambot --from-file=names=<path to names.yml file>
```

Create `deploymentconfig` and `imagestream` from files in this directory, which were auto-generated using kompose but modified to use `secrets` and `configmaps`
```bash
$ oc create -f .
```
